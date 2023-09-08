package database

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"tempme-book-store/internal/dberrors"
	"tempme-book-store/internal/models"

	"gorm.io/gorm"
)

func (c Client) AddUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.UserID = uuid.NewString()
	result := c.DB.WithContext(ctx).
		Create(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return user, nil
}

func (c Client) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	result := c.DB.WithContext(ctx).
		Where(&models.User{Email: email}).
		First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}
