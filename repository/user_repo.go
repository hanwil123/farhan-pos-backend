package repository

import (
	"Farhan-Backend-POS/database"
	"Farhan-Backend-POS/models"
	"errors"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

var users = []models.UserModel{
	// {
	// 	Id:       "1",
	// 	Name:     "John Doe",
	// 	Email:    "john.doe@example.com",
	// 	Password: "password123",
	// },
	// {
	// 	Id:       "2",
	// 	Name:     "Jane Doe",
	// 	Email:    "jane.doe@example.com",
	// 	Password: "password123",
	// },
}

func RegisterUser(name, email, password string) (*models.UserModel, error) {
	fmt.Println("DEBUG: name =", name) // Tambahkan ini
	for _, u := range users {
		if u.Email == email {
			return nil, errors.New("email already exist")
		}
	}
	passwordBcrypt, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user := &models.UserModel{
		Id:       uint64(rand.Uint32()),
		Name:     name,
		Email:    email,
		Password: passwordBcrypt,
	}
	users = append(users, *user)
	fmt.Printf("DEBUG: user struct = %+v\n", user) // Tambahkan ini
	database.UDB.Create(&user)
	return user, nil
}

func LoginUser(email, password string) (models.UserModel, error) {
	for _, u := range users {
		if u.Email == email {
			err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
			if err == nil {
				return u, nil
			}
		}
	}
	return models.UserModel{}, errors.New("user not found")
}

func GetUser(id uint64) (models.UserModel, error) {
	for _, u := range users {
		if u.Id == id {
			return u, nil
		}
	}
	return models.UserModel{}, errors.New("user not found")
}
