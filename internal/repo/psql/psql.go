package psql

import "github.com/pokrovsky-io/msg-store/internal/entity"

type DB struct {
}

// TODO: Добавить подключение к БД в качестве аргумента
func New() *DB {
	return nil
}

func (c *DB) Create(order *entity.Order) {
}

func (c *DB) Get(id int) (*entity.Order, error) {
	return nil, nil
}

//func (pg *Postgres) SaveOrders(orders ...*entity.Order) {
//
//}
//
//func (pg *Postgres) GetOrder(id int) (*entity.Order, error) {
//	return nil, nil
//}
//
//func (pg *Postgres) RemoveOrders(ids ...int) {
//
//}
//
//func (pg *Postgres) ClearStorage() {
//
//}
