package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ApiServer - http сервер на основе gin регистрирующий хэндлеры и запускающий http сервер
type ApiServer struct {
	router   *gin.Engine
	handlers *ApiHandlers
}

// NewApiServer - создание нового экземпляра ApiServer
func NewApiServer() *ApiServer {
	validate := validator.New()
	eng := en.New()
	uni := ut.New(eng, eng)
	ts, _ := uni.GetTranslator("en")
	err := enTranslations.RegisterDefaultTranslations(validate, ts)
	if err != nil {
		logrus.Panicf("[NewApiServer]: failed to RegisterDefaultTranslations. Error: %v", err)
	}
	return &ApiServer{
		router:   gin.Default(),
		handlers: NewApiHandlers(validate, ts),
	}
}

// RunHttpApi - регистрация хэндлеров и запуск http сервера
func (r ApiServer) RunHttpApi() error {
	api := r.router.Group("/api")
	{
		api.POST("/goods/add", r.handlers.GoodsAdd)
		api.GET("/goods/get", r.handlers.GoodsGet)
		api.PUT("/goods/update", r.handlers.GoodsUpdate)
		api.DELETE("/goods/delete", r.handlers.GoodsDelete)
		api.POST("/carts/create", r.handlers.CartCreate)
		api.PUT("/carts/goods/add", r.handlers.CartGoodsAdd)
		api.GET("/carts/goods/get", r.handlers.CartGoodsGet)
		api.DELETE("/carts/goods/delete", r.handlers.CartGoodsDelete)
		api.DELETE("/carts/delete", r.handlers.CartDelete)
		api.POST("/orders/create", r.handlers.OrderCreate)
		api.GET("/orders/get", r.handlers.OrderGet)
		api.PUT("/orders/update", r.handlers.OrderUpdate)
		api.DELETE("/orders/delete", r.handlers.OrderDelete)
	}
	err := r.router.Run(viper.GetString("server.host"))
	if err != nil {
		return fmt.Errorf("[RunHttpApi]: failed to run gin router. Error: %v", err)
	}
	return nil
}
