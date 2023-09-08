package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"tempme-book-store/internal/models"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*models.User), args.Error(1)
}

// the following methods to be implemented later
func (m *MockClient) Ready() bool {
	return false
}

func (m *MockClient) AddUser(ctx context.Context, user *models.User) (*models.User, error) {
	return nil, nil
}

func (m *MockClient) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	return nil, nil
}

func (m *MockClient) CreateOrder(ctx context.Context, order *models.Order, items []models.OrderRequestItem) error {
	return nil
}

func (m *MockClient) GetOrdersByUserId(ctx context.Context, user models.User) (*models.UserOrderResponse, error) {
	return nil, nil
}
