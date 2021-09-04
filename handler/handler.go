package handler

import (
	"test/model"
	"test/service"

	"github.com/gofiber/fiber"
)

type ProductHandler struct {
	ProductService service.ProductService
}

func(r *ProductHandler) GetProducts(c *fiber.Ctx) {
	products, err := r.ProductService.GetAllProducts()

	if err != nil {
		c.Status(403).Send(err)
		return
	}

	c.JSON(products)
}

func(r *ProductHandler) InsertProduct(c *fiber.Ctx) {
	product := &model.Product{}

	if err := c.BodyParser(product); err != nil {
		c.Status(403).Send(err)
		return
	}

	r.ProductService.InsertProduct(*product)
}