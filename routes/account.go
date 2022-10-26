package routes

import (
	"name/shoppingcart/controllers"

	"github.com/gofiber/fiber/v2"
)

func InitAccRoute(r *fiber.App) {
	userControllers := controllers.InitUserController()
	authController := controllers.InitAuthController()
	r.Post("/register", userControllers.PostedCreatedAccount)
	r.Post("/login", authController.LoginPosted)
}
