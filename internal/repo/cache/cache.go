package cache

import (
	"errors"
	"github.com/pokrovsky-io/msg-store/internal/entity"
)

var (
	ErrOrderNotFound = errors.New("order not found")
)

type Cache struct {
	orders map[int]*entity.Order
}

func New() *Cache {
	// TODO Установить capacity в зависимости от кол-ва элементов
	return &Cache{
		orders: make(map[int]*entity.Order),
	}
}

func (c *Cache) SaveOrders(orders ...*entity.Order) {
	startId := len(c.orders) + 1

	for i, order := range orders {
		c.orders[startId+i] = order
	}
}

func (c *Cache) GetOrders(ids ...int) ([]*entity.Order, error) {
	res := make([]*entity.Order, 0, len(ids))

	for _, id := range ids {
		order, ok := c.orders[id]
		if !ok {
			return nil, ErrOrderNotFound
		}

		res = append(res, order)
	}

	return res, nil
}

func (c *Cache) RemoveOrders(ids ...int) {
	for _, id := range ids {
		delete(c.orders, id)
	}
}

func (c *Cache) ClearStorage() {
	c.orders = make(map[int]*entity.Order)
}
