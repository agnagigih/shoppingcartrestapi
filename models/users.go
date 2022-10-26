package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `form:"id" json:"id" gorm:"primaryKey"`
	Name      string         `form:"name" json:"name" validate:"required"`
	Email     string         `form:"email" json:"email" validate:"required"`
	Username  string         `form:"username" json:"username" validate:"required"`
	Password  string         `form:"password" json:"password" validate:"required"`
	CreatedAt time.Time      `form:"createdAt" json:"createdAt"`
	UpdatedAt time.Time      `form:"updatedAt" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CRUD
func CreateUser(db *gorm.DB, newUser *User) (err error) {
	if err := db.Create(newUser).Error; err != nil {
		return err
	}
	return nil
}

func FindUserByUsername(db *gorm.DB, user *User, username string) (err error) {
	if err := db.Where("username=?", username).First(user).Error; err != nil {
		return err
	}
	return nil
}
func UpdateUserProfile(db *gorm.DB, user *User) (err error) {
	db.Save(user)

	return nil
}
