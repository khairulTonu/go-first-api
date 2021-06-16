package impl

import (
	"encoding/json"
	"first-api/app/model"
	"first-api/app/repository"
	"first-api/app/serializers"
	"first-api/app/svc"
	"first-api/infra/errors"
)

type user struct {
	urepo repository.IUser
}

func NewUserService(urepo repository.IUser) svc.IUser {
	return &user{
		urepo: urepo,
	}
}

func (u *user) GetAllUsers() (*[]model.User, *errors.RestErr) {
	resp, getErr := u.urepo.GetAll()
	if getErr != nil {
		return nil, getErr
	}
	return resp, nil
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

func (u *user) UpdateUser(userId string) (*model.User, *errors.RestErr) {
	resp, getErr := u.urepo.Update(userId)
	if getErr != nil {
		return nil, getErr
	}
	return resp, nil
}

func (u *user) DeleteUser(userId string) (*model.User, *errors.RestErr) {
	resp, getErr := u.urepo.Delete(userId)
	if getErr != nil {
		return nil, getErr
	}
	return resp, nil
}
