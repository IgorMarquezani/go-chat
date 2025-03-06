package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"app/adapters/database/repository"
	"app/adapters/hasher"
	"app/adapters/storage"
	"app/adapters/validator"
	"app/core/controllers"
	"app/core/models"
	"app/core/ports"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var signUpPool = sync.Pool{
	New: func() any {
		return make([]byte, 512)
	},
}

func Signup(c *fiber.Ctx) error {
	type Form struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type Service struct {
		Saver   ports.UserSaver
		Deleter ports.UserDeleterByPK
		Hasher  ports.PasswdHasher
		Storage ports.Storage
		Passwd  ports.PasswdValidator
	}

	var (
		Err  = controllers.Error{}
		form = Form{}
	)

	fileHeader, err := c.FormFile("profile-image")
	if err != nil {
		Err.Message = "invalid file"
		Err.Description = "expected form field name 'profile-image' containing a image file"
		c.Status(http.StatusBadRequest)
		return c.JSON(&Err)
	}

	src, err := fileHeader.Open()
	if err != nil {
		Err.Message = "server error"
		Err.Description = "internal server error"
		c.Status(http.StatusInternalServerError)
		return c.JSON(&Err)
	}
	defer src.Close()

	buf := signUpPool.Get().([]byte)

	if _, err := src.Read(buf); err != nil && !errors.Is(err, io.EOF) {
		Err.Message = "server error"
		Err.Description = "internal server error"
		c.Status(http.StatusInternalServerError)
		return c.JSON(&Err)
	}

	contentType := http.DetectContentType(buf)
	if contentType != "image/jpeg" && contentType != "image/png" {
		Err.Message = "invalid file"
		Err.Description = "invalid file type for image"
		Err.Details = append(Err.Details, fmt.Sprintf("file is of type %s and expected image/jpeg or image/png", contentType))
		c.Status(http.StatusBadRequest)
		return c.JSON(&Err)
	}

	signUpPool.Put(buf)

	rawJSON := c.FormValue("json")

	reader := strings.NewReader(rawJSON)

	if err := json.NewDecoder(reader).Decode(&form); err != nil {
		Err.Message = "JSON error"
		Err.Description = "invalid json syntax or wrong data type"
		Err.Details = append(Err.Details, err.Error())
		c.Status(http.StatusBadRequest)
		return c.JSON(&Err)
	}

	repo, err := repository.NewUserRepository()
	if err != nil {
		Err.Message = "server error"
		Err.Description = "internal server error"
		c.Status(http.StatusInternalServerError)
		return c.JSON(&Err)
	}

	svc := Service{
		Saver:   repo,
		Deleter: repo,
		Hasher:  hasher.Hasher{},
		Storage: storage.FS{},
		Passwd:  validator.Password{},
	}

	if errs := svc.Passwd.Validate(form.Password); len(errs) > 0 {
		Err.Message = "invalid passwod"
		Err.Description = fmt.Sprintf("the given passwod is not valid")

		for _, err := range errs {
			Err.Details = append(Err.Details, err.Error())
		}

		c.Status(http.StatusBadRequest)
		return c.JSON(&Err)
	}

	hashedPasswd, err := svc.Hasher.HashPasswd(form.Password)
	if err != nil {
		Err.Message = "server error"
		Err.Description = "internal server error"
		c.Status(http.StatusInternalServerError)
		return c.JSON(&Err)
	}

	userID := uuid.NewString()

	format := strings.Split(contentType, "/")[1]

	u := models.User{
		ID:           userID,
		Name:         form.Name,
		Email:        form.Email,
		Password:     hashedPasswd,
		ProfileImage: userID + "." + format,
	}

	if _, err := svc.Saver.Save(c.Context(), &u); err != nil {
		strs := strings.SplitN(err.Error(), " ", 4)

		if len(strs) > 2 && strs[1] == "duplicate" && strs[2] == "key" {
			Err.Message = "email already in use"
			Err.Description = fmt.Sprintf("the given email (%s) is already in use", form.Email)
			c.Status(http.StatusBadRequest)
			return c.JSON(&Err)
		}

		Err.Message = "server error"
		Err.Description = "internal server error"
		c.Status(http.StatusInternalServerError)
		return c.JSON(&Err)
	}

	src.Seek(0, 0)

	if err := svc.Storage.Create(userID+"."+format, src); err != nil {
		svc.Deleter.DeleteByPK(c.Context(), userID)

		Err.Message = "server error"
		Err.Description = "internal server error"

		c.Status(http.StatusInternalServerError)
		return c.JSON(&Err)
	}

	return c.SendString("ok")
}
