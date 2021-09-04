package main

import (
	"test/handler"

	"test/model"
	"test/repository"
	"test/service"

	"github.com/gofiber/fiber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func helloProduct(c *fiber.Ctx) {
	product := model.Product{
		Model:    gorm.Model{},
		Name:     "Ultramilk",
		Quantity: 2,
		Supplier: "Indofood",
	}
	c.JSON(product)
}

func helloPost(c *fiber.Ctx) {
	product := &model.Product{}

	if err := c.BodyParser(product); err != nil {
		c.Status(403).Send(err)
		return
	}

	c.JSON(product)
}

func main() {
	app := fiber.New()

	app.Get("/", helloProduct)
	app.Post("/", helloPost)

	gormDB, err := gorm.Open(mysql.Open("root:74712331@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"))

	if err != nil {
		panic(err.Error())
	}

	model.MigrateProduct(gormDB)

	productRepo := repository.ProductRepository{
		DB: gormDB,
	}
	productService := service.ProductService{
		ProductRepo: productRepo,
	}
	productHandler := handler.ProductHandler{
		ProductService: productService,
	}

	app.Get("/product", productHandler.GetProducts)
	app.Post("/product", productHandler.InsertProduct)

	app.Listen(3000)
}
