package handler

import (
	"Farhan-Backend-POS/proto"
	"Farhan-Backend-POS/repository"
	"context"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
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

func (s *UserServiceServer) LoginUser(ctx context.Context, req *proto.LoginUserRequest) (*proto.LoginUserResponse, error) {
	loginUser, err := repository.LoginUser(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	errBcrypt := bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(req.Password))
	if errBcrypt != nil {
		return nil, errBcrypt
	}
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.FormatUint(loginUser.Id, 10),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, errToken := claim.SignedString([]byte(SecretKey))
	if errToken != nil {
		return nil, errToken
	}
	return &proto.LoginUserResponse{
		Token:                   token,
		StatusCodeBerhasilLogin: strconv.Itoa(200),
	}, nil
}
