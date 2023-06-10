package model

type Address struct {
	Street   string `json:"street"`
	Building int    `json:"building"`
	Apt      int    `json:"apt"`
}

type Order struct {
	Id         int     `json:"id"`
	SupplierId int     `json:"supplier_id"`
	UserId     int     `json:"user_id"`
	Address    Address `json:"address"`
	Status     string  `json:"status"`
}
