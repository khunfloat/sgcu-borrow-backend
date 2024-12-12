package repository

import (
	"strconv"

	repository "github.com/khunfloat/sgcu-borrow-backend/repository/model"
	"gorm.io/gorm"
)

type itemRepositoryDB struct {
	db *gorm.DB	
}

func NewItemRepositoryDB(db *gorm.DB) itemRepositoryDB {
	db.AutoMigrate(repository.Item{})
	return itemRepositoryDB{db: db}
}

func (r itemRepositoryDB) GetAll() ([]repository.Item, error) {

	items := []repository.Item{}
	
	// query
	tx := r.db.Find(&items)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return items, nil
}

func (r itemRepositoryDB) GetById(id string) (*repository.Item, error) {

	itemId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	item := repository.Item{}
	tx := r.db.First(&item, itemId)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &item, nil
}

func (r itemRepositoryDB) Create(name string, currentAmount int, imgUrl string) (*repository.Item, error) {

	item := repository.Item{
		Name: name,
		CurrentAmount: currentAmount,
		ImgUrl: imgUrl,
	}

	tx := r.db.Create(&item)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &item, nil
}

func (r itemRepositoryDB) Update(id string, name string, currentAmount int, imgUrl string, borrowCount int) (*repository.Item, error) {

	itemId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	// Get data
	item := repository.Item{}
	tx := r.db.First(&item, itemId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	
	// Update data
	item.Name = name
	item.CurrentAmount = currentAmount
	item.ImgUrl = imgUrl
	item.BorrowCount = borrowCount	

	tx = r.db.Save(&item)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &item, nil
}