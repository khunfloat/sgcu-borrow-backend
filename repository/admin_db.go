package repository

import (
	repository "github.com/khunfloat/sgcu-borrow-backend/repository/model"
	"gorm.io/gorm"
)

type adminRepositoryDB struct {
	db *gorm.DB
}

func NewAdminRepositoryDB(db *gorm.DB) adminRepositoryDB {
	db.AutoMigrate(repository.Admin{})
	return adminRepositoryDB{db: db}
}

func (r adminRepositoryDB) GetAll() ([]repository.Admin, error) {

	admins := []repository.Admin{}
	
	// query
	tx := r.db.Find(&admins)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return admins, nil
}

func (r adminRepositoryDB) GetById(id string) (*repository.Admin, error) {

	admin := repository.Admin{}
	tx := r.db.First(&admin, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &admin, nil
}

func (r adminRepositoryDB) Create(id string, name string, password string) (*repository.Admin, error) {

	admin := repository.Admin{
		ID: id,
		Name: name,
		Password: password,
	}

	tx := r.db.Create(&admin)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &admin, nil
}

func (r adminRepositoryDB) Update(id string, name string, password string, banStatus int) (*repository.Admin, error) {

	// Get data
	admin := repository.Admin{}
	tx := r.db.First(&admin, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	
	// Update data
	admin.Name = name
	admin.Password = password

	tx = r.db.Save(&admin)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &admin, nil
}