package ports

import (
	"app/core/models"
	"context"
)

type UserSaver interface {
	Save(context.Context, *models.User) (uint, error)
}

type UserSelectorByPK interface {
	SelectByPK(context.Context, uint) (models.User, error)
}

type UserDeleterByPK interface {
	DeleteByPK(context.Context, string) (uint, error)
}

type UserRepo interface {
	UserSaver
	UserSelectorByPK
	UserDeleterByPK
}
