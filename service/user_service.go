package service

import (
	"context"
	"fmt"
	"go-template/entity"
	"go-template/model"
	"go-template/repository"
	"os"
	"time"

	// errwrapper used in every service method  to pass errors to handler layer
	ew "go-template/sdk/apires/errwrapper"
	sdk_jwt "go-template/sdk/jwt"

	"go-template/sdk/password"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type UserService interface {
	Register(ctx context.Context, req model.CreateUserRequest) (model.CreateUserResponse, ew.ErrWrapper)
	Login(ctx context.Context, req model.LoginUserRequest) (model.LoginUserResponse, ew.ErrWrapper)
}


type userServiceInterfaceImpl struct {
	userRepository repository.UserRepository
}
func NewUserService(repo *repository.UserRepository) UserService {
	return &userServiceInterfaceImpl{
		userRepository: *repo,
	}
}

func (s *userServiceInterfaceImpl) Register(ctx context.Context, req model.CreateUserRequest) (model.CreateUserResponse, ew.ErrWrapper){
	//hashing password
	hashedPassword, err := password.Hash(req.Password)
	if err != nil { // password hashing failed
		return model.CreateUserResponse{}, ew.New(http.StatusInternalServerError, err)
	}

	// create uuid
	userUUID := uuid.New()
	_, err = s.userRepository.FindByUUID(ctx, userUUID.String())
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return model.CreateUserResponse{}, ew.New(500, fmt.Errorf("user uuid error:%s", err.Error()))
		}
	}

	user := entity.User{
		Email: req.Email,
		UUID: userUUID.String(),
		Password: hashedPassword,
		FName: req.FName,
		LName: req.LName,
	}
	//create user
	err = s.userRepository.Create(ctx, &user)
	if err != nil{
		return model.CreateUserResponse{}, ew.New(http.StatusInternalServerError, err)
	}

	res := model.CreateUserResponse{
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		Email: user.Email,
		FName: user.FName,
		LName: user.LName,
	}
	return res, nil
}

func (s *userServiceInterfaceImpl) Login(ctx context.Context, req model.LoginUserRequest) (model.LoginUserResponse, ew.ErrWrapper){

	user, err := s.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return model.LoginUserResponse{}, ew.New(http.StatusNotFound, fmt.Errorf("Account not found"))
	}

	// if password is not correct, then return error
	if !password.Compare(user.Password, req.Password){
		return model.LoginUserResponse{}, ew.New(http.StatusForbidden, fmt.Errorf("Wrong password"))
	}

	// IMPLEMENT JWT
	// get exp time
	claimsExp, err := time.ParseDuration(os.Getenv("USER_CLAIMS_EXPIRE"))
	if err != nil {
		return model.LoginUserResponse{}, ew.New(http.StatusInternalServerError, fmt.Errorf("user login:exp parse err:%s", err.Error()))
	}
	claims := model.NewUserClaims(user.UUID, claimsExp) 
	
	// getting JWT_KEY in .env
	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		return model.LoginUserResponse{}, ew.New(http.StatusInternalServerError, fmt.Errorf("user login:JWT_KEY is not set"))
	}
	// creating the jwt
	token, err := sdk_jwt.NewToken(claims, []byte(jwtKey))
	if err != nil {
		return model.LoginUserResponse{}, ew.New(http.StatusInternalServerError, fmt.Errorf("user login:create token fail:%s",err.Error()))
	}
	
	// give response
	res := model.LoginUserResponse{
		Authorization: "Bearer " + token,
	}
	return res, nil
}