package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/khunfloat/sgcu-borrow-backend/errs"
)

func handlerError(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case errs.AppError:
		return c.Status(e.Code).SendString(e.Message)
	}
	return c.Status(http.StatusInternalServerError).SendString(err.Error())
}