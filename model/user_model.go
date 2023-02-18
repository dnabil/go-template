package model

import (
	customvalidation "go-template/sdk/custom_validation"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// for creating an account (register)
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FName    string `json:"f_name"`
	LName    string `json:"l_name"`
}
func (m *CreateUserRequest) Validate() error {
	return validation.ValidateStruct(m,
		validation.Field(&m.Email, validation.Required, is.Email),
		validation.Field(&m.Password, customvalidation.PasswordValidation...),
		validation.Field(&m.FName, validation.Required, validation.Length(0, 50), validation.By(customvalidation.CheckName)),
		validation.Field(&m.LName, validation.Required, validation.Length(0, 50), validation.By(customvalidation.CheckName)),
	)
}

type CreateUserResponse struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Email    string `json:"email"`
	FName    string `json:"f_name"`
	LName    string `json:"l_name"`
}

type LoginUserRequest struct{
	Email	string `json:"email"`
	Password string `json:"password"`
}
func (m *LoginUserRequest) Validate() error{
	return validation.ValidateStruct(m,
		validation.Field(&m.Email, validation.Required, is.Email),
		validation.Field(&m.Password, customvalidation.PasswordRequiredLength...),
	)
}

type LoginUserResponse struct{
	Authorization string `json:"Authorization"`
}

// for jwt authentication
type UserClaims struct{
	jwt.RegisteredClaims
}
func NewUserClaims(uuid string, t time.Duration) (UserClaims){
	return UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: uuid,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t)), // bisa dimasukin config kali?
		},
	}
}