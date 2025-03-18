package repository

import (
	"portto/entity"
	"time"
)

type coinDAO struct {
	ID              uint      `gorm:"primaryKey"`
	Name            string    `gorm:"not null;unique"`
	Description     string    `gorm:"type:text"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	PopularityScore int       `gorm:"default:0"`
}

func fromEntity(c *entity.Coin) *coinDAO {
	return &coinDAO{
		ID:              c.ID,
		Name:            c.Name,
		Description:     c.Description,
		CreatedAt:       c.CreatedAt,
		PopularityScore: c.PopularityScore,
	}
}

func (dao *coinDAO) toEntity() *entity.Coin {
	return &entity.Coin{
		ID:              dao.ID,
		Name:            dao.Name,
		Description:     dao.Description,
		CreatedAt:       dao.CreatedAt,
		PopularityScore: dao.PopularityScore,
	}
}
