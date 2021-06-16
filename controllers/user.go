package controllers

import (
	"first-api/infra/errors"
	"first-api/serializers"
	"first-api/svc"
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
	//g := gin.Default()
	g := 
	userGrp := g.Group("/user-api")
	{
		userGrp.POST("user", uc.CreateUser)
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
