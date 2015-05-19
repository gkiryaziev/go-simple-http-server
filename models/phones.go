package models

type Phones struct {
	Id int				`json:"id"`
	Phone string		`json:"phone"`
	Person string		`json:"person"`
	Street string		`json:"street"`
	Building string		`json:"building"`
	Appartment string	`json:"appartment"`
}
