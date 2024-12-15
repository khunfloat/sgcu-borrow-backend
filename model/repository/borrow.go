package model

type Borrow struct {
	ID             int       `gorm:"primaryKey;autoIncrement;not null"`
	OrderId        int       `gorm:"not null"`
	ItemId         int       `gorm:"not null"`
	Amount         int       `gorm:"not null"`
}

type BorrowRepository interface {
	GetAll() ([]Borrow, error)
	GetById(int) (*Borrow, error)
	GetByOrderId(int) ([]Borrow, error)
	GetByOrderIdAndItemId(int, int) (*Borrow, error)
	Create(int, int, int) (*Borrow, error)
	Update(int, int, int, int) (*Borrow, error)
	DeleteById(int) error
}
