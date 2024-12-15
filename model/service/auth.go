package model

type UserSignUpRequest struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Tel      string `json:"tel"`
	Password string `json:"password"`
}

type StaffSignUpRequest struct {
	StaffId  string `json:"staff_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserSignInRequest struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

type StaffSignInRequest struct {
	StaffId  string `json:"staff_id"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
	Exp   int64  `json:"exp"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

type UserAuthService interface {
	SignUp(UserSignUpRequest) (*TokenResponse, error)
	SignIn(UserSignInRequest) (*TokenResponse, error)
}

type StaffAuthService interface {
	SignUp(StaffSignUpRequest) (*TokenResponse, error)
	SignIn(StaffSignInRequest) (*TokenResponse, error)
}
