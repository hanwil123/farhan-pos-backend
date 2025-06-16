package handler

import (
	"Farhan-Backend-POS/proto"
	"Farhan-Backend-POS/repository"
	"context"
	"fmt"
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

func (s *BakeryProductServiceServer) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) (*proto.ProductResponse, error) {
	categoryId, err := strconv.ParseUint(req.CategoryId, 10, 64)
	if err != nil {
		return nil, err
	}

	product, err := repository.AddProduct(req.Name, req.Description, req.Price, int(req.StockQuantity), categoryId, req.ImageUrl)
	if err != nil {
		return nil, err
	}

	return &proto.ProductResponse{
		Product: &proto.Product{
			Id:            strconv.FormatUint(product.ID, 10),
			Name:          product.Name,
			Description:   product.Description,
			Price:         product.Price,
			StockQuantity: int32(product.StockQuantity),
			CategoryId:    strconv.FormatUint(product.CategoryID, 10),
			ImageUrl:      product.ImageURL,
		},
	}, nil
}

func (s *BakeryProductServiceServer) GetCategoryById(ctx context.Context, req *proto.GetCategoryByIdRequest) (*proto.CategoryResponse, error) {
	categoryId, err := strconv.ParseUint(req.Id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid category id: %v", err)
	}

	category, err := repository.GetCategoryById(categoryId)
	if err != nil {
		return nil, err
	}

	return &proto.CategoryResponse{
		Id:   strconv.FormatUint(category.ID, 10),
		Name: category.Name,
	}, nil
}

func (s *BakeryProductServiceServer) ListCategories(ctx context.Context, req *proto.Empty) (*proto.CategoryList, error) {
	cats, err := repository.ListCategories()
	if err != nil {
		return nil, err
	}

	var protoCategories []*proto.CategoryResponse
	for _, cat := range cats {
		protoCategories = append(protoCategories, &proto.CategoryResponse{
			Id:   strconv.FormatUint(cat.ID, 10),
			Name: cat.Name,
		})
	}

	return &proto.CategoryList{
		Categories: protoCategories,
	}, nil
}
