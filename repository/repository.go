package repository

import (
	"test/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func(r *ProductRepository) FindAll() ([]model.Product, error) {
	var products []model.Product
	findResult := r.DB.Find(&products)

	return products, findResult.Error
}

func(r *ProductRepository) FindById(id string) (model.Product, error) {
	return model.Product{}, nil
}

func(r *ProductRepository) Save(product model.Product) (model.Product, error) {
	trx := r.DB.Create(&product)

	return product, trx.Error
}

// func(r *ProductRepository) Update(id string, Product model.Product) {
// 	//
// }

// func(r *ProductRepository) Delete(id string) {
// 	//
// }
