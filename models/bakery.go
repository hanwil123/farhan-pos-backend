package models

type Product struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
	CategoryID    string  `json:"category_id"`
	ImageURL      string  `json:"image_url"`
}

type ProductCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Customer struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type Cart struct {
	ID         string     `json:"id"`
	CustomerID string     `json:"customer_id"`
	Status     string     `json:"status"`
	Items      []CartItem `json:"items"`
}

type CartItem struct {
	ID        string  `json:"id"`
	CartID    string  `json:"cart_id"`
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}

type Purchase struct {
	ID          string         `json:"id"`
	CustomerID  string         `json:"customer_id"`
	TotalAmount float64        `json:"total_amount"`
	PaymentType string         `json:"payment_type"`
	PurchasedAt string         `json:"purchased_at"`
	Items       []PurchaseItem `json:"items"`
}

type PurchaseItem struct {
	ID         string  `json:"id"`
	PurchaseID string  `json:"purchase_id"`
	ProductID  string  `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
	Subtotal   float64 `json:"subtotal"`
}

type StockLog struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	Change    int    `json:"change"`
	Reason    string `json:"reason"`
}
