package model

type Item struct {
	ID            int    `gorm:"primaryKey;autoIncrement;not null"`
	Name          string `gorm:"not null"`
	CurrentAmount int    `gorm:"not null"`
	ImgUrl        string `gorm:"not null"`
	BorrowCount   int    `gorm:"default 0"`
}

type ItemRepository interface {
	GetAll() ([]Item, error)
	GetFrequentlyBorrowed() ([]Item, error)
	GetById(int) (*Item, error)
	Create(string, int, string) (*Item, error)
	Update(int, string, int, string, int) (*Item, error)
	DeleteById(int) (error)
}
