package model

import "time"

type Borrow struct {
	ID             int       `gorm:"primaryKey;autoIncrement;not null"`
	OrderId        int       `gorm:"not null"`
	ItemId         int       `gorm:"not null"`
	Amount         int       `gorm:"not null"`
	PickupDatetime time.Time `gorm:"default null"`
}

type BorrowRepository interface {
	GetAll() ([]Borrow, error)
	GetById(string) (*Borrow, error)
	Create(string, string, int) (*Borrow, error)
	Update(string, string, string, int, time.Time) (*Borrow, error)
	DeleteById(string) error
}
