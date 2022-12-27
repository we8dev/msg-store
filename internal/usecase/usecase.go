package usecase

import (
	"github.com/pokrovsky-io/msg-store/internal/entity"
)

type OrderUseCase struct {
	cache OrderRepo
	db    OrderRepo
}

func New(cache, db OrderRepo) *OrderUseCase {
	return &OrderUseCase{
		cache: cache,
		db:    db,
	}
}

// TODO Добавить взаимодействия кэша и БД

func (uc *OrderUseCase) Create(order *entity.Order) {
	uc.cache.Create(order)
}

func (uc *OrderUseCase) Get(id int) (*entity.Order, error) {
	return uc.cache.Get(id)
}
