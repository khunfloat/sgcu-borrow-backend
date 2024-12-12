package repository

import (
	repository "github.com/khunfloat/sgcu-borrow-backend/repository/model"
	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) userRepositoryDB {
	db.AutoMigrate(repository.User{})
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) GetAll() ([]repository.User, error) {

	users := []repository.User{}
	
	// query
	tx := r.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

func (r userRepositoryDB) GetById(id string) (*repository.User, error) {

	user := repository.User{}
	tx := r.db.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (r userRepositoryDB) Create(id string, name string, tel string, password string) (*repository.User, error) {

	user := repository.User{
		ID: id,
		Name: name,
		Tel: tel,
		Password: password,
	}

	tx := r.db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r userRepositoryDB) Update(id string, name string, tel string, password string, banStatus int) (*repository.User, error) {

	// Get data
	user := repository.User{}
	tx := r.db.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	
	// Update data
	user.Name = name
	user.Tel = tel
	user.Password = password
	user.BanStatus = banStatus

	tx = r.db.Save(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}