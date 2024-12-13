package repository

import (
	"fmt"
	"strconv"
	"time"

	modelRepo "github.com/khunfloat/sgcu-borrow-backend/model/repository"
	"gorm.io/gorm"
)

type borrowRepositoryDB struct {
	db *gorm.DB
}

func NewBorrowRepositoryDB(db *gorm.DB) borrowRepositoryDB {
	db.AutoMigrate(modelRepo.Borrow{})
	return borrowRepositoryDB{db: db}
}

func (r borrowRepositoryDB) GetAll() ([]modelRepo.Borrow, error) {

	borrows := []modelRepo.Borrow{}

	// query
	tx := r.db.Find(&borrows)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return borrows, nil
}

func (r borrowRepositoryDB) GetById(id string) (*modelRepo.Borrow, error) {

	borrowId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	borrow := modelRepo.Borrow{}
	tx := r.db.First(&borrow, borrowId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &borrow, nil
}

func (r borrowRepositoryDB) Create(orderId string, itemId string, amount int) (*modelRepo.Borrow, error) {

	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		return nil, err
	}

	itemIdInt, err := strconv.Atoi(itemId)
	if err != nil {
		return nil, err
	}

	borrow := modelRepo.Borrow{
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

func (r borrowRepositoryDB) Update(id string, orderId string, itemId string, amount int, pickupDatetime time.Time) (*modelRepo.Borrow, error) {

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
	borrow := modelRepo.Borrow{}
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

	borrow := modelRepo.Borrow{}
	tx := r.db.Delete(&borrow, borrowId)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
        return fmt.Errorf("no record found with id %d", borrowId)
    }
	
	return nil
}
