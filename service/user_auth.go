package service

import (
	"github.com/khunfloat/sgcu-borrow-backend/errs"
	"github.com/khunfloat/sgcu-borrow-backend/logs"
	"github.com/khunfloat/sgcu-borrow-backend/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	modelRepo "github.com/khunfloat/sgcu-borrow-backend/model/repository"
	modelServ "github.com/khunfloat/sgcu-borrow-backend/model/service"
)

type userAuthService struct {
	userRepository modelRepo.UserRepository
}

func NewUserAuthService(userRepository modelRepo.UserRepository) userAuthService {
	return userAuthService{userRepository: userRepository}
}

func (s userAuthService) SignUp(userSignUpRequest modelServ.UserSignUpRequest) (*modelServ.TokenResponse, error) {
	userId := userSignUpRequest.UserId
	name := userSignUpRequest.Name
	tel := userSignUpRequest.Tel
	password := userSignUpRequest.Password

	if userId == "" || name == "" || tel == "" || password == ""{
		return nil, errs.NewBadRequestError("invalid signup credentials")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userSignUpRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}
	password = string(hash)

	_, err = s.userRepository.Create(userId, name, tel ,password)
	if err != nil {

		if err == gorm.ErrDuplicatedKey {
			logs.Error(err)
			return nil, errs.NewBadRequestError("user already exists")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	token, exp, err := utils.CreateJWTToken(userId, "user")
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	tokenResponse := modelServ.TokenResponse{
		Token: token,
		Exp:   exp,
		User: userId,
	}

	return &tokenResponse, nil
}

func (s userAuthService) SignIn(userSignInRequest modelServ.UserSignInRequest) (*modelServ.TokenResponse, error) {
	userId := userSignInRequest.UserId
	password := userSignInRequest.Password

	if userId == "" || password == "" {
		return nil, errs.NewBadRequestError("invalid signip credentials")
	}

	user, err := s.userRepository.GetById(userId)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewNotFoundError("user doesn't existed")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errs.NewBadRequestError("invalid password")
	}

	token, exp, err := utils.CreateJWTToken(userId, "user")
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	tokenResponse := modelServ.TokenResponse{
		Token: token,
		Exp:   exp,
		User: userId,
	}

	return &tokenResponse, nil
}