package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        uint           `form:"id" json:"id" gorm:"primaryKey"`
	UserID    uint           `form:"userId" json:"userId" validate:"required"`
	CartID    uint           `form:"cartId" json:"cartId"`
	Total     float64        `form:"total json:total"`
	CreatedAt time.Time      `form:"createdAt" json:"createdAt"`
	UpdatedAt time.Time      `form:"updatedAt" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateTransaksi(db *gorm.DB, newTransaction *Transaction, userId, cartId uint) (err error) {
	newTransaction.UserID = userId
	newTransaction.CartID = cartId
	if err := db.Create(newTransaction).Error; err != nil {
		return err
	}
	return nil
}

func ReadTransactionById(db *gorm.DB, transaction *Transaction, id int) (err error) {
	if err = db.Where("id=?", id).First(transaction).Error; err != nil {
		return err
	}
	return nil
}

func SaveTransaction(db *gorm.DB, transaction *Transaction) (err error) {
	if err := db.Save(transaction).Error; err != nil {
		return err
	}
	return nil
}
