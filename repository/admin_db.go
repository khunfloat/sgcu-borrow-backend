package repository

import (
	modelRepo "github.com/khunfloat/sgcu-borrow-backend/model/repository"
	"gorm.io/gorm"
)

type adminRepositoryDB struct {
	db *gorm.DB
}

func NewAdminRepositoryDB(db *gorm.DB) adminRepositoryDB {
	db.AutoMigrate(modelRepo.Admin{})
	return adminRepositoryDB{db: db}
}

func (r adminRepositoryDB) GetAll() ([]modelRepo.Admin, error) {

	admins := []modelRepo.Admin{}
	
	// query
	tx := r.db.Find(&admins)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return admins, nil
}

func (r adminRepositoryDB) GetById(id string) (*modelRepo.Admin, error) {

	admin := modelRepo.Admin{}
	tx := r.db.First(&admin, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &admin, nil
}

func (r adminRepositoryDB) Create(id string, name string, password string) (*modelRepo.Admin, error) {

	admin := modelRepo.Admin{
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

func (r adminRepositoryDB) Update(id string, name string, password string, banStatus int) (*modelRepo.Admin, error) {

	// Get data
	admin := modelRepo.Admin{}
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