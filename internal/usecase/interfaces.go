package usecase

import "github.com/pokrovsky-io/msgstore/internal/entity"

type Order interface {
	Create(order *entity.Order) error
	Get(id int) (*entity.Order, error)
}

type OrderRepo interface {
	Create(order *entity.Order) error
	Get(ids ...int) ([]entity.Order, error)
}
