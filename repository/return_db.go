package repository

import (
	"fmt"
	"strconv"
	"time"

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

func (r returnRepositoryDB) GetById(id string) (*modelRepo.Return, error) {

	returnItemId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	returnItem := modelRepo.Return{}
	tx := r.db.First(&returnItem, returnItemId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &returnItem, nil
}

func (r returnRepositoryDB) Create(orderId string, itemId string, amount int) (*modelRepo.Return, error) {

	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		return nil, err
	}

	itemIdInt, err := strconv.Atoi(itemId)
	if err != nil {
		return nil, err
	}

	returnItem := modelRepo.Return{
		OrderId: orderIdInt,
		ItemId: itemIdInt,
		Amount: amount,
	}

	tx := r.db.Create(&returnItem)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &returnItem, nil
}

func (r returnRepositoryDB) Update(id string, orderId string, itemId string, amount int, dropoffDatetime time.Time) (*modelRepo.Return, error) {

	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		return nil, err
	}

	itemIdInt, err := strconv.Atoi(itemId)
	if err != nil {
		return nil, err
	}

	returnItemId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	// Get data
	returnItem := modelRepo.Return{}
	tx := r.db.First(&returnItem, returnItemId)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Update data
	returnItem.OrderId = orderIdInt
	returnItem.ItemId = itemIdInt
	returnItem.Amount = amount
	returnItem.DropoffDatetime = dropoffDatetime

	tx = r.db.Save(&returnItem)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &returnItem, nil
}

func (r returnRepositoryDB) DeleteById(id string) (error) {

	returnId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	returnItem := modelRepo.Borrow{}
	tx := r.db.Delete(&returnItem, returnId)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
        return fmt.Errorf("no record found with id %d", returnId)
    }
	
	return nil
}
