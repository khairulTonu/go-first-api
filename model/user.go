package model

import "first-api/config"

func GetAllUsers(user *[]User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(user *User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}
