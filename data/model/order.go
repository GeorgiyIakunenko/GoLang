package model

type Address struct {
	Street   string `json:"street"`
	Building int    `json:"building"`
	Apt      int    `json:"apt"`
}

type Order struct {
	UserId     int     `json:"user_id"`
	Address    Address `json:"address"`
	SupplierId int     `json:"supplier_id"`
	Status     string  `json:"status"`
}
