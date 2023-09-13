package models

type Invoices struct {
	Id     int `gorm:"primaryKey"`
	UserId int
	Status string
	Label  string
	Amount float64
}
