package repository

import (
	"Farhan-Backend-POS/database"
	"Farhan-Backend-POS/models"
	"errors"
	"fmt"
	"strings"
)

func AddCategory(name string) (*models.ProductCategory, error) {
	var existingCategory models.ProductCategory

	// Check if category exists
	resultCategory := database.CDB.Where("name = ?", name).First(&existingCategory)
	if resultCategory.Error == nil {
		return nil, errors.New("category already exists")
	}

	// Create new category
	categoryReq := models.ProductCategory{
		Name: name,
	}

	// Save to database
	resultCategory = database.CDB.Create(&categoryReq)
	if resultCategory.Error != nil {
		return nil, fmt.Errorf("failed to create category: %v", resultCategory.Error)
	}

	return &categoryReq, nil
}

// AddProduct - Smart function to add product or update stock if duplicate exists
func AddProduct(name, description string, price float64, stockQuantity int, categoryID uint64, imageURL string) (*models.Product, error) {
	// Normalize input untuk pengecekan duplikasi
	normalizedName := strings.TrimSpace(strings.ToLower(name))
	normalizedDescription := strings.TrimSpace(strings.ToLower(description))

	// Validasi input dasar
	if normalizedName == "" {
		return nil, errors.New("product name cannot be empty")
	}
	if price < 0 {
		return nil, errors.New("product price cannot be negative")
	}
	if stockQuantity < 0 {
		return nil, errors.New("stock quantity cannot be negative")
	}

	// Cek apakah produk sudah ada dengan kriteria:
	// 1. Name sama (case-insensitive)
	// 2. Description sama (case-insensitive)
	// 3. CategoryID sama
	// 4. Price sama (dengan toleransi kecil untuk float)
	var existingProduct models.Product

	// Query untuk mencari produk yang mirip
	result := database.UDB.Where(
		"LOWER(TRIM(name)) = ? AND LOWER(TRIM(description)) = ? AND category_id = ? AND ABS(price - ?) < 0.01",
		normalizedName, normalizedDescription, categoryID, price,
	).First(&existingProduct)

	if result.Error == nil {
		// Produk sudah ada, update stock quantity
		newStockQuantity := existingProduct.StockQuantity + stockQuantity

		updateResult := database.UDB.Model(&existingProduct).Update("stock_quantity", newStockQuantity)
		if updateResult.Error != nil {
			return nil, fmt.Errorf("failed to update stock quantity: %v", updateResult.Error)
		}

		// Refresh data dari database
		database.UDB.First(&existingProduct, existingProduct.ID)

		fmt.Printf("DEBUG: Product already exists. Updated stock from %d to %d\n",
			existingProduct.StockQuantity-stockQuantity, existingProduct.StockQuantity)

		return &existingProduct, nil
	}

	// Produk belum ada, buat produk baru
	newProduct := &models.Product{
		Name:          strings.TrimSpace(name), // Simpan dengan format asli (tidak lowercase)
		Description:   strings.TrimSpace(description),
		Price:         price,
		StockQuantity: stockQuantity,
		CategoryID:    categoryID,
		ImageURL:      strings.TrimSpace(imageURL),
	}

	createResult := database.UDB.Create(&newProduct)
	if createResult.Error != nil {
		return nil, fmt.Errorf("failed to create new product: %v", createResult.Error)
	}

	fmt.Printf("DEBUG: New product created with ID: %d\n", newProduct.ID)
	return newProduct, nil
}
