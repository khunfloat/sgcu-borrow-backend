package service

import (
	"time"

	"github.com/khunfloat/sgcu-borrow-backend/errs"
	"github.com/khunfloat/sgcu-borrow-backend/logs"
	modelRepo "github.com/khunfloat/sgcu-borrow-backend/model/repository"
	modelServ "github.com/khunfloat/sgcu-borrow-backend/model/service"
	"github.com/khunfloat/sgcu-borrow-backend/utils"
)

type orderService struct {
	orderRepository  modelRepo.OrderRepository
	itemRepository   modelRepo.ItemRepository
	borrowRepository modelRepo.BorrowRepository
	returnRepository modelRepo.ReturnRepository
	lostRespository  modelRepo.LostRepository
}

func NewOrderService(
	orderRepository modelRepo.OrderRepository,
	itemRepository modelRepo.ItemRepository,
	borrowRepository modelRepo.BorrowRepository,
	returnRepository modelRepo.ReturnRepository,
	lostRespository modelRepo.LostRepository,
) orderService {
	return orderService{
		orderRepository: orderRepository,
		itemRepository:  itemRepository,
		borrowRepository: borrowRepository,
		returnRepository: returnRepository,
		lostRespository: lostRespository,
	}
}

func (s orderService) CreateOrder(orderRequest modelServ.NewOrderRequest) (*modelServ.OrderResponse, error) {
	userId := orderRequest.UserId
	userOrg := orderRequest.UserOrg
	borrowDt := orderRequest.BorrowDatetime
	returnDt := orderRequest.ReturnDatetime
	items := orderRequest.Items

	if userId == "" || userOrg == "" || borrowDt == "" || returnDt == "" || items == nil {
		return nil, errs.NewBadRequestError("invalid request body")
	}

	borrowDatetime, err := utils.String2Time(borrowDt)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError("invalid borrow datetime")
	}

	returnDatetime, err := utils.String2Time(returnDt)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError("invalid return datetime")
	}

	// 1. create order
	order, err := s.orderRepository.Create(
		userId,
		userOrg,
		borrowDatetime,
		returnDatetime,
	)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	// 2. create borrow
	itemResponses := []modelServ.ItemInOrderResponse{}

	for _, item := range items {
		itemId := item.ID
		amount := item.Amount

		// get item data
		existedItem, err := s.itemRepository.GetById(itemId)
		if err != nil {

			s.DeleteOrder(order.ID)

			logs.Error(err)
			return nil, errs.NewBadRequestError("item not found")
		}

		// check item amount
		if existedItem.CurrentAmount < amount {

			s.DeleteOrder(order.ID)

			logs.Error(err)
			return nil, errs.NewBadRequestError("item amount is not enough")
		}

		// create borrow item
		_, err = s.borrowRepository.Create(order.ID, itemId, amount)
		if err != nil {

			s.DeleteOrder(order.ID)

			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}

		// update item amount
		newCurrentAmount := existedItem.CurrentAmount - amount
		newBorrowCount := existedItem.BorrowCount + 1

		_, err = s.itemRepository.Update(itemId, existedItem.Name, newCurrentAmount, existedItem.ImgUrl, newBorrowCount)
		if err != nil {

			s.DeleteOrder(order.ID)

			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}

		itemResponse := modelServ.ItemInOrderResponse{
			ID: item.ID,
			Name: existedItem.Name,
			Amount: item.Amount,
			ImgUrl: existedItem.ImgUrl,
		}

		itemResponses = append(itemResponses, itemResponse)
	}


	orderResponse := modelServ.OrderResponse{
		ID: order.ID,
		UserId: userId,
		UserOrg: userOrg,
		BorrowDatetime: utils.Time2String(order.BorrowDatetime),
		ReturnDatetime: utils.Time2String(order.ReturnDatetime),
		Items: itemResponses,
	}

	return &orderResponse, nil
}

func (s orderService) GetOrders() ([]modelServ.OrderResponse, error) {
	
	orders, err := s.orderRepository.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	orderResponses := []modelServ.OrderResponse{}


	for _, order := range orders {
	
		orderId := order.ID
		userId := order.UserId
		userOrg := order.UserOrg
		borrowDatetime := order.BorrowDatetime
		returnDatetime := order.ReturnDatetime
		pickupDatetime := order.PickupDatetime
		dropoffDatetime := order.DropoffDatetime

		items, err := s.borrowRepository.GetByOrderId(orderId)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}
		
		itemResponses := []modelServ.ItemInOrderResponse{}

		for _, item := range items {
			borrowedItem, err := s.itemRepository.GetById(item.ItemId)
			if err != nil {
				logs.Error(err)
				return nil, errs.NewUnexpectedError()
			}

			itemResponse := modelServ.ItemInOrderResponse{
				ID: borrowedItem.ID,
				Name: borrowedItem.Name,
				Amount: item.Amount,
				ImgUrl: borrowedItem.ImgUrl,
			}

			itemResponses = append(itemResponses, itemResponse)
		}

		orderResponse := modelServ.OrderResponse{
			ID: orderId,
			UserId: userId,
			UserOrg: userOrg,
			BorrowDatetime: utils.Time2String(borrowDatetime),
			ReturnDatetime: utils.Time2String(returnDatetime),
			PickupDatetime: utils.Time2String(pickupDatetime),
			DropoffDatetime: utils.Time2String(dropoffDatetime),
			Items: itemResponses,
		}

		orderResponses = append(orderResponses, orderResponse)
	}
	
	return orderResponses, nil
}

func (s orderService) GetOrder(id int) (*modelServ.OrderResponse, error) {

	order, err := s.orderRepository.GetById(id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError("order not found")
	}

	itemResponses := []modelServ.ItemInOrderResponse{}

	items, err := s.borrowRepository.GetByOrderId(id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError("item not found")
	}

	for _, item := range items {
		existedItem, err := s.itemRepository.GetById(item.ItemId)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewBadRequestError("item not found")
		}

		itemResponse := modelServ.ItemInOrderResponse{
			ID: item.ItemId,
			Name: existedItem.Name,
			Amount: item.Amount,
			ImgUrl: existedItem.ImgUrl,
		}

		itemResponses = append(itemResponses, itemResponse)
	}


	orderResponse := modelServ.OrderResponse{
		ID: order.ID,
		UserId: order.UserId,
		UserOrg: order.UserOrg,
		BorrowDatetime: utils.Time2String(order.BorrowDatetime),
		ReturnDatetime: utils.Time2String(order.ReturnDatetime),
		Items: itemResponses,
	}

	return &orderResponse, nil
}
	
func (s orderService) UpdateOrder(orderRequest modelServ.UpdateOrderRequest) (*modelServ.OrderResponse, error) {
	id := orderRequest.ID
	userId := orderRequest.UserId
	userOrg := orderRequest.UserOrg
	borrowDt := orderRequest.BorrowDatetime
	returnDt := orderRequest.ReturnDatetime
	items := orderRequest.Items

	if id == 0 || userId == "" || userOrg == "" || borrowDt == "" || returnDt == "" || items == nil {
		return nil, errs.NewBadRequestError("invalid request body")
	}

	order, err := s.orderRepository.GetById(id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError("order not found")
	}

	if (order.PickupDatetime != time.Time{}) {
		logs.Error(err)
		return nil, errs.NewBadRequestError("order has been picked up")
	}

	borrowDatetime, err := utils.String2Time(borrowDt)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError("invalid borrow datetime")
	}

	returnDatetime, err := utils.String2Time(returnDt)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewBadRequestError("invalid return datetime")
	}

	//1. update order
	order, err = s.orderRepository.UpdateInfo(
		id,
		userId,
		userOrg,
		borrowDatetime,
		returnDatetime,
	)

	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	
	// 2. create borrow
	itemResponses := []modelServ.ItemInOrderResponse{}

	for _, item := range items {
		itemId := item.ID
		amount := item.Amount

		// get current borrow
		currentBorrow, err := s.borrowRepository.GetByOrderIdAndItemId(id, itemId)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewBadRequestError("Item doesn't in this order")
		}

		// get item data
		existedItem, err := s.itemRepository.GetById(itemId)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewBadRequestError("item not found")
		}

		// check item amount
		if existedItem.CurrentAmount + currentBorrow.Amount < amount {
			logs.Error(err)
			return nil, errs.NewBadRequestError("item amount is not enough")
		}

		// update borrow item
		_, err = s.borrowRepository.Update(currentBorrow.ID, order.ID, itemId, amount)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}

		// update item amount
		newCurrentAmount := existedItem.CurrentAmount + currentBorrow.Amount - amount

		_, err = s.itemRepository.Update(itemId, existedItem.Name, newCurrentAmount, existedItem.ImgUrl, existedItem.BorrowCount)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}

		itemResponse := modelServ.ItemInOrderResponse{
			ID: item.ID,
			Name: existedItem.Name,
			Amount: item.Amount,
			ImgUrl: existedItem.ImgUrl,
		}

		itemResponses = append(itemResponses, itemResponse)
	}

	orderResponse := modelServ.OrderResponse{
		ID: order.ID,
		UserId: userId,
		UserOrg: userOrg,
		BorrowDatetime: utils.Time2String(order.BorrowDatetime),
		ReturnDatetime: utils.Time2String(order.ReturnDatetime),
		Items: itemResponses,
	}

	return &orderResponse, nil
}

func (s orderService) DeleteOrder(id int) error {

	order, err := s.orderRepository.GetById(id)
	if err != nil {
		logs.Error(err)
		return errs.NewBadRequestError("order not found")
	}

	if (order.PickupDatetime != time.Time{}) {
		logs.Error(err)
		return errs.NewBadRequestError("order has been picked up")
	}

	borrows, err := s.borrowRepository.GetByOrderId(id)
	if err != nil {
		logs.Error(err)
		return errs.NewBadRequestError("borrow not found")
	}
	for _, borrow := range borrows {
		item, err := s.itemRepository.GetById(borrow.ItemId)
		if err != nil {
			logs.Error(err)
			return errs.NewBadRequestError("item not found")
		}

		newAmount := item.CurrentAmount + borrow.Amount
		newBorrowCount := item.BorrowCount - 1

		_, err = s.itemRepository.Update(item.ID, item.Name, newAmount, item.ImgUrl, newBorrowCount)
		if err != nil {
			logs.Error(err)
			return errs.NewBadRequestError("update item failed")
		}

		err = s.borrowRepository.DeleteById(borrow.ID)
		if err != nil {
			logs.Error(err)
			return errs.NewBadRequestError("delete borrow failed")
		}
	}

	s.orderRepository.DeleteById(id)

	return nil
}

func (s orderService) PickupOrder(modelServ.CheckOrderRequest) (*modelServ.OrderResponse, error) {
	return nil, nil
}
func (s orderService) DropoffOrder(modelServ.CheckOrderRequest) (*modelServ.OrderResponse, error) {
	return nil, nil
}
