package repository

import "http/data/model"

type SuppliersRepository interface {
	GetAll() ([]*model.Supplier, error)
	GetById(id int) (model.Supplier, error)
	UpdateSupplierById(id int, newSup model.Supplier) error
	RefreshSuppliers(data []*model.Supplier) error
}
