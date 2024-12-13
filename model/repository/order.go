package model

import "time"

type Order struct {
	ID             int       `gorm:"primaryKey;autoIncrement;not null"`
	UserId         string    `gorm:"not null"`
	UserOrg        string    `gorm:"not null"`
	BorrowDatetime time.Time `gorm:"not null"`
	ReturnDatetime time.Time `gorm:"not null"`
}

type OrderRepository interface {
	GetAll() ([]Order, error)
	GetById(string) (*Order, error)
	Create(string, string, time.Time, time.Time) (*Order, error)
	Update(string, string, string, time.Time, time.Time) (*Order, error)
	DeleteById(string) (error)
}
