package model

type Lost struct {
	ID              int       `gorm:"primaryKey;autoIncrement;not null"`
	OrderId         int       `gorm:"not null"`
	ItemId          int       `gorm:"not null"`
	Amount          int       `gorm:"not null"`
}

type LostRepository interface {
	GetAll() ([]Lost, error)
	GetById(int) (*Lost, error)
	Create(int, int, int) (*Lost, error)
	Update(int, int, int, int) (*Lost, error)
	DeleteById(int) (error)
}