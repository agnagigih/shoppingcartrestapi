package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"name/shoppingcart/database"
	"name/shoppingcart/models"
)

type TransactionController struct {
	Db *gorm.DB
}

func InitTransactionController() *TransactionController {
	db := database.InitDb()
	db.AutoMigrate(&models.Transaction{})
	return &TransactionController{Db: db}
}

// Get"/checkout/:userid"
func (controller *TransactionController) GetTransaction(c *fiber.Ctx) error {
	params := c.AllParams()
	userId, _ := strconv.Atoi(params["userid"])
	var cart models.Cart
	if err := models.ReadCartByUserId(controller.Db, &cart, uint(userId)); err != nil {
		return c.SendStatus(500)
	}
	var transaction models.Transaction
	if err := models.CreateTransaksi(controller.Db, &transaction, uint(userId), cart.ID); err != nil {
		return c.SendStatus(500)
	}
	sliceProductID := ConvertToSliceUint(cart.ProductIDs)
	var someProduct []*ProductCart
	var total float64 = 0
	for i := 0; i < len(sliceProductID); i++ {
		var product models.Product
		if err := models.ReadProductById(controller.Db, &product, int(sliceProductID[i])); err != nil {
			return c.JSON(fiber.Map{
				"message": err,
			})
		}
		var currentProduct ProductCart
		currentProduct.Name = product.Name
		currentProduct.Price = product.Price
		currentProduct.Quantity = 1
		total += float64(currentProduct.Price) * float64(currentProduct.Quantity)
		someProduct = append(someProduct, &currentProduct)
	}
	transaction.Total = total
	if err := models.SaveTransaction(controller.Db, &transaction); err != nil {
		return c.SendStatus(500) //error server
	}
	// reset cart.productIDs to ""
	cart.ProductIDs = ""
	if err := models.SaveCart(controller.Db, &cart); err != nil {
		return c.SendStatus(500) //error server
	}
	return c.JSON(fiber.Map{
		"transaction": transaction,
		"product":     someProduct,
	})
}
