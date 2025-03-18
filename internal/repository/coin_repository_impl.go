package repository

import (
	"context"
	"portto/entity"
	"time"

	"gorm.io/gorm"
)

type coinRepositoryImpl struct {
	rw *gorm.DB
}

// NewCoinRepository is a constructor function that creates a new instance of entity.CoinRepository.
func NewCoinRepository(rw *gorm.DB) (entity.CoinRepository, error) {
	err := rw.AutoMigrate(&coinDAO{})
	if err != nil {
		return nil, err
	}

	return &coinRepositoryImpl{
		rw: rw,
	}, nil
}

func (i *coinRepositoryImpl) Create(c context.Context, coin *entity.Coin) error {
	dao := fromEntity(coin)
	if dao.CreatedAt.IsZero() {
		dao.CreatedAt = time.Now()
	}
	if err := i.rw.WithContext(c).Create(dao).Error; err != nil {
		return err
	}
	coin.ID = dao.ID
	coin.CreatedAt = dao.CreatedAt
	return nil
}

func (i *coinRepositoryImpl) GetByID(c context.Context, id uint) (*entity.Coin, error) {
	var dao coinDAO
	if err := i.rw.WithContext(c).First(&dao, id).Error; err != nil {
		return nil, err
	}
	return dao.toEntity(), nil
}

func (i *coinRepositoryImpl) GetByName(c context.Context, name string) (*entity.Coin, error) {
	var dao coinDAO
	if err := i.rw.WithContext(c).Where("name = ?", name).First(&dao).Error; err != nil {
		return nil, err
	}
	return dao.toEntity(), nil
}

func (i *coinRepositoryImpl) List(c context.Context, cond entity.ListCondition) ([]*entity.Coin, int, error) {
	var daos []coinDAO
	var total int64

	if err := i.rw.WithContext(c).Model(&coinDAO{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query := i.rw.WithContext(c).Model(&coinDAO{})
	if cond.Limit > 0 {
		query = query.Limit(cond.Limit)
	}
	if cond.Offset > 0 {
		query = query.Offset(cond.Offset)
	}
	if err := query.Find(&daos).Error; err != nil {
		return nil, 0, err
	}

	coins := make([]*entity.Coin, len(daos))
	for idx, dao := range daos {
		coins[idx] = dao.toEntity()
	}
	return coins, int(total), nil
}

func (i *coinRepositoryImpl) UpdateDescription(c context.Context, id uint, description string) error {
	return i.rw.WithContext(c).
		Model(&coinDAO{}).
		Where("id = ?", id).
		Update("description", description).Error
}

func (i *coinRepositoryImpl) Delete(c context.Context, id uint) error {
	return i.rw.WithContext(c).Delete(&coinDAO{}, id).Error
}

func (i *coinRepositoryImpl) Poke(c context.Context, id uint, score int) error {
	return i.rw.WithContext(c).
		Model(&coinDAO{}).
		Where("id = ?", id).
		UpdateColumn("popularity_score", gorm.Expr("popularity_score + ?", score)).
		Error
}
