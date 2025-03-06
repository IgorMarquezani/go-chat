package repository

import (
	"app/adapters/database"
	"app/core/models"
	"context"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() (UserRepository, error) {
	db, err := database.GetDBConn(5)

	return UserRepository{
		db: db,
	}, err
}

func (r UserRepository) Save(ctx context.Context, u *models.User) (uint, error) {
	db := r.db.WithContext(ctx).Save(u)
	return uint(db.RowsAffected), db.Error
}

func (r UserRepository) SelectByPK(ctx context.Context, id uint) (models.User, error) {
	u := models.User{}

	err := r.db.WithContext(ctx).Where("id = ?", id).First(&u).Error

	return u, err
}

func (r UserRepository) DeleteByPK(ctx context.Context, id string) (uint, error) {
	db := r.db.Where("id = ?", id).Delete(models.User{})

	return uint(db.RowsAffected), db.Error
}
