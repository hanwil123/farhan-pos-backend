package repository

import (
	"Farhan-Backend-POS/database"
	"Farhan-Backend-POS/models"
	"errors"
	"fmt"
	"strings"
)

func AddCategory(name string) (*models.ProductCategory, error) {
	if database.CDB == nil {
		return nil, errors.New("category database connection is not initialized")
	}

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

func ListCategories() ([]models.ProductCategory, error) {
	fmt.Println("DEBUG: Masuk ke repository.ListCategories")
	if database.CDB == nil {
		fmt.Println("ERROR: Koneksi database kategori tidak diinisialisasi.")
		return nil, errors.New("category database connection is not initialized")
	}

	var categories []models.ProductCategory
	fmt.Println("DEBUG: Melakukan query database untuk kategori...")
	result := database.CDB.Find(&categories)
	if result.Error != nil {
		fmt.Printf("ERROR: Gagal mendapatkan kategori dari database: %v\n", result.Error)
		return nil, fmt.Errorf("failed to get categories: %v", result.Error)
	}
	fmt.Printf("DEBUG: Berhasil mendapatkan %d kategori dari database.\n", len(categories))
	return categories, nil
}

// AddProduct - Smart function to add product or update stock if duplicate exists
func AddProduct(name, description string, price float64, stockQuantity int, categoryID uint64, imageURL string) (*models.Product, error) {
	if database.PDB == nil {
		return nil, errors.New("product database connection is not initialized")
	}

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
	result := database.PDB.Where(
		"LOWER(TRIM(name)) = ? AND LOWER(TRIM(description)) = ? AND category_id = ? AND ABS(price - ?) < 0.01",
		normalizedName, normalizedDescription, categoryID, price,
	).First(&existingProduct)

	if result.Error == nil {
		// Produk sudah ada, update stock quantity
		newStockQuantity := existingProduct.StockQuantity + stockQuantity

		updateResult := database.PDB.Model(&existingProduct).Update("stock_quantity", newStockQuantity)
		if updateResult.Error != nil {
			return nil, fmt.Errorf("failed to update stock quantity: %v", updateResult.Error)
		}

		// Refresh data dari database
		database.PDB.First(&existingProduct, existingProduct.ID)

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

	createResult := database.PDB.Create(&newProduct)
	if createResult.Error != nil {
		return nil, fmt.Errorf("failed to create new product: %v", createResult.Error)
	}

	fmt.Printf("DEBUG: New product created with ID: %d\n", newProduct.ID)
	return newProduct, nil
}

func GetCategoryById(id uint64) (*models.ProductCategory, error) {
	if database.CDB == nil {
		return nil, errors.New("category database connection is not initialized")
	}

	var category models.ProductCategory

	result := database.CDB.First(&category, id)
	if result.Error != nil {
		return nil, fmt.Errorf("category not found: %v", result.Error)
	}

	return &category, nil
}
