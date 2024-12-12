package repository

type User struct {
	ID        string `gorm:"primaryKey;size:10;not null"`
	Name      string `gorm:"not null"`
	Tel       string `gorm:"not null"`
	Password  string `gorm:"not null"`
	BanStatus int    `gorm:"default 0"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(string) (*User, error)
	Create(string, string, string, string) (*User, error)
	Update(string, string, string, string, int) (*User, error)
}
