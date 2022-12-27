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

func (c *Cache) Create(order *entity.Order) {
	id := len(c.orders) + 1

	c.orders[id] = order
}

func (c *Cache) Get(id int) (*entity.Order, error) {
	order, ok := c.orders[id]
	if !ok {
		return nil, ErrOrderNotFound
	}

	return order, nil
}

//func (c *Cache) SaveOrders(orders ...*entity.Order) {
//	startId := len(c.orders) + 1
//
//	for i, order := range orders {
//		c.orders[startId+i] = order
//	}
//}

//
//func (c *Cache) RemoveOrders(ids ...int) {
//	for _, id := range ids {
//		delete(c.orders, id)
//	}
//}
//
//func (c *Cache) ClearStorage() {
//	c.orders = make(map[int]*entity.Order)
//}
