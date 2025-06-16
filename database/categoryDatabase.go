package database

import (
	"Farhan-Backend-POS/modules/bakery/dto"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var CDB *gorm.DB

func ConnectCategory() {
	connectCategory, err := gorm.Open(postgres.Open("postgresql://postgres:hanhan123@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	CDB = connectCategory
	connectCategory.AutoMigrate(&dto.ProductCategory{})
}
