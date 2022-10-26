package routes

import (
	"name/shoppingcart/controllers"

	"github.com/gofiber/fiber/v2"
)

func InitShoppingCartRoute(r *fiber.App) {
	cartController := controllers.InitCartController()
	transactionController := controllers.InitTransactionController()
	s := r.Group("/shoppingcart")
	s.Post("/addtocart/:productid/:userid", cartController.AddProductToCart) // input params path cartid & productid
	s.Get("/:userid", cartController.GetCart)                                // input params path userid
	s.Get("/checkout/:userid", transactionController.GetTransaction)         // input params path userid
}
