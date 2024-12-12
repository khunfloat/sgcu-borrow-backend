package repository

import "time"

type Return struct {
	ID              int       `gorm:"primaryKey;autoIncrement;not null"`
	OrderId         int       `gorm:"not null"`
	ItemId          int       `gorm:"not null"`
	Amount          int       `gorm:"not null"`
	DropoffDatetime time.Time `gorm:"default null"`
}

type ReturnRepository interface {
	GetAll() ([]Return, error)
	GetById(string) (*Return, error)
	Create(string, string, int) (*Return, error)
	Update(string, string, string, int, time.Time) (*Return, error)
	DeleteById(string) (error)
}