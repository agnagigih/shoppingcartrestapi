package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID         uint           `form:"id" json:"id" gorm:"primaryKey"`
	UserID     uint           `form:"userid" json:"userid"`
	ProductIDs string         `form:"productid" json:"productid"`
	CreatedAt  time.Time      `form:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time      `form:"updatedAt" json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func CreateCart(db *gorm.DB, cart *Cart, userID uint) (err error) {
	cart.UserID = userID
	if err = db.Create(cart).Error; err != nil {
		return err
	}
	return nil
}

func ReadCartByUserId(db *gorm.DB, cart *Cart, id uint) (err error) {
	if err = db.Where("user_id=?", id).Last(cart).Error; err != nil {
		return err
	}
	return nil
}

func ReadCartById(db *gorm.DB, cart *Cart, cartId uint) (err error) {
	if err = db.Where("id=?", cartId).First(cart).Error; err != nil {
		return err
	}
	return nil
}

func SaveCart(db *gorm.DB, insertedCart *Cart) (err error) {
	if err := db.Save(insertedCart).Error; err != nil {
		return err
	}
	return nil
}
