package main

import (
	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/jwt/v3"

	"name/shoppingcart/routes"
)

func main() {
	app := fiber.New()
	// Init route login & register
	routes.InitAccRoute(app)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secretpassword"),
	}))
	//restricted route
	routes.InitProductsRoute(app)
	routes.InitShoppingCartRoute(app)

	app.Listen(":3000")
}
