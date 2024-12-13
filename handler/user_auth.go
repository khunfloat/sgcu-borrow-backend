package handler

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt"
	"github.com/khunfloat/sgcu-borrow-backend/errs"
	modelServ "github.com/khunfloat/sgcu-borrow-backend/model/service"
	"github.com/spf13/viper"
)

type userAuthHandler struct {
	userAuthService modelServ.UserAuthService
}

func NewUserAuthHandler(userAuthService modelServ.UserAuthService) userAuthHandler {
	return userAuthHandler{userAuthService: userAuthService}
}

func (h userAuthHandler) SignUp(c *fiber.Ctx) error {
	
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

func (h userAuthHandler) SignIn(c *fiber.Ctx) error {
	
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

func (h userAuthHandler) AuthErrorHandler(c *fiber.Ctx, err error) error {
	return handlerError(c, errs.NewUnAuthorizedError())
}

func (h userAuthHandler) AuthSuccessHandler(c *fiber.Ctx) error {
	c.Next()
	return nil
}

func (h userAuthHandler) AuthorizationRequired() fiber.Handler {
    return jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:   []byte(viper.GetString("app.jwt-secret")),
		ErrorHandler: h.AuthErrorHandler,
		SuccessHandler: h.AuthSuccessHandler,
	})
}

func (h userAuthHandler) IsUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
  
	if role != "user" {
	  return handlerError(c, errs.NewUnAuthorizedError())
	}
  
	return c.Next()
}