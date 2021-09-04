package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `json:"name"`
	Quantity int `json:"qty"`
	Supplier string `json:"supplier"`
}

func MigrateProduct(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}