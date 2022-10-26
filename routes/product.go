package routes

import (
	"name/shoppingcart/controllers"

	"github.com/gofiber/fiber/v2"
)

func InitProductsRoute(r *fiber.App) {
	productControllers := controllers.InitProductController()
	p := r.Group("/products")
	p.Get("/", productControllers.IndexProduct)
	p.Post("/create", productControllers.AddPostedProduct)           // input form name, qunatity, price, picture
	p.Get("/detail/:id", productControllers.GetDetailProduct)        // input params path id
	p.Put("/edit/product/:id", productControllers.EditProductAPI)    // input params path id
	p.Delete("/deleteproduct/:id", productControllers.DeleteProduct) // input params path id
}
