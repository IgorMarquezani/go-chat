package hasher

import "golang.org/x/crypto/bcrypt"

type Hasher struct{}

func (h Hasher) HashPasswd(str string) (string, error) {
	arr, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(arr), err
}
