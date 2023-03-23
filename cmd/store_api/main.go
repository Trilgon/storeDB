package main

import (
	"github.com/sirupsen/logrus"
	"store_api/internal/app"
	"store_api/internal/config"
)

func main() {
	err := config.InitViper()
	if err != nil {
		logrus.Panic(err)
	}
	server := app.NewStoreWebApi()
	err = server.StartApp()
	if err != nil {
		logrus.Panic(err)
	}
}
