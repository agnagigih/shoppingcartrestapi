package controllers

import (
	"name/shoppingcart/database"
	"name/shoppingcart/models"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginForm struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

type AuthController struct {
	// declare variables
	Db *gorm.DB
}

func InitAuthController() *AuthController {
	db := database.InitDb()

	// migrate the schema
	db.AutoMigrate(&models.Product{})

	return &AuthController{Db: db}
}

// post /login
func (controller *AuthController) LoginPosted(c *fiber.Ctx) error {
	var form LoginForm
	if err := c.BodyParser(&form); err != nil {
		return c.SendStatus(500)
	}

	// auth
	var userDb models.User
	usernameDb := form.Username

	// mencari user dari database

	if err := models.FindUserByUsername(controller.Db, &userDb, usernameDb); err != nil {
		return c.JSON(fiber.Map{
			"message": "akun tidak ditemukan",
		}) // akun tidak ditemukan
	}
	// bytes, _ := bcrypt.GenerateFromPassword([]byte(userDb.SHash), 10)

	if err := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(form.Password)); err == nil {
		// Create the Claims
		exp := time.Now().Add(time.Hour * 72)
		claims := jwt.MapClaims{
			"name":  userDb.Name,
			"admin": true,
			"exp":   exp.Unix(),
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secretpassword"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"message": "berhasil login! lanjut ke /products",
			"token":   t,
			"expired": exp.Format("2006-01-02 15:04:05"),
			"name":    userDb.Name,
		}) // berhasil login
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}

// logout
func (controller *AuthController) Logout(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"message": "logut sukses!",
	})
}
