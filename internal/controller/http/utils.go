package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

// translateError - приводит ошибку валидации к человекочитабельному виду
func translateError(err error, ts ut.Translator) error {
	if err == nil {
		return nil
	}
	var finalErr string
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		finalErr += fmt.Sprintf("%s", e.Translate(ts))
	}
	return fmt.Errorf("%s", finalErr)
}

// catchErrGin - логирует ошибку и отправляет статус код и сообщение клиенту
func catchErrGin(ctx *gin.Context, code int, msg string, err error) {
	if err == nil {
		logrus.Errorf("%s.", msg)
		ctx.AbortWithStatusJSON(code, gin.H{"error": msg})
		return
	}
	logrus.Errorf("%s. Error: %s", msg, err)
	ctx.AbortWithStatusJSON(code, gin.H{"error": msg})
}
