package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"name/shoppingcart/database"
	"name/shoppingcart/models"
)

type UserController struct {
	Db *gorm.DB
}

func InitUserController() *UserController {
	db := database.InitDb()

	// migrate the schema
	db.AutoMigrate(&models.User{})

	return &UserController{Db: db}
}

// Post /register
func (controller *UserController) PostedCreatedAccount(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(400)
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	sHash := string(bytes)
	// password dalam bentuk hash yang akan disimpan di database
	user.Password = sHash

	// save user ke database
	if err := models.CreateUser(controller.Db, &user); err != nil {
		return c.SendStatus(500)
	}

	// generate cartId setelah
	var cart models.Cart
	if err := models.CreateCart(controller.Db, &cart, uint(user.ID)); err != nil {
		c.SendStatus(500) //error server
	}

	return c.JSON(fiber.Map{
		"message": "registrasi berhasil!, lanjut ke /login",
		"user":    user,
		"cart":    cart,
	})
}
