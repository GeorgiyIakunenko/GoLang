package repository

import "http/data/model"

type OrdersRepository interface {
	Create(orderData model.Order) error
}
