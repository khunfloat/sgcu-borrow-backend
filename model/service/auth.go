package model

type SignUpRequest struct {
	StudentId string `json:"student_id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
}

type SignInRequest struct {
	StudentId string `json:"student_id"`
	Password  string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
	Exp   int64  `json:"exp"`
	User  string `json:"user"`
}

type AuthService interface {
	SignUp(SignUpRequest) (*TokenResponse, error)
	SignIn(SignInRequest) (*TokenResponse, error)
}