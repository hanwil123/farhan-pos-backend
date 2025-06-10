// user service - FIXED VERSION
package handler

import (
	"Farhan-Backend-POS/proto"
	"Farhan-Backend-POS/repository"
	"context"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserServiceServer struct {
	proto.UnimplementedUserServiceServer
}

const SecretKey = "secret"

func (s *UserServiceServer) RegisterUser(ctx context.Context, req *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {
	user, err := repository.RegisterUser(req.Name, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &proto.RegisterUserResponse{
		Id:                         strconv.FormatUint(user.Id, 10),
		Name:                       user.Name,
		Email:                      user.Email,
		Message:                    "Register Successfully",
		StatusCodeBerhasilRegister: strconv.Itoa(200),
	}, nil
}

// FIXED: Hapus parameter fiber.Ctx dan duplikasi validasi password
func (s *UserServiceServer) LoginUser(ctx context.Context, req *proto.LoginUserRequest) (*proto.LoginUserResponse, error) {
	// Repository sudah handle validasi password, jadi tidak perlu duplikasi
	loginUser, err := repository.LoginUser(req.Email, req.Password)
	if err != nil {
		return &proto.LoginUserResponse{
			Message:                 "Login failed: " + err.Error(),
			StatusCodeBerhasilLogin: strconv.Itoa(400),
		}, nil
	}

	// Generate JWT token
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.FormatUint(loginUser.Id, 10),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, errToken := claim.SignedString([]byte(SecretKey))
	if errToken != nil {
		return &proto.LoginUserResponse{
			Message:                 "Token generation failed",
			StatusCodeBerhasilLogin: strconv.Itoa(500),
		}, nil
	}

	return &proto.LoginUserResponse{
		Id:                      strconv.FormatUint(loginUser.Id, 10),
		Token:                   token,
		Message:                 "Login Successfully",
		StatusCodeBerhasilLogin: strconv.Itoa(200),
	}, nil
}
