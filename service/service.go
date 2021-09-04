package service

import (
	"test/model"
	"test/repository"
)

type ProductService struct {
	ProductRepo repository.ProductRepository
}

func(s *ProductService) GetAllProducts() ([]model.Product, error) {
	products, err := s.ProductRepo.FindAll()

	return products, err
}

func(s *ProductService) InsertProduct(product model.Product) error {
	_, err := s.ProductRepo.Save(product)

	return err
}