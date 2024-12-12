package repository

import (
	"fmt"
	"strconv"
	"time"

	repository "github.com/khunfloat/sgcu-borrow-backend/repository/model"
	"gorm.io/gorm"
)

type returnRepositoryDB struct {
	db *gorm.DB
}

func NewReturnRepositoryDB(db *gorm.DB) returnRepositoryDB {
	db.AutoMigrate(repository.Return{})
	return returnRepositoryDB{db: db}
}

func (r returnRepositoryDB) GetAll() ([]repository.Return, error) {

	returnItems := []repository.Return{}

	// query
	tx := r.db.Find(&returnItems)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return returnItems, nil
}

func (r returnRepositoryDB) GetById(id string) (*repository.Return, error) {

	returnItemId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	returnItem := repository.Return{}
	tx := r.db.First(&returnItem, returnItemId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &returnItem, nil
}

func (r returnRepositoryDB) Create(orderId string, itemId string, amount int) (*repository.Return, error) {

	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		return nil, err
	}

	itemIdInt, err := strconv.Atoi(itemId)
	if err != nil {
		return nil, err
	}

	returnItem := repository.Return{
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

func (r returnRepositoryDB) Update(id string, orderId string, itemId string, amount int, dropoffDatetime time.Time) (*repository.Return, error) {

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
	returnItem := repository.Return{}
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

	returnItem := repository.Borrow{}
	tx := r.db.Delete(&returnItem, returnId)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
        return fmt.Errorf("no record found with id %d", returnId)
    }
	
	return nil
}
