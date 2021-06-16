package svc

import (
	"first-api/infra/errors"
	"first-api/model"
	"first-api/serializers"
)

type IUser interface {
	CreateUser(serializers.UserReq) (*serializers.UserResp, *errors.RestErr)
	GetUser(userId string) (*model.User, *errors.RestErr)
}
