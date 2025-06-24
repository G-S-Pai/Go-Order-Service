// models/order.go
package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID        string      `gorm:"type:STRING(36);primaryKey" json:"id"`
	UserID    string      `json:"user_id"`
	Amount    float64     `json:"amount"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate a new UUID and assign it to the ID field
	order.ID = uuid.New().String()
	return nil
}