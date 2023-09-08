package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

const userEmail = "User"
const unAuthorized = "UnAuthorized"

func (s *EchoServer) GetAllBooks(ctx echo.Context) error {
	books, err := s.DB.GetAllBooks(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, books)
}
