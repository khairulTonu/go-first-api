package svc

import (
	"first-api/app/model"
	"first-api/app/serializers"
	"first-api/infra/errors"
)

type IUser interface {
	GetAllUsers() (*[]model.User, *errors.RestErr)
	CreateUser(serializers.UserReq) (*serializers.UserResp, *errors.RestErr)
	GetUser(userId string) (*model.User, *errors.RestErr)
	UpdateUser(userId string) (*model.User, *errors.RestErr)
	DeleteUser(userId string) (*model.User, *errors.RestErr)
}
