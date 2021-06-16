package controllers

import (
	"first-api/app/serializers"
	"first-api/app/svc"
	"first-api/infra/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	uSvc svc.IUser
}

func NewUserController(grp interface{}, uSvc svc.IUser) {
	uc := &user{
		uSvc: uSvc,
	}
	g := grp.(*gin.Engine)
	userGrp := g.Group("/user-api")
	{
		userGrp.POST("/user", uc.CreateUser)
	}
}

func (ctr *user) CreateUser(c *gin.Context) {
	var us serializers.UserReq
	err := c.BindJSON(&us)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := ctr.uSvc.CreateUser(us)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (ctr *user) GetUser(c *gin.Context) {
	userId := c.Params.ByName("id")
	result, getErr := ctr.uSvc.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (ctr *user) UpdateUser(c *gin.Context) {
	userId := c.Params.ByName("id")
	result, getErr := ctr.uSvc.UpdateUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (ctr *user) DeleteUser(c *gin.Context) {
	userId := c.Params.ByName("id")
	result, getErr := ctr.uSvc.DeleteUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

// func GetUsers(c *gin.Context) {
// 	var user []model.User
// 	err := model.GetAllUsers(&user)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }

// func CreateUser(c *gin.Context) {
// 	var user model.User
// 	c.BindJSON(&user)
// 	err := model.CreateUser(&user)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }

// func GetUserById(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user model.User
// 	err := model.GetUserById(&user, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }

// func UpdateUser(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user model.User
// 	err := model.UpdateUser(&user, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }

// func DeleteUser(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user model.User
// 	err := model.DeleteUser(&user, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }
