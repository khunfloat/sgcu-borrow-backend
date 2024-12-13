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

type staffAuthService struct {
	staffRepository modelRepo.StaffRepository
}

func NewStaffAuthService(staffRepository modelRepo.StaffRepository) staffAuthService {
	return staffAuthService{staffRepository: staffRepository}
}

func (s staffAuthService) SignUp(staffSignUpRequest modelServ.StaffSignUpRequest) (*modelServ.TokenResponse, error) {
	staffId := staffSignUpRequest.StaffId
	name := staffSignUpRequest.Name
	password := staffSignUpRequest.Password

	if staffId == "" || name == "" || password == ""{
		return nil, errs.NewBadRequestError("invalid signup credentials")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(staffSignUpRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}
	password = string(hash)

	staff, err := s.staffRepository.Create(staffId, name, password)
	if err != nil {

		if err == gorm.ErrDuplicatedKey {
			logs.Error(err)
			return nil, errs.NewBadRequestError("staff already exists")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	role := staff.Role

	token, exp, err := utils.CreateJWTToken(staffId, role)
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	tokenResponse := modelServ.TokenResponse{
		Token: token,
		Exp:   exp,
		User: staffId,
	}

	return &tokenResponse, nil
}

func (s staffAuthService) SignIn(staffSignInRequest modelServ.StaffSignInRequest) (*modelServ.TokenResponse, error) {
	staffId := staffSignInRequest.StaffId
	password := staffSignInRequest.Password

	if staffId == "" || password == "" {
		return nil, errs.NewBadRequestError("invalid signip credentials")
	}

	staff, err := s.staffRepository.GetById(staffId)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewNotFoundError("staff doesn't existed")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(password)); err != nil {
		return nil, errs.NewBadRequestError("invalid password")
	}

	role := staff.Role

	token, exp, err := utils.CreateJWTToken(staffId, role)
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	tokenResponse := modelServ.TokenResponse{
		Token: token,
		Exp:   exp,
		User: staffId,
	}

	return &tokenResponse, nil
}