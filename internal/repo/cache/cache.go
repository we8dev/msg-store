package cache

import (
	"errors"
	"github.com/pokrovsky-io/msg-store/internal/entity"
)

const (
	_defaultCapacity = 100
)

var (
	ErrOrderNotFound = errors.New("order not found")
)

type Cache struct {
	// TODO Указатели или сами сущности?
	orders []*entity.Order
}

func New() *Cache {
	// TODO Установить capacity в зависимости от кол-ва элементов
	return &Cache{
		orders: make([]*entity.Order, 0, _defaultCapacity),
	}
}

func (c *Cache) Create(order *entity.Order) {
	c.orders = append(c.orders, order)
}

func (c *Cache) Get(id int) (*entity.Order, error) {
	if id < len(c.orders) {
		return c.orders[id], nil
	}

	return nil, ErrOrderNotFound
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
