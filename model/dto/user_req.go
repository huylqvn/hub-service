package dto

import "encoding/json"

// LoginDto defines a data transfer object for login.
type LoginDto struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// NewLoginDto is constructor.
func NewLoginDto() *LoginDto {
	return &LoginDto{}
}

// ToString is return string of object
func (l *LoginDto) ToString() (string, error) {
	bytes, err := json.Marshal(l)
	return string(bytes), err
}

type AuthenticationFail struct {
	Message string `json:"message"`
}

type GetUserInfo struct {
	Name string `json:"name" query:"name"`
}
