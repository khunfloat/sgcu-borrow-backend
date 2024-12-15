package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	modelServ "github.com/khunfloat/sgcu-borrow-backend/model/service"
)

type orderHandler struct {
	orderService modelServ.OrderService
}

func NewOrderHandler(orderService modelServ.OrderService) orderHandler {
	return orderHandler{orderService: orderService}
}

func (h orderHandler) CreateOrder(c *fiber.Ctx) error {
	
	var request modelServ.NewOrderRequest

    if err := c.BodyParser(&request); err != nil {
       return handlerError(c, err)
    }

	order, err := h.orderService.CreateOrder(request)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(order)
}

func (h orderHandler) UpdateOrder(c *fiber.Ctx) error {
	
	var request modelServ.UpdateOrderRequest

    if err := c.BodyParser(&request); err != nil {
       return handlerError(c, err)
    }

	order, err := h.orderService.UpdateOrder(request)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(order)
}

func (h orderHandler) GetOrder(c *fiber.Ctx) error {
	
	id := c.Params("order_id")

	orderId, err := strconv.Atoi(id)
	if err != nil {
		return handlerError(c, err)
	}

	order, err := h.orderService.GetOrder(orderId)
	if err != nil {
		return handlerError(c, err)
	}
	return c.JSON(order)
}

func (h orderHandler) GetOrders(c *fiber.Ctx) error {
	
	orders, err := h.orderService.GetOrders()
	if err != nil {
		return handlerError(c, err)
	}
	return c.JSON(fiber.Map{
		"orders": orders,
	})
}

func (h orderHandler) DeleteOrder(c *fiber.Ctx) error {
	
	id := c.Params("order_id")

	orderId, err := strconv.Atoi(id)
	if err != nil {
		return handlerError(c, err)
	}

	err = h.orderService.DeleteOrder(orderId)
	if err != nil {
		return handlerError(c, err)
	}
	return nil
}

func (h orderHandler) PickupOrder(c *fiber.Ctx) error {
	
	var request modelServ.CheckOrderRequest

    if err := c.BodyParser(&request); err != nil {
       return handlerError(c, err)
    }

	order, err := h.orderService.PickupOrder(request)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(order)
}

func (h orderHandler) DropoffOrder(c *fiber.Ctx) error {
	
	var request modelServ.CheckOrderRequest

    if err := c.BodyParser(&request); err != nil {
       return handlerError(c, err)
    }

	order, err := h.orderService.DropoffOrder(request)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(order)
}