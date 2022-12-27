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

func (uc *OrderUseCase) Create(orders ...*entity.Order) {

}
