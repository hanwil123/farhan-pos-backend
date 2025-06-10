package database

import (
	"Farhan-Backend-POS/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var UDB *gorm.DB

func ConnectUser() {
	connectUser, err := gorm.Open(postgres.Open("postgresql://postgres:hanhan123@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	UDB = connectUser
	connectUser.AutoMigrate(&models.UserModel{})
}
