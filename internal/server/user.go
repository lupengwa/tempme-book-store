package server

import (
	"net/http"
	"tempme-book-store/internal/dberrors"
	"tempme-book-store/internal/models"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) AddUser(ctx echo.Context) error {
	user := new(models.User)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	user, err := s.DB.AddUser(ctx.Request().Context(), user)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, user)
}
