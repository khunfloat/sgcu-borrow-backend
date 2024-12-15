package repository

import (
	"fmt"

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

func (r lostRepositoryDB) GetById(id int) (*modelRepo.Lost, error) {

	lost := modelRepo.Lost{}
	tx := r.db.First(&lost, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &lost, nil
}

func (r lostRepositoryDB) Create(orderId int, itemId int, amount int) (*modelRepo.Lost, error) {

	lost := modelRepo.Lost{
		OrderId: orderId,
		ItemId: itemId,
		Amount: amount,
	}

	tx := r.db.Create(&lost)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &lost, nil
}

func (r lostRepositoryDB) Update(id int, orderId int, itemId int, amount int) (*modelRepo.Lost, error) {

	// Get data
	lost := modelRepo.Lost{}
	tx := r.db.First(&lost, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Update data
	lost.OrderId = orderId
	lost.ItemId = itemId
	lost.Amount = amount

	tx = r.db.Save(&lost)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &lost, nil
}

func (r lostRepositoryDB) DeleteById(id int) (error) {

	lost := modelRepo.Borrow{}
	tx := r.db.Delete(&lost, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
        return fmt.Errorf("no record found with id %d", id)
    }
	
	return nil
}
