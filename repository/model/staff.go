package repository

type Staff struct {
	ID       string `gorm:"primaryKey;size:10;not null"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type StaffRepository interface {
	GetAll() ([]Staff, error)
	GetById(string) (*Staff, error)
	Create(string, string, string) (*Staff, error)
	Update(string, string, string) (*Staff, error)
}
