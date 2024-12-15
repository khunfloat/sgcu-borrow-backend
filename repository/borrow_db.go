package repository

import (
	"fmt"

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

func (r borrowRepositoryDB) GetByOrderId(orderId int) ([]modelRepo.Borrow, error) {

	borrows := []modelRepo.Borrow{}

	// query
	tx := r.db.Where("order_id = ?", orderId).Find(&borrows)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return borrows, nil
}

func (r borrowRepositoryDB) GetByOrderIdAndItemId(orderId int, itemId int) (*modelRepo.Borrow, error) {

	borrow := modelRepo.Borrow{}

	// query
	tx := r.db.Where("order_id = ?", orderId).Where("item_id", itemId).First(&borrow)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &borrow, nil
}

func (r borrowRepositoryDB) GetById(id int) (*modelRepo.Borrow, error) {

	borrow := modelRepo.Borrow{}
	tx := r.db.First(&borrow, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &borrow, nil
}

func (r borrowRepositoryDB) Create(orderId int, itemId int, amount int) (*modelRepo.Borrow, error) {

	borrow := modelRepo.Borrow{
		OrderId: orderId,
		ItemId: itemId,
		Amount: amount,
	}

	tx := r.db.Create(&borrow)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &borrow, nil
}

func (r borrowRepositoryDB) Update(id int, orderId int, itemId int, amount int) (*modelRepo.Borrow, error) {

	// Get data
	borrow := modelRepo.Borrow{}
	tx := r.db.First(&borrow, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Update data
	borrow.OrderId = orderId
	borrow.ItemId = itemId
	borrow.Amount = amount

	tx = r.db.Save(&borrow)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &borrow, nil
}

func (r borrowRepositoryDB) DeleteById(id int) (error) {

	borrow := modelRepo.Borrow{}
	tx := r.db.Delete(&borrow, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
        return fmt.Errorf("no record found with id %d", id)
    }
	
	return nil
}
