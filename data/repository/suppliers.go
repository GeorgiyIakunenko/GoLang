package repository

type SuppliersRepository interface {
	GetAll() ([]byte, error)
}
