package usecase

import "github.com/pokrovsky-io/msg-store/internal/entity"

type Order interface {
	Create(orders ...*entity.Order)
}

type OrderRepo interface {
	SaveOrders(orders ...*entity.Order)
	GetOrders(ids ...int) ([]*entity.Order, error)
	RemoveOrders(ids ...int)
	ClearStorage()
}
