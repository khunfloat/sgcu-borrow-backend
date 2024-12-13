package handler

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt"
	"github.com/khunfloat/sgcu-borrow-backend/errs"
	modelServ "github.com/khunfloat/sgcu-borrow-backend/model/service"
	"github.com/spf13/viper"
)

type staffAuthHandler struct {
	staffAuthService modelServ.StaffAuthService
}

func NewStaffAuthHandler(staffAuthService modelServ.StaffAuthService) staffAuthHandler {
	return staffAuthHandler{staffAuthService: staffAuthService}
}

func (h staffAuthHandler) SignUp(c *fiber.Ctx) error {
	
	var request modelServ.StaffSignUpRequest
    if err := c.BodyParser(&request); err != nil {
       return handlerError(c, err)
    }

	token, err := h.staffAuthService.SignUp(request)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(token)
}

func (h staffAuthHandler) SignIn(c *fiber.Ctx) error {
	
	var request modelServ.StaffSignInRequest
    if err := c.BodyParser(&request); err != nil {
       return handlerError(c, err)
    }

	token, err := h.staffAuthService.SignIn(request)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(token)
}

func (h staffAuthHandler) AuthErrorHandler(c *fiber.Ctx, err error) error {
	return handlerError(c, errs.NewUnAuthorizedError())
}

func (h staffAuthHandler) AuthSuccessHandler(c *fiber.Ctx) error {
	c.Next()
	return nil
}

func (h staffAuthHandler) AuthorizationRequired() fiber.Handler {
    return jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:   []byte(viper.GetString("app.jwt-secret")),
		ErrorHandler: h.AuthErrorHandler,
		SuccessHandler: h.AuthSuccessHandler,
	})
}

func (h staffAuthHandler) IsStaff(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
  
	if role != "staff" {
	  return handlerError(c, errs.NewUnAuthorizedError())
	}
  
	return c.Next()
}

func (h staffAuthHandler) IsAdmin(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
  
	if role != "admin" {
	  return handlerError(c, errs.NewUnAuthorizedError())
	}
  
	return c.Next()
}