package impl

import (
	"first-api/infra/errors"
	"first-api/model"
	"first-api/repository"
	"log"

	"gorm.io/gorm"
)

type user struct {
	*gorm.DB
}

// NewMySqlUserRepository will create an object that represent the Company.Repository implementations
func NewMySqlUserRepository(db *gorm.DB) repository.IUser {
	return &user{
		DB: db,
	}
}

func (db *user) Save(user *model.User) (*model.User, *errors.RestErr) {
	if err := db.Model(&model.User{}).Create(&user).Error; err != nil {
		log.Println("error occurred while creating user ", err)
		restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
		return nil, restErr
	}
	return user, nil
}

func (db *user) Get(userId string) (*model.User, *errors.RestErr) {
	var user model.User
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		log.Println("error occurred while getting user ", err)
		restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
		return nil, restErr
	}
	return &user, nil
}

func (db *user) Update(userId string) (*model.User, *errors.RestErr) {
	var user model.User
	if err := db.Where("id = ?", userId).Save(&user).Error; err != nil {
		log.Println("error occurred while getting user ", err)
		restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
		return nil, restErr
	}
	return &user, nil
}

func (db *user) Delete(userId string) (*model.User, *errors.RestErr) {
	var user model.User
	if err := db.Where("id = ?", userId).Delete(&user).Error; err != nil {
		log.Println("error occurred while getting user ", err)
		restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
		return nil, restErr
	}
	return &user, nil
}
