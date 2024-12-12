package repository

type Admin struct {
	ID       string `gorm:"primaryKey;size:10;not null"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type AdminRepository interface {
	GetAll() ([]Admin, error)
	GetById(string) (*Admin, error)
	Create(string, string, string) (*Admin, error)
	Update(string, string, string) (*Admin, error)
}