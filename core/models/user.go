package models

type User struct {
	ID           string `json:"id" gorm:"type:uuid;primaryKey"`
	Name         string `json:"name" gorm:"not null"`
	Email        string `json:"email" gorm:"unique;not null"`
	Password     string `json:"-" gorm:"not null"`
	ProfileImage string `json:"profile_image"`
}
