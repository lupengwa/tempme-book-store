package models

type Book struct {
	BookID string  `gorm:"primaryKey" json:"bookId"`
	Name   string  `json:"name"`
	Price  float32 `gorm:"type:numeric" json:"price"`
}
