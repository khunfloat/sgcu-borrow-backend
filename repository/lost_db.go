package repository

import (
	"fmt"
	"strconv"

	modelRepo "github.com/khunfloat/sgcu-borrow-backend/model/repository"
	"gorm.io/gorm"
)

type lostRepositoryDB struct {
	db *gorm.DB
}

func NewLostRepositoryDB(db *gorm.DB) lostRepositoryDB {
	db.AutoMigrate(modelRepo.Lost{})
	return lostRepositoryDB{db: db}
}

func (r lostRepositoryDB) GetAll() ([]modelRepo.Lost, error) {

	losts := []modelRepo.Lost{}

	// query
	tx := r.db.Find(&losts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return losts, nil
}

func (r lostRepositoryDB) GetById(id string) (*modelRepo.Lost, error) {

	lostId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	lost := modelRepo.Lost{}
	tx := r.db.First(&lost, lostId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &lost, nil
}

func (r lostRepositoryDB) Create(orderId string, itemId string, amount int) (*modelRepo.Lost, error) {

	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		return nil, err
	}

	itemIdInt, err := strconv.Atoi(itemId)
	if err != nil {
		return nil, err
	}

	lost := modelRepo.Lost{
		OrderId: orderIdInt,
		ItemId: itemIdInt,
		Amount: amount,
	}

	tx := r.db.Create(&lost)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &lost, nil
}

func (r lostRepositoryDB) Update(id string, orderId string, itemId string, amount int) (*modelRepo.Lost, error) {

	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		return nil, err
	}

	itemIdInt, err := strconv.Atoi(itemId)
	if err != nil {
		return nil, err
	}

	lostId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	// Get data
	lost := modelRepo.Lost{}
	tx := r.db.First(&lost, lostId)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Update data
	lost.OrderId = orderIdInt
	lost.ItemId = itemIdInt
	lost.Amount = amount

	tx = r.db.Save(&lost)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &lost, nil
}

func (r lostRepositoryDB) DeleteById(id string) (error) {

	lostId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	lost := modelRepo.Borrow{}
	tx := r.db.Delete(&lost, lostId)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
        return fmt.Errorf("no record found with id %d", lostId)
    }
	
	return nil
}
