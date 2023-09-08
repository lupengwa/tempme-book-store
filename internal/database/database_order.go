package database

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"tempme-book-store/internal/models"
)

func (c Client) CreateOrder(ctx context.Context, order *models.Order, items []models.OrderRequestItem) error {
	order.OrderID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&order)
	if result.Error != nil {
		return result.Error
	}
	orderItems := make([]models.OrderItem, 0)
	for _, item := range items {
		orderItem := models.OrderItem{
			OrderId:  order.OrderID,
			BookId:   item.BookID,
			Quantity: item.Quantity,
		}
		orderItems = append(orderItems, orderItem)
	}
	result = c.DB.WithContext(ctx).Create(&orderItems)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c Client) GetOrdersByUserId(ctx context.Context, user models.User) (*models.UserOrderResponse, error) {
	userOrderResponse := &models.UserOrderResponse{UserEmail: user.Email, Orders: make([]models.OrderResponse, 0)}
	// fetch related orders
	var orders []models.Order
	result := c.DB.WithContext(ctx).Where(&models.Order{UserID: user.UserID}).Find(&orders)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	//fetch related items for each order
	for _, order := range orders {
		var orderItems []models.OrderItem
		result = c.DB.WithContext(ctx).Where(&models.OrderItem{OrderId: order.OrderID}).Find(&orderItems)
		if result.Error != nil || len(orderItems) == 0 {
			continue
		}
		orderResponse := models.OrderResponse{
			OrderId: order.OrderID,
		}
		orderItemsResponse := make([]models.OrderItemResponse, 0)
		for _, orderItem := range orderItems {
			orderItemsResponse = append(orderItemsResponse, models.OrderItemResponse{
				BookId:   orderItem.BookId,
				Quantity: orderItem.Quantity,
			})
		}
		orderResponse.OrderItems = orderItemsResponse
		userOrderResponse.Orders = append(userOrderResponse.Orders, orderResponse)
	}

	return userOrderResponse, nil
}
