package database

import (
	"context"
	"tempme-book-store/internal/models"
)

func (c Client) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	result := c.DB.WithContext(ctx).Find(&books)
	return books, result.Error
}
