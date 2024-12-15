package repository

import (
	"fmt"

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

func (r itemRepositoryDB) GetFrequentlyBorrowed() ([]modelRepo.Item, error) {
	
	items := []modelRepo.Item{}

	tx := r.db.Order("borrow_count DESC").Limit(10).Find(&items)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return items, nil
}

func (r itemRepositoryDB) GetById(id int) (*modelRepo.Item, error) {

	item := modelRepo.Item{}
	tx := r.db.First(&item, id)
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

func (r itemRepositoryDB) Update(id int, name string, currentAmount int, imgUrl string, borrowCount int) (*modelRepo.Item, error) {

	// Get data
	item := modelRepo.Item{}
	tx := r.db.First(&item, id)
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

func (r itemRepositoryDB) AddCurrentAmount(id int, amount int) (*modelRepo.Item, error) {

	// Get data
	item := modelRepo.Item{}
	tx := r.db.First(&item, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	
	// Update data
	item.CurrentAmount += amount

	tx = r.db.Save(&item)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &item, nil
}

func (r itemRepositoryDB) UpdateCurrentAmount(id int, amount int) (*modelRepo.Item, error) {

	// Get data
	item := modelRepo.Item{}
	tx := r.db.First(&item, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	
	// Update data
	item.CurrentAmount = amount

	tx = r.db.Save(&item)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &item, nil
}

func (r itemRepositoryDB) SubtractCurrentAmount(id int, amount int) (*modelRepo.Item, error) {

	// Get data
	item := modelRepo.Item{}
	tx := r.db.First(&item, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	
	// Update data
	item.CurrentAmount -= amount

	tx = r.db.Save(&item)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &item, nil
}

func (r itemRepositoryDB) DeleteById(id int) (error) {

	item := modelRepo.Item{}
	tx := r.db.Delete(&item, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
        return fmt.Errorf("no record found with id %d", id)
    }
	
	return nil
}