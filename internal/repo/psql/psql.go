package psql

import "github.com/pokrovsky-io/msg-store/internal/entity"

type Postgres struct {
}

// TODO: Добавить подключение к БД в качестве аргумента
func New() *Postgres {
	return nil
}

func (pg *Postgres) SaveOrders(orders ...*entity.Order) {

}

func (pg *Postgres) GetOrders(ids ...int) ([]*entity.Order, error) {
	return nil, nil
}

func (pg *Postgres) RemoveOrders(ids ...int) {

}

func (pg *Postgres) ClearStorage() {

}
