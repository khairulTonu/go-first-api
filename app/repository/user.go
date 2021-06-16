package repository

import (
	"first-api/app/model"
	"first-api/infra/errors"
)

type IUser interface {
	GetAll() (*[]model.User, *errors.RestErr)
	Save(user *model.User) (*model.User, *errors.RestErr)
	Get(userId string) (*model.User, *errors.RestErr)
	Update(userId string) (*model.User, *errors.RestErr)
	Delete(userId string) (*model.User, *errors.RestErr)
}
