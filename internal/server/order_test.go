package server

import (
	"context"
	"github.com/stretchr/testify/assert"
	"tempme-book-store/internal/mocks"
	"tempme-book-store/internal/models"
	"testing"
)

var mockClient *mocks.MockClient
var server *EchoServer

func setupMockObj() {
	mockClient = new(mocks.MockClient)
	server = &EchoServer{
		DB: mockClient,
	}

}

func TestValidation(t *testing.T) {
	setupMockObj()
	testUser := &models.User{
		UserID: "testUserId",
		Email:  "testEmail",
	}
	email := "testEmail"
	mockClient.On("GetUserByEmail", context.Background(), "testEmail").Return(testUser, nil)
	ctx := context.Background()
	result, err := server.ValidateUser(ctx, email)
	assert.Equal(t, testUser, result)
	assert.NoError(t, err)

}
