package handler

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/khunfloat/sgcu-borrow-backend/errs"
	modelServ "github.com/khunfloat/sgcu-borrow-backend/model/service"
	"github.com/spf13/viper"
)

type authHandler struct {
	userAuthService modelServ.UserAuthService
	staffAuthService modelServ.StaffAuthService
}

func NewAuthHandler(userAuthService modelServ.UserAuthService, staffAuthService modelServ.StaffAuthService) authHandler {
	return authHandler{
		userAuthService: userAuthService,
		staffAuthService: staffAuthService,
	}
}

func (h authHandler) UserSignUp(c *fiber.Ctx) error {
	
	var request modelServ.UserSignUpRequest
    if err := c.BodyParser(&request); err != nil {
       return handlerError(c, err)
    }

	token, err := h.userAuthService.SignUp(request)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(token)
}

func (h authHandler) UserSignIn(c *fiber.Ctx) error {
	
	var request modelServ.UserSignInRequest
    if err := c.BodyParser(&request); err != nil {
       return handlerError(c, err)
    }

	token, err := h.userAuthService.SignIn(request)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(token)
}

func (h authHandler) StaffSignUp(c *fiber.Ctx) error {
	
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

func (h authHandler) StaffSignIn(c *fiber.Ctx) error {
	
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

func (h authHandler) AuthErrorHandler(c *fiber.Ctx, err error) error {
	return handlerError(c, errs.NewUnAuthorizedError())
}

func (h authHandler) AuthSuccessHandler(c *fiber.Ctx) error {
	c.Next()
	return nil
}

func (h authHandler) AuthorizationRequired() fiber.Handler {
    return jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:   []byte(viper.GetString("app.jwt-secret")),
		ErrorHandler: h.AuthErrorHandler,
		SuccessHandler: h.AuthSuccessHandler,
	})
}

func (h authHandler) IsStaff(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
  
	if role != "staff" {
	  return handlerError(c, errs.NewUnAuthorizedError())
	}
  
	return c.Next()
}

func (h authHandler) IsAdmin(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
  
	if role != "admin" {
	  return handlerError(c, errs.NewUnAuthorizedError())
	}
  
	return c.Next()
}