package database

import (
	"Farhan-Backend-POS/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PDB *gorm.DB

func ConnectProduct() {
	connectProduct, err := gorm.Open(postgres.Open("postgresql://postgres:hanhan123@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	PDB = connectProduct
	connectProduct.AutoMigrate(&models.Product{})
}
