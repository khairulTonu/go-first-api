package http

import (
	"first-api/app/container"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	container.Init(r)

	r.Run()
}
