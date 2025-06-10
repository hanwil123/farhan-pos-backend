package handler

import (
	"Farhan-Backend-POS/proto"
	"Farhan-Backend-POS/repository"
	"context"
	"strconv"
)

type BakeryProductServiceServer struct {
	proto.UnimplementedBakeryPOSServiceServer
}

func (s *BakeryProductServiceServer) CreateCategory(ctx context.Context, req *proto.CategoryRequest) (*proto.CategoryResponse, error) {
	categoryCreate, err := repository.AddCategory(req.Name)
	if err != nil {
		return nil, err
	}
	return &proto.CategoryResponse{
		Id:   strconv.FormatUint(categoryCreate.ID, 10),
		Name: req.Name,
	}, nil
}
