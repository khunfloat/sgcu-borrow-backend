package repository

import (
	"strconv"

	modelRepo "github.com/khunfloat/sgcu-borrow-backend/model/repository"
	"gorm.io/gorm"
)

type itemRepositoryDB struct {
	db *gorm.DB	
}

func NewItemRepositoryDB(db *gorm.DB) itemRepositoryDB {
	db.AutoMigrate(modelRepo.Item{})
	return itemRepositoryDB{db: db}
}

func (r itemRepositoryDB) GetAll() ([]modelRepo.Item, error) {

	items := []modelRepo.Item{}
	
	// query
	tx := r.db.Find(&items)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return items, nil
}

func (r itemRepositoryDB) GetById(id string) (*modelRepo.Item, error) {

	itemId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	item := modelRepo.Item{}
	tx := r.db.First(&item, itemId)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &item, nil
}

func (r itemRepositoryDB) Create(name string, currentAmount int, imgUrl string) (*modelRepo.Item, error) {

	item := modelRepo.Item{
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

func (r itemRepositoryDB) Update(id string, name string, currentAmount int, imgUrl string, borrowCount int) (*modelRepo.Item, error) {

	itemId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	// Get data
	item := modelRepo.Item{}
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