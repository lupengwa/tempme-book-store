package server

import (
	"context"
	"net/http"
	"tempme-book-store/internal/models"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) CreateOrder(ctx echo.Context) error {
	user := ctx.Request().Header[userEmail]
	if user == nil {
		return ctx.JSON(http.StatusUnauthorized, unAuthorized)
	}
	userDB, err := s.ValidateUser(ctx.Request().Context(), user[0])
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, unAuthorized)
	}
	if userDB == nil {
		return ctx.JSON(http.StatusUnauthorized, unAuthorized)
	}

	orderRequest := new(models.OrderRequest)
	if err := ctx.Bind(orderRequest); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	orderEntity := &models.Order{UserID: userDB.UserID}

	err = s.DB.CreateOrder(ctx.Request().Context(), orderEntity, orderRequest.Items)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, orderRequest)
}

func (s *EchoServer) GetOrdersByUserId(ctx echo.Context) error {
	user := ctx.Request().Header[userEmail]
	if user == nil {
		return ctx.JSON(http.StatusUnauthorized, unAuthorized)
	}
	userDB, err := s.ValidateUser(ctx.Request().Context(), user[0])
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, unAuthorized)
	}
	if userDB == nil {
		return ctx.JSON(http.StatusUnauthorized, unAuthorized)
	}

	resp, err := s.DB.GetOrdersByUserId(ctx.Request().Context(), *userDB)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *EchoServer) ValidateUser(ctx context.Context, userEmail string) (*models.User, error) {
	result, err := s.DB.GetUserByEmail(ctx, userEmail)
	if err != nil {
		return nil, err
	}
	return result, nil
}
