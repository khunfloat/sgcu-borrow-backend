package model

import "time"

type Order struct {
	ID              int       `gorm:"primaryKey;autoIncrement;not null"`
	UserId          string    `gorm:"not null"`
	UserOrg         string    `gorm:"not null"`
	BorrowDatetime  time.Time `gorm:"not null"`
	ReturnDatetime  time.Time `gorm:"not null"`
	PickupDatetime  time.Time `gorm:"default null"`
	DropoffDatetime time.Time `gorm:"default null"`
}

type OrderRepository interface {
	GetAll() ([]Order, error)
	GetById(int) (*Order, error)
	Create(string, string, time.Time, time.Time) (*Order, error)
	UpdateInfo(int, string, string, time.Time, time.Time) (*Order, error)
	UpdatePickupDatetime(int, time.Time) (*Order, error)
	UpdateDropoffDatetime(int, time.Time) (*Order, error)
	DeleteById(int) error
}
