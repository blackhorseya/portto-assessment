//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package entity

import (
	"context"
)

type ListCondition struct {
	Limit  int
	Offset int
}

type CoinRepository interface {
	Create(c context.Context, coin *Coin) error
	GetByID(c context.Context, id uint) (*Coin, error)
	GetByName(c context.Context, name string) (*Coin, error)
	List(c context.Context, cond ListCondition) ([]*Coin, int, error)
	UpdateDescription(c context.Context, id uint, description string) error
	Delete(c context.Context, id uint) error
	Poke(c context.Context, id uint, score int) error
}
