package repository

import (
	"fmt"
	"strconv"
	"time"

	modelRepo "github.com/khunfloat/sgcu-borrow-backend/model/repository"
	"gorm.io/gorm"
)

type orderRepositoryDB struct {
	db *gorm.DB
}

func NewOrderRepositoryDB(db *gorm.DB) orderRepositoryDB {
	db.AutoMigrate(modelRepo.Order{})
	return orderRepositoryDB{db: db}
}

func (r orderRepositoryDB) GetAll() ([]modelRepo.Order, error) {

	orders := []modelRepo.Order{}

	// query
	tx := r.db.Find(&orders)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return orders, nil
}

func (r orderRepositoryDB) GetById(id string) (*modelRepo.Order, error) {

	orderId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	order := modelRepo.Order{}
	tx := r.db.First(&order, orderId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &order, nil
}

func (r orderRepositoryDB) Create(userId string, userOrg string, borrowDatetime time.Time, returnDatetime time.Time) (*modelRepo.Order, error) {

	order := modelRepo.Order{
		UserId:         userId,
		UserOrg:        userOrg,
		BorrowDatetime: borrowDatetime,
		ReturnDatetime: returnDatetime,
	}

	tx := r.db.Create(&order)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &order, nil
}

func (r orderRepositoryDB) Update(id string, userId string, userOrg string, borrowDatetime time.Time, returnDatetime time.Time) (*modelRepo.Order, error) {

	orderId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	// Get data
	order := modelRepo.Order{}
	tx := r.db.First(&order, orderId)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Update data
	order.UserId = userId
	order.UserOrg = userOrg
	order.BorrowDatetime = borrowDatetime
	order.ReturnDatetime = returnDatetime

	tx = r.db.Save(&order)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &order, nil
}

func (r orderRepositoryDB) DeleteById(id string) (error) {

	orderId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	order := modelRepo.Order{}
	tx := r.db.Delete(&order, orderId)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
        return fmt.Errorf("no record found with id %d", orderId)
    }
	
	return nil
}
