package repository

import (
	"context"
	"go-template/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (err error)
	FindByID(ctx context.Context, id uint) (user entity.User, err error)
	FindByUUID(ctx context.Context, uuid string) (user entity.User, err error)
	FindByEmail(ctx context.Context, email string) (user entity.User, err error)
}

func NewUserRepository(db *gorm.DB) (u UserRepository){
	return &userRepositoryImpl{
		db: db,
	}
}
// Must implement all methods from UserRepository interface.
type userRepositoryImpl struct {
	db *gorm.DB
}


func (r *userRepositoryImpl) Create(ctx context.Context, user *entity.User) (err error){
	err = r.db.WithContext(ctx).Create(user).Error
	return
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id uint) (user entity.User, err error){
	err = r.db.WithContext(ctx).First(&user, id).Error
	if err == gorm.ErrRecordNotFound {
		return entity.User{}, err
	}
	return
}

func (r *userRepositoryImpl) FindByUUID(ctx context.Context, uuid string) (user entity.User, err error){
	err = r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return entity.User{}, err
	}
	return
}

func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (user entity.User, err error) {
	err = r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return entity.User{}, err
	}
	return 
}