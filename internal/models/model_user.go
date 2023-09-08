package models

type User struct {
	UserID string `gorm:"primaryKey" json:"userId"`
	Email  string `json:"email"`
}
