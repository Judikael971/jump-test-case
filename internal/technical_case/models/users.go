package models

type Users struct {
	Id       int    `gorm:"primaryKey"`
	FistName string `gorm:"column:first_name"`
	LastName string `gorm:"column:last_name"`
	Balance  float64
}
