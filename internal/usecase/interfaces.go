package usecase

import "github.com/pokrovsky-io/msg-store/internal/entity"

type Order interface {
	Create(order *entity.Order)
	Get(id int) (*entity.Order, error)
}

type OrderRepo interface {
	Create(order *entity.Order)
	Get(id int) (*entity.Order, error)
}
