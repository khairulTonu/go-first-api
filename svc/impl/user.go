package impl

import (
	"encoding/json"
	"first-api/infra/errors"
	"first-api/model"
	"first-api/repository"
	"first-api/serializers"
	"first-api/svc"
)

type user struct {
	urepo repository.IUser
}

func NewUserService(urepo repository.IUser) svc.IUser {
	return &user{
		urepo: urepo,
	}
}

func (u *user) CreateUser(us serializers.UserReq) (*serializers.UserResp, *errors.RestErr) {
	var userObj model.User
	jsonData, _ := json.Marshal(us)
	_ = json.Unmarshal(jsonData, &userObj)
	userResult, createErr := u.urepo.Save(&userObj)
	if createErr != nil {
		return nil, createErr
	}
	var userResp serializers.UserResp
	jsonDataResp, _ := json.Marshal(userResult)
	_ = json.Unmarshal(jsonDataResp, &userResp)
	return &userResp, nil
}

func (u *user) GetUser(userId string) (*model.User, *errors.RestErr) {
	resp, getErr := u.urepo.Get(userId)
	if getErr != nil {
		return nil, getErr
	}
	return resp, nil
}
