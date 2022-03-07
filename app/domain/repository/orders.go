package repository

import "github.com/spro80/apiGolang/app/domain/entity"

type OrdersRepository interface {
	Create(order *entity.Order) (*entity.Order, error)
	Find(id string) (*entity.Order, error)
}
