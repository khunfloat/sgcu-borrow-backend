package model

type Return struct {
	ID              int       `gorm:"primaryKey;autoIncrement;not null"`
	OrderId         int       `gorm:"not null"`
	ItemId          int       `gorm:"not null"`
	Amount          int       `gorm:"not null"`
}

type ReturnRepository interface {
	GetAll() ([]Return, error)
	GetById(int) (*Return, error)
	Create(int, int, int) (*Return, error)
	Update(int, int, int, int) (*Return, error)
	DeleteById(int) (error)
}