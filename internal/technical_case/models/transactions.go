package models

type Transactions struct {
	Id        int `gorm:"primaryKey"`
	InvoiceId int
	Status    string
	Reference string
	Amount    float64
}
