package models

// api map object
type UserOrderResponse struct {
	Orders    []OrderResponse `json:"orders"`
	UserEmail string          `json:"user"`
}

type OrderResponse struct {
	OrderId    string              `json:"orderId"`
	OrderItems []OrderItemResponse `json:"orderItems"`
}

type OrderItemResponse struct {
	BookId   string `json:"bookId"`
	Quantity int    `json:"quantity"`
}

type OrderRequest struct {
	Items []OrderRequestItem `json:"items"`
}
type OrderRequestItem struct {
	BookID   string `json:"bookId"`
	Quantity int    `json:"quantity"`
}

// db map object
type OrderItem struct {
	OrderId  string `gorm:"primaryKey" json:"orderId"`
	BookId   string `gorm:"column:book_id" json:"bookId"`
	Quantity int    `json:"quantity"`
}

type Order struct {
	OrderID string `gorm:"primaryKey""`
	UserID  string `gorm:"column:user_id""`
}
