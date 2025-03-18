package entity

import (
	"time"
)

// Coin represents a cryptocurrency entity with its attributes.
type Coin struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" gorm:"unique;not null"`
	Description     string    `json:"description" gorm:"type:text"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	PopularityScore int       `json:"popularity_score" gorm:"default:0"`
}

func NewCoin(name, description string) *Coin {
	return &Coin{
		ID:              0,
		Name:            name,
		Description:     description,
		CreatedAt:       time.Time{},
		PopularityScore: 0,
	}
}
