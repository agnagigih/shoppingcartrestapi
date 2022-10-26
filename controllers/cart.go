package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"name/shoppingcart/database"
	"name/shoppingcart/models"
	"strconv"
	"strings"
)

type CartController struct {
	// declare variables
	Db *gorm.DB
}
type ProductCart struct {
	Name     string
	Price    float32
	Quantity uint
}

// initial
func InitCartController() *CartController {
	db := database.InitDb()

	// migrate the schema
	db.AutoMigrate(&models.Cart{})

	return &CartController{Db: db}
}
func ConvertToSliceUint(str string) []uint {
	sliceString := strings.Split(str, "; ")
	var conv []uint
	for i := 0; i < len(sliceString); i++ {
		num, _ := strconv.Atoi(sliceString[i])
		conv = append(conv, uint(num))
	}
	return conv
}

func (controller *CartController) GetCart(c *fiber.Ctx) error {
	params := c.AllParams()
	userId, _ := strconv.Atoi(params["userid"])
	var cart models.Cart
	if err := models.ReadCartByUserId(controller.Db, &cart, uint(userId)); err != nil {
		return c.SendStatus(500) //error server
	}
	var someProduct []*ProductCart
	var total float64 = 0
	if cart.ProductIDs != "" {
		sliceProductID := ConvertToSliceUint(cart.ProductIDs)
		// var someProduct []*ProductCart
		// var total float64 = 0
		for i := 0; i < len(sliceProductID); i++ {
			var product models.Product
			if err := models.ReadProductById(controller.Db, &product, int(sliceProductID[i])); err != nil {
				return c.JSON(fiber.Map{
					"message": "Cart kosong!",
				})
			}
			var currentProduct ProductCart
			currentProduct.Name = product.Name
			currentProduct.Price = product.Price
			currentProduct.Quantity = 1
			total += float64(currentProduct.Price) * float64(currentProduct.Quantity)
			someProduct = append(someProduct, &currentProduct)
		}
	}

	return c.JSON(fiber.Map{
		"cart":     cart,
		"products": someProduct,
		"total":    total,
	})
}

func (controller *CartController) AddProductToCart(c *fiber.Ctx) error {
	params := c.AllParams()

	productId, _ := strconv.Atoi(params["productid"])
	userId, _ := strconv.Atoi(params["userid"])
	// quantity := strconv.Atoi(params("quantity"))

	var cart models.Cart
	var product models.Product

	if err := models.ReadProductById(controller.Db, &product, productId); err != nil {
		// return c.SendStatus(400) // user salah input productId
		return c.JSON(fiber.Map{
			"message": "product tidak ditemukan",
		})
	}

	cart.UserID = uint(userId)

	if err := models.ReadCartByUserId(controller.Db, &cart, uint(userId)); err != nil {
		// return c.SendStatus(500) //error server
		return c.JSON(fiber.Map{
			"message": "cart tidak ditemukan",
		})
	}
	if cart.ProductIDs == "" {
		cart.ProductIDs = strconv.FormatUint(uint64(product.ID), 10)
	} else {
		cart.ProductIDs = cart.ProductIDs + "; " + strconv.FormatUint(uint64(product.ID), 10)
	}
	if err := models.SaveCart(controller.Db, &cart); err != nil {
		return c.SendStatus(500) //error server
	}
	return c.JSON(fiber.Map{
		"cart": cart,
	})
}
