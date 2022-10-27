package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"name/shoppingcart/database"
	"name/shoppingcart/models"
)

type ProductController struct {
	Db *gorm.DB
}

// initial
func InitProductController() *ProductController {
	db := database.InitDb()

	// migrate the schema
	db.AutoMigrate(&models.Product{})

	return &ProductController{Db: db}
}

// routings
// Get products
func (controller *ProductController) IndexProduct(c *fiber.Ctx) error {
	var products []models.Product
	err := models.ReadProducts(controller.Db, &products)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	return c.JSON(fiber.Map{
		"products": products,
	})
}

// POST /products/create
func (controller *ProductController) AddPostedProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.SendStatus(400)
	}
	if form, err := c.MultipartForm(); err == nil {
		files := form.File["picture"]

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			// Save the files to disk:
			fmt.Println(product.Picture)
			if err := c.SaveFile(file, fmt.Sprintf("public/upload/%s", file.Filename)); err != nil {
				return err
			}
			product.Picture = fmt.Sprintf(file.Filename)
		}
	}

	// save product
	err := models.CreateProduct(controller.Db, &product)
	if err != nil {
		return c.SendStatus(500)
	}
	// if succeed
	return c.JSON(product)
}

// GET /products/detail/:id
func (controller *ProductController) GetDetailProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, id)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	return c.JSON(fiber.Map{
		"product": product,
	})
}

// Put product/edit/:id
func (controller *ProductController) EditProductAPI(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, id)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	var updateProduct models.Product

	if form, err := c.MultipartForm(); err == nil {
		files := form.File["picture"]
		fmt.Println(files)
		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			// Save the files to disk:
			fmt.Println(updateProduct.Picture)
			if err := c.SaveFile(file, fmt.Sprintf("public/upload/%s", file.Filename)); err != nil {
				return err
			}
			updateProduct.Picture = fmt.Sprintf(file.Filename)
		}
	}

	if err := c.BodyParser(&updateProduct); err != nil {
		return c.SendStatus(400)
	}
	product.Name = updateProduct.Name
	product.Quantity = updateProduct.Quantity
	product.Price = updateProduct.Price
	product.Picture = updateProduct.Picture

	// save product
	models.UpdateProduct(controller.Db, &product)

	return c.JSON(product)
}

// / GET /products/deleteproduct/xx
func (controller *ProductController) DeleteProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	var product models.Product
	models.DeleteProductById(controller.Db, &product, id)
	return c.JSON(fiber.Map{
		"message": "data was deleted",
	})
}
