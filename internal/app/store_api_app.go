package app

import (
	"store_api/internal/controller/http"
)

type StoreWebApiApp struct {
	router *http.ApiServer
}

func NewStoreWebApi() *StoreWebApiApp {
	return &StoreWebApiApp{router: http.NewApiServer()}
}

func (a StoreWebApiApp) StartApp() error {
	err := a.router.RunHttpApi()
	if err != nil {
		return err
	}
	return nil
}
