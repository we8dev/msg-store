package usecase

import "github.com/pokrovsky-io/msg-store/internal/model"

type Storage interface {
	SaveOrders(orders ...*model.Order)
	GetOrders(ids ...int) ([]*model.Order, error)
	RemoveOrders(ids ...int)
	ClearStorage()
}
