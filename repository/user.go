package repository

import (
	"first-api/infra/errors"
	"first-api/model"
)

type IUser interface {
	Save(user *model.User) (*model.User, *errors.RestErr)
	Get(userId string) (*model.User, *errors.RestErr)
	Update(userId string) (*model.User, *errors.RestErr)
	Delete(userId string) (*model.User, *errors.RestErr)
}
