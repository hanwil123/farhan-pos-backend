package grpcServiceBakery

import (
	"Farhan-Backend-POS/modules/bakery/repository"
	"Farhan-Backend-POS/proto"
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
	categoryId, err := strconv.ParseUint(req.Product.CategoryId, 10, 64)
	if err != nil {
		return nil, err
	}

	product, err := repository.AddProduct(req.Product.Name, req.Product.Description, req.Product.Price, int(req.Product.StockQuantity), categoryId, req.Product.ImageUrl)
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

func (s *BakeryProductServiceServer) UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest) (*proto.ProductResponse, error) {
	fmt.Printf("DEBUG: Received Product ID for update: %s\n", req.Product.Id)
	id, err := strconv.ParseUint(req.Product.Id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse product ID: %v", err)
	}

	fmt.Printf("DEBUG: Received Category ID for update: %s\n", req.Product.CategoryId)
	categoryId, err := strconv.ParseUint(req.Product.CategoryId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse category ID: %v", err)
	}

	product, err := repository.UpdateProductRepo(id, req.Product.Name, req.Product.Description, req.Product.Price, int(req.Product.StockQuantity), categoryId, req.Product.ImageUrl)
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

func (s *BakeryProductServiceServer) ListProducts(ctx context.Context, req *proto.Empty) (*proto.ProductListResponse, error) {
	products, err := repository.ListProductsRepo()
	if err != nil {
		return nil, err
	}

	var protoProducts []*proto.Product
	for _, prd := range products {
		protoProducts = append(protoProducts, &proto.Product{
			Id:            strconv.FormatUint(prd.ID, 10),
			Name:          prd.Name,
			Description:   prd.Description,
			Price:         prd.Price,
			StockQuantity: int32(prd.StockQuantity),
			CategoryId:    strconv.FormatUint(prd.CategoryID, 10),
			ImageUrl:      prd.ImageURL,
		})
	}

	return &proto.ProductListResponse{
		Products: protoProducts,
	}, nil
}

func (s *BakeryProductServiceServer) DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) (*proto.DeleteProductResponse, error) {
	id, err := strconv.ParseUint(req.Id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse product ID for deletion: %v", err)
	}

	_, err = repository.DeleteProductRepo(id)
	if err != nil {
		return nil, err
	}
	return &proto.DeleteProductResponse{
		MessageSuccesfull: "successfull delete product",
	}, nil
}
