package repository

import (
	"fmt"

	modelRepo "github.com/khunfloat/sgcu-borrow-backend/model/repository"
	"gorm.io/gorm"
)

type returnRepositoryDB struct {
	db *gorm.DB
}

func NewReturnRepositoryDB(db *gorm.DB) returnRepositoryDB {
	db.AutoMigrate(modelRepo.Return{})
	return returnRepositoryDB{db: db}
}

func (r returnRepositoryDB) GetAll() ([]modelRepo.Return, error) {

	returnItems := []modelRepo.Return{}

	// query
	tx := r.db.Find(&returnItems)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return returnItems, nil
}

func (r returnRepositoryDB) GetById(id int) (*modelRepo.Return, error) {

	returnItem := modelRepo.Return{}
	tx := r.db.First(&returnItem, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &returnItem, nil
}

func (r returnRepositoryDB) Create(orderId int, itemId int, amount int) (*modelRepo.Return, error) {

	returnItem := modelRepo.Return{
		OrderId: orderId,
		ItemId: itemId,
		Amount: amount,
	}

	tx := r.db.Create(&returnItem)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &returnItem, nil
}

func (r returnRepositoryDB) Update(id int, orderId int, itemId int, amount int) (*modelRepo.Return, error) {

	// Get data
	returnItem := modelRepo.Return{}
	tx := r.db.First(&returnItem, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Update data
	returnItem.OrderId = orderId
	returnItem.ItemId = itemId
	returnItem.Amount = amount

	tx = r.db.Save(&returnItem)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &returnItem, nil
}

func (r returnRepositoryDB) DeleteById(id int) (error) {

	returnItem := modelRepo.Return{}
	tx := r.db.Delete(&returnItem, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
        return fmt.Errorf("no record found with id %d", id)
    }
	
	return nil
}
