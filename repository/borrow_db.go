package repository

import (
	"fmt"
	"strconv"
	"time"

	repository "github.com/khunfloat/sgcu-borrow-backend/repository/model"
	"gorm.io/gorm"
)

type borrowRepositoryDB struct {
	db *gorm.DB
}

func NewBorrowRepositoryDB(db *gorm.DB) borrowRepositoryDB {
	db.AutoMigrate(repository.Borrow{})
	return borrowRepositoryDB{db: db}
}

func (r borrowRepositoryDB) GetAll() ([]repository.Borrow, error) {

	borrows := []repository.Borrow{}

	// query
	tx := r.db.Find(&borrows)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return borrows, nil
}

func (r borrowRepositoryDB) GetById(id string) (*repository.Borrow, error) {

	borrowId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	borrow := repository.Borrow{}
	tx := r.db.First(&borrow, borrowId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &borrow, nil
}

func (r borrowRepositoryDB) Create(orderId string, itemId string, amount int) (*repository.Borrow, error) {

	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		return nil, err
	}

	itemIdInt, err := strconv.Atoi(itemId)
	if err != nil {
		return nil, err
	}

	borrow := repository.Borrow{
		OrderId: orderIdInt,
		ItemId: itemIdInt,
		Amount: amount,
	}

	tx := r.db.Create(&borrow)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &borrow, nil
}

func (r borrowRepositoryDB) Update(id string, orderId string, itemId string, amount int, pickupDatetime time.Time) (*repository.Borrow, error) {

	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		return nil, err
	}

	itemIdInt, err := strconv.Atoi(itemId)
	if err != nil {
		return nil, err
	}

	borrowId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	// Get data
	borrow := repository.Borrow{}
	tx := r.db.First(&borrow, borrowId)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Update data
	borrow.OrderId = orderIdInt
	borrow.ItemId = itemIdInt
	borrow.Amount = amount
	borrow.PickupDatetime = pickupDatetime

	tx = r.db.Save(&borrow)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &borrow, nil
}

func (r borrowRepositoryDB) DeleteById(id string) (error) {

	borrowId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	borrow := repository.Borrow{}
	tx := r.db.Delete(&borrow, borrowId)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
        return fmt.Errorf("no record found with id %d", borrowId)
    }
	
	return nil
}
