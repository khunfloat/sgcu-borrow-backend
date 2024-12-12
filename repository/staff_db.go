package repository

import (
	repository "github.com/khunfloat/sgcu-borrow-backend/repository/model"
	"gorm.io/gorm"
)

type staffRepositoryDB struct {
	db *gorm.DB	
}

func NewStaffRepositoryDB(db *gorm.DB) staffRepositoryDB {
	db.AutoMigrate(repository.Staff{})
	return staffRepositoryDB{db: db}
}

func (r staffRepositoryDB) GetAll() ([]repository.Staff, error) {

	staffs := []repository.Staff{}
	
	// query
	tx := r.db.Find(&staffs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return staffs, nil
}

func (r staffRepositoryDB) GetById(id string) (*repository.Staff, error) {

	staff := repository.Staff{}
	tx := r.db.First(&staff, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &staff, nil
}

func (r staffRepositoryDB) Create(id string, name string, password string) (*repository.Staff, error) {

	staff := repository.Staff{
		ID: id,
		Name: name,
		Password: password,
	}

	tx := r.db.Create(&staff)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &staff, nil
}

func (r staffRepositoryDB) Update(id string, name string, password string) (*repository.Staff, error) {

	// Get data
	staff := repository.Staff{}
	tx := r.db.First(&staff, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	
	// Update data
	staff.Name = name
	staff.Password = password

	tx = r.db.Save(&staff)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &staff, nil
}