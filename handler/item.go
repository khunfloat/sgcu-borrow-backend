package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	modelServ "github.com/khunfloat/sgcu-borrow-backend/model/service"
)

type itemHandler struct {
	itemService modelServ.ItemService
}

func NewItemHandler(itemService modelServ.ItemService) itemHandler {
	return itemHandler{itemService: itemService}
}

func (h itemHandler) GetItems(c *fiber.Ctx) error {
	
	items, err := h.itemService.GetItems()
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(fiber.Map{
		"items": items,
	})
}

func (h itemHandler) GetFrequentlyBorrowed(c *fiber.Ctx) error {
	
	items, err := h.itemService.GetFrequentlyBorrowed()
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(fiber.Map{
		"items": items,
	})
}

func (h itemHandler) GetItem(c *fiber.Ctx) error {
	
	id := c.Params("item_id")

	itemId, err := strconv.Atoi(id)
	if err != nil {
		return handlerError(c, err)
	}

	item, err := h.itemService.GetItem(itemId)
	if err != nil {
		return handlerError(c, err)
	}
	return c.JSON(item)
}

func (h itemHandler) CreateItem(c *fiber.Ctx) error {
	
	var request modelServ.NewItemRequest
    if err := c.BodyParser(&request); err != nil {
       return handlerError(c, err)
    }

	item, err := h.itemService.CreateItem(request)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(item)
}

func (h itemHandler) UpdateItem(c *fiber.Ctx) error {
	
	var request modelServ.UpdateItemRequest
    if err := c.BodyParser(&request); err != nil {
       return handlerError(c, err)
    }

	item, err := h.itemService.UpdateItem(request)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(item)
}

func (h itemHandler) DeleteItem(c *fiber.Ctx) error {
	
	id := c.Params("item_id")

	itemId, err := strconv.Atoi(id)
	if err != nil {
		return handlerError(c, err)
	}

	err = h.itemService.DeleteItem(itemId)
	if err != nil {
		return handlerError(c, err)
	}
	return nil
}