package file

import (
	"encoding/json"
	"http/data/model"
	"io/ioutil"
	"path/filepath"
)

type SuppliersRepository struct {
	directory     string            `json:"directory"`
	suppliersList []*model.Supplier `json:"suppliers_list"`
}

func NewSuppliersRepository() (*SuppliersRepository, error) {
	return &SuppliersRepository{
		directory:     "data/suppliers",
		suppliersList: nil,
	}, nil
}

func (r *SuppliersRepository) GetAll() ([]byte, error) {
	suppliers := make([]model.Supplier, 0)

	files, err := ioutil.ReadDir(r.directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			filepath := filepath.Join(r.directory, file.Name())

			data, err := ioutil.ReadFile(filepath)
			if err != nil {
				return nil, err
			}

			var supplier model.Supplier
			err = json.Unmarshal(data, &supplier)
			if err != nil {
				return nil, err
			}

			suppliers = append(suppliers, supplier)
		}

	}

	jsonDate, err := json.Marshal(suppliers)
	if err != nil {
		return nil, err
	}

	return jsonDate, nil
}
