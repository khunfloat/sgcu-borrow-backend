package repository

type Lost struct {
	ID              int       `gorm:"primaryKey;autoIncrement;not null"`
	OrderId         int       `gorm:"not null"`
	ItemId          int       `gorm:"not null"`
	Amount          int       `gorm:"not null"`
}

type LostRepository interface {
	GetAll() ([]Lost, error)
	GetById(string) (*Lost, error)
	Create(string, string, int) (*Lost, error)
	Update(string, string, string, int) (*Lost, error)
	DeleteById(string) (error)
}