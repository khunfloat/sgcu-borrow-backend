package service

import (
	"github.com/khunfloat/sgcu-borrow-backend/errs"
	"github.com/khunfloat/sgcu-borrow-backend/logs"
	modelRepo "github.com/khunfloat/sgcu-borrow-backend/model/repository"
	modelServ "github.com/khunfloat/sgcu-borrow-backend/model/service"
	"gorm.io/gorm"
)

type itemService struct {
	itemRepository modelRepo.ItemRepository
}

func NewItemService(itemRepository modelRepo.ItemRepository) itemService {
	return itemService{itemRepository: itemRepository}
}

func (s itemService) GetItems() ([]modelServ.ItemResponse, error) {

	items, err := s.itemRepository.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	itemResponses := []modelServ.ItemResponse{}
	for _, item := range items {
		itemResponse := modelServ.ItemResponse{
			ID: item.ID,
			Name: 	item.Name,
			CurrentAmount: item.CurrentAmount,
			ImgUrl: item.ImgUrl,
			BorrowCount: item.BorrowCount,
		}
		itemResponses = append(itemResponses, itemResponse)
	}

	return itemResponses, nil
}

func (s itemService) GetFrequentlyBorrowed() ([]modelServ.ItemResponse, error) {

	items, err := s.itemRepository.GetFrequentlyBorrowed()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	itemResponses := []modelServ.ItemResponse{}
	for _, item := range items {
		itemResponse := modelServ.ItemResponse{
			ID: item.ID,
			Name: 	item.Name,
			CurrentAmount: item.CurrentAmount,
			ImgUrl: item.ImgUrl,
			BorrowCount: item.BorrowCount,
		}
		itemResponses = append(itemResponses, itemResponse)
	}

	return itemResponses, nil
}

func (s itemService) GetItem(id int) (*modelServ.ItemResponse, error) {
	item, err := s.itemRepository.GetById(id)
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			logs.Error(err)
			return nil, errs.NewNotFoundError("user not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	itemResponse := modelServ.ItemResponse{
		ID: item.ID,
		Name: 	item.Name,
		CurrentAmount: item.CurrentAmount,
		ImgUrl: item.ImgUrl,
		BorrowCount: item.BorrowCount,
	}

	return &itemResponse, nil
}

func (s itemService) CreateItem(itemRequest modelServ.NewItemRequest) (*modelServ.ItemResponse, error) {
	name := itemRequest.Name
	currentAmount := itemRequest.CurrentAmount
	imgUrl := itemRequest.ImgUrl

	if name == "" || currentAmount == 0 || imgUrl == "" {
		return nil, errs.NewBadRequestError("invalid request body")
	}

	item, err := s.itemRepository.Create(name, currentAmount, imgUrl)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	itemResponse := modelServ.ItemResponse{
		ID: item.ID,
		Name: 	item.Name,
		CurrentAmount: item.CurrentAmount,
		ImgUrl: item.ImgUrl,
		BorrowCount: item.BorrowCount,
	}

	return &itemResponse, nil
}

func (s itemService) UpdateItem(itemRequest modelServ.UpdateItemRequest) (*modelServ.ItemResponse, error) {
	id := itemRequest.ID
	name := itemRequest.Name
	currentAmount := itemRequest.CurrentAmount
	imgUrl := itemRequest.ImgUrl
	borrowCount := itemRequest.BorrowCount

	if id == nil || name == "" || currentAmount == nil || imgUrl == "" || borrowCount == nil {
		return nil, errs.NewBadRequestError("invalid request body")
	}

	idInt := *itemRequest.ID
	currentAmountInt := *itemRequest.CurrentAmount
	borrowCountInt := *itemRequest.BorrowCount


	item, err := s.itemRepository.Update(idInt, name, currentAmountInt, imgUrl, borrowCountInt)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	itemResponse := modelServ.ItemResponse{
		ID: item.ID,
		Name: 	item.Name,
		CurrentAmount: item.CurrentAmount,
		ImgUrl: item.ImgUrl,
		BorrowCount: item.BorrowCount,
	}

	return &itemResponse, nil
}

func (s itemService) DeleteItem(id int) (error) {
	err := s.itemRepository.DeleteById(id)
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			logs.Error(err)
			return errs.NewNotFoundError("user not found")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	return nil
}