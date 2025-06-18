// user repo - FIXED VERSION
package repository

import (
	"Farhan-Backend-POS/database"
	"Farhan-Backend-POS/modules/auth/dto"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var users = []dto.User{
	// Hapus slice lokal ini karena kita menggunakan database
}

func RegisterUser(name, email, password, role string) (*dto.User, error) {
	fmt.Println("DEBUG: name =", name)

	// Check di database, bukan di slice lokal
	var existingUser dto.User
	result := database.UDB.Where("email = ?", email).First(&existingUser)
	if result.Error == nil {
		return nil, errors.New("email already exist")
	}

	passwordBcrypt, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	user := &dto.User{
		Name:     name,
		Email:    email,
		Password: passwordBcrypt,
		Role:     role,
	}

	// Simpan ke database
	result = database.UDB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Printf("DEBUG: user struct = %+v\n", user)
	return user, nil
}

func LoginUser(email, password string) (dto.User, error) {
	var user dto.User
	result := database.UDB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return dto.User{}, errors.New("user not found")
	}

	// Validasi password
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return dto.User{}, errors.New("invalid password")
	}

	return user, nil
}

func GetUser(id uint64) (dto.User, error) {
	var user dto.User
	result := database.UDB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return dto.User{}, errors.New("user not found")
	}
	return user, nil
}
