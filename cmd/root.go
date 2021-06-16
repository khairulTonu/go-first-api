package cmd

import (
	"first-api/app/http"
	"first-api/infra/config"
	"first-api/infra/conn"
)

func Execute() {
	config.LoadConfig()
	conn.ConnectDB()
	http.Start()
}
