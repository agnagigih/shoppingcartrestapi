package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `form:"id" json:"id" gorm:"primaryKey"`
	Name      string         `form:"name" json:"name" validate:"required"`
	Quantity  int            `form:"quantity" json:"quantity" validate:"required"`
	Price     float32        `form:"price" json:"price" validate:"required"`
	Picture   string         `form:"picture" json:"picture" validate:"required"`
	CreatedAt time.Time      `form:"createdAt" json:"createdAt"`
	UpdatedAt time.Time      `form:"updatedAt" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CRUD
func CreateProduct(db *gorm.DB, newProduct *Product) (err error) {
	if err := db.Create(newProduct).Error; err != nil {
		return err
	}
	return nil
}
func ReadProducts(db *gorm.DB, products *[]Product) (err error) {
	if err := db.Find(products).Error; err != nil {
		return err
	}
	return nil
}
func ReadProductById(db *gorm.DB, product *Product, id int) (err error) {
	if err = db.Where("id=?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}
func UpdateProduct(db *gorm.DB, product *Product) (err error) {
	db.Save(product)

	return nil
}
func DeleteProductById(db *gorm.DB, product *Product, id int) (err error) {
	db.Where("id=?", id).Delete(product)

	return nil
}
