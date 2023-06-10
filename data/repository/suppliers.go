package repository

import "http/data/model"

type SuppliersRepository interface {
	GetAll() ([]*model.Supplier, error)
	GetById() (*model.Supplier, error)
}
