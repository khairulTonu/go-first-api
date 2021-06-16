package main

import (
	"first-api/container"
	"first-api/infra/config"
	"first-api/infra/conn"

	"github.com/gin-gonic/gin"
)

func main() {

	// config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{
	// 	PrepareStmt: true,
	// })
	// if err != nil {
	// 	fmt.Println("Status:", err)
	// }
	// config.DB.AutoMigrate(&model.User{})

	config.LoadConfig()
	conn.ConnectDB()
	r := gin.Default()
	container.Init(r)

	r.Run()

}
