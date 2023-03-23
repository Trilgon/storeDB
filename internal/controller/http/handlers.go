package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"store_api/internal/domain/models"
	"store_api/internal/domain/models/dto"
	"store_api/internal/domain/service"
	"strconv"
)

// ApiHandlers - структура хэндлеров для эндпоинтов ApiServer Store Web API
type ApiHandlers struct {
	service   service.StoreService
	validator *validator.Validate
	ts        ut.Translator
}

// NewApiHandlers - создание экземпляра структуры с хэндлерами для ApiServer
func NewApiHandlers(validate *validator.Validate, ts ut.Translator) *ApiHandlers {
	acceptance, err := service.NewStore()
	if err != nil {
		logrus.Panicf("[NewApiHandlers]: failed to create StoreService. %v", err)
	}
	return &ApiHandlers{
		service:   acceptance,
		validator: validate,
		ts:        ts,
	}
}

func (h *ApiHandlers) GoodsAdd(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Failed to read body", fmt.Errorf(
			"[GoodsAdd]: %v",
			err,
		))
		return
	}
	goods := models.Goods{}
	err = jsoniter.Unmarshal(body, &goods)
	if err != nil {
		catchErrGin(ctx, http.StatusUnprocessableEntity, "Failed to unmarshal body", fmt.Errorf(
			"[GoodsAdd]: %v",
			err,
		))
		return
	}

	err = h.validator.Struct(goods)
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("Body validation failed. %v", translatedErr),
			fmt.Errorf("[GoodsAdd]: %v", err),
		)
		return
	}

	err = h.service.GoodsAdd(&goods)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Request to DB doesn't succeed", fmt.Errorf(
			"[GoodsAdd]: %v",
			err,
		))
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (h *ApiHandlers) GoodsGet(ctx *gin.Context) {
	key := ctx.Request.URL.Query().Get("goods_id")
	if key == "" {
		catchErrGin(ctx, http.StatusBadRequest, "The goods_id in query required", fmt.Errorf(
			"[GoodsGet]: no value in query error"))
		return
	}
	goodsId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		catchErrGin(ctx, http.StatusBadRequest, "Failed to parse int64 goods_id from query", fmt.Errorf(
			"[GoodsGet]: %v", err,
		))
		return
	}

	err = h.validator.Var(goodsId, "required,gt=0")
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("The goods_id validation failed. %v", translatedErr),
			fmt.Errorf("[GoodsGet]: %v", err),
		)
		return
	}

	goods, err := h.service.GoodsGet(goodsId)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Request to DB doesn't succeed", fmt.Errorf(
			"[GoodsGet]: %v",
			err,
		))
		return
	}

	respBody, err := jsoniter.Marshal(&goods)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Failed to marshal response body", fmt.Errorf(
			"[GoodsGet]: %v",
			err,
		))
		return
	}
	_, err = ctx.Writer.Write(respBody)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Failed to write response body", fmt.Errorf(
			"[GoodsGet]: %v",
			err,
		))
		return
	}
}

func (h *ApiHandlers) GoodsUpdate(ctx *gin.Context) {
	key := ctx.Request.URL.Query().Get("goods_id")
	if key == "" {
		catchErrGin(ctx, http.StatusBadRequest, "The goods_id in query required", fmt.Errorf(
			"[GoodsUpdate]: no value in query error"))
		return
	}
	goodsId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		catchErrGin(ctx, http.StatusBadRequest, "Failed to parse int64 goods_id from query", fmt.Errorf(
			"[GoodsUpdate]: %v", err,
		))
		return
	}

	err = h.validator.Var(goodsId, "required,gt=0")
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("The goods_id validation failed. %v", translatedErr),
			fmt.Errorf("[GoodsUpdate]: %v", err),
		)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Failed to read body", fmt.Errorf(
			"[GoodsUpdate]: %v",
			err,
		))
		return
	}
	goods := dto.GoodsUpdate{}
	err = jsoniter.Unmarshal(body, &goods)
	if err != nil {
		catchErrGin(ctx, http.StatusUnprocessableEntity, "Failed to unmarshal body", fmt.Errorf(
			"[GoodsUpdate]: %v",
			err,
		))
		return
	}

	err = h.validator.Struct(goods)
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("Body validation failed. %v", translatedErr),
			fmt.Errorf("[GoodsUpdate]: %v", err),
		)
		return
	}

	err = h.service.GoodsUpdate(goodsId, &goods)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Request to DB doesn't succeed", fmt.Errorf(
			"[GoodsAdd]: %v",
			err,
		))
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (h *ApiHandlers) GoodsDelete(ctx *gin.Context) {
	key := ctx.Request.URL.Query().Get("goods_id")
	if key == "" {
		catchErrGin(ctx, http.StatusBadRequest, "The goods_id in query required", fmt.Errorf(
			"[GoodsDelete]: no value in query error"))
		return
	}
	goodsId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		catchErrGin(ctx, http.StatusBadRequest, "Failed to parse int64 goods_id from query", fmt.Errorf(
			"[GoodsDelete]: %v", err,
		))
		return
	}

	err = h.validator.Var(goodsId, "required,gt=0")
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("The goods_id validation failed. %v", translatedErr),
			fmt.Errorf("[GoodsDelete]: %v", err),
		)
		return
	}

	err = h.service.GoodsDelete(goodsId)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Request to DB doesn't succeed", fmt.Errorf(
			"[GoodsDelete]: %v",
			err,
		))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *ApiHandlers) CartCreate(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Failed to read body", fmt.Errorf(
			"[CartCreate]: %v",
			err,
		))
		return
	}
	cart := models.Cart{}
	err = jsoniter.Unmarshal(body, &cart)
	if err != nil {
		catchErrGin(ctx, http.StatusUnprocessableEntity, "Failed to unmarshal body", fmt.Errorf(
			"[CartCreate]: %v",
			err,
		))
		return
	}

	err = h.validator.Struct(cart)
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("Body validation failed. %v", translatedErr),
			fmt.Errorf("[CartCreate]: %v", err),
		)
		return
	}

	err = h.service.CartCreate(&cart)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Request to DB doesn't succeed", fmt.Errorf(
			"[CartCreate]: %v",
			err,
		))
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (h *ApiHandlers) CartGoodsAdd(ctx *gin.Context) {
	key := ctx.Request.URL.Query().Get("cart_id")
	if key == "" {
		catchErrGin(ctx, http.StatusBadRequest, "The cart_id in query required", fmt.Errorf(
			"[CartGoodsAdd]: no value in query error"))
		return
	}
	cartId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		catchErrGin(ctx, http.StatusBadRequest, "Failed to parse int64 cart_id from query", fmt.Errorf(
			"[CartGoodsAdd]: %v", err,
		))
		return
	}

	err = h.validator.Var(cartId, "required,gt=0")
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("The goods_id validation failed. %v", translatedErr),
			fmt.Errorf("[CartGoodsAdd]: %v", err),
		)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Failed to read body", fmt.Errorf(
			"[CartGoodsAdd]: %v",
			err,
		))
		return
	}
	goods := dto.GoodsAdd{}
	err = jsoniter.Unmarshal(body, &goods)
	if err != nil {
		catchErrGin(ctx, http.StatusUnprocessableEntity, "Failed to unmarshal body", fmt.Errorf(
			"[CartGoodsAdd]: %v",
			err,
		))
		return
	}

	err = h.validator.Struct(goods)
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("Body validation failed. %v", translatedErr),
			fmt.Errorf("[CartGoodsAdd]: %v", err),
		)
		return
	}

	err = h.service.CartAddGoods(&goods)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Request to DB doesn't succeed", fmt.Errorf(
			"[CartGoodsAdd]: %v",
			err,
		))
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (h *ApiHandlers) CartGoodsGet(ctx *gin.Context) {
	goods, err := h.service.CartGetGoods()
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Request to DB doesn't succeed", fmt.Errorf(
			"[CartGetGoods]: %v",
			err,
		))
		return
	}

	respBody, err := jsoniter.Marshal(&goods)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Failed to marshal response body", fmt.Errorf(
			"[CartGetGoods]: %v",
			err,
		))
		return
	}
	_, err = ctx.Writer.Write(respBody)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Failed to write response body", fmt.Errorf(
			"[CartGetGoods]: %v",
			err,
		))
		return
	}
}

func (h *ApiHandlers) CartGoodsUpdate(ctx *gin.Context) {
	key := ctx.Request.URL.Query().Get("cart_id")
	if key == "" {
		catchErrGin(ctx, http.StatusBadRequest, "The cart_id in query required", fmt.Errorf(
			"[CartGoodsUpdate]: no value in query error"))
		return
	}
	cartId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		catchErrGin(ctx, http.StatusBadRequest, "Failed to parse int64 cart_id from query", fmt.Errorf(
			"[CartGoodsUpdate]: %v", err,
		))
		return
	}

	key = ctx.Request.URL.Query().Get("goods_id")
	if key == "" {
		catchErrGin(ctx, http.StatusBadRequest, "The goods_id in query required", fmt.Errorf(
			"[CartGoodsUpdate]: no value in query error"))
		return
	}
	goodsId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		catchErrGin(ctx, http.StatusBadRequest, "Failed to parse int64 goods_id from query", fmt.Errorf(
			"[CartGoodsUpdate]: %v", err,
		))
		return
	}

	err = h.validator.Var(cartId, "required,gt=0")
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("The cart_id validation failed. %v", translatedErr),
			fmt.Errorf("[CartGoodsUpdate]: %v", err),
		)
		return
	}
	err = h.validator.Var(goodsId, "required,gt=0")
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("The goods_id validation failed. %v", translatedErr),
			fmt.Errorf("[CartGoodsUpdate]: %v", err),
		)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Failed to read body", fmt.Errorf(
			"[CartGoodsUpdate]: %v",
			err,
		))
		return
	}
	goods := struct {
		Quantity int64 `json:"quantity" validate:"required,gt=0"`
	}{}
	err = jsoniter.Unmarshal(body, &goods)
	if err != nil {
		catchErrGin(ctx, http.StatusUnprocessableEntity, "Failed to unmarshal body", fmt.Errorf(
			"[CartGoodsUpdate]: %v",
			err,
		))
		return
	}

	err = h.validator.Struct(goods)
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("Body validation failed. %v", translatedErr),
			fmt.Errorf("[CartGoodsUpdate]: %v", err),
		)
		return
	}

	err = h.service.CartGoodsUpdate(cartId, goodsId, goods.Quantity)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Request to DB doesn't succeed", fmt.Errorf(
			"[CartGoodsUpdate]: %v",
			err,
		))
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (h *ApiHandlers) CartGoodsDelete(ctx *gin.Context) {
	key := ctx.Request.URL.Query().Get("cart_id")
	if key == "" {
		catchErrGin(ctx, http.StatusBadRequest, "The cart_id in query required", fmt.Errorf(
			"[CartGoodsDelete]: no value in query error"))
		return
	}
	cartId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		catchErrGin(ctx, http.StatusBadRequest, "Failed to parse int64 cart_id from query", fmt.Errorf(
			"[CartGoodsDelete]: %v", err,
		))
		return
	}
	key = ctx.Request.URL.Query().Get("goods_id")
	if key == "" {
		catchErrGin(ctx, http.StatusBadRequest, "The goods_id in query required", fmt.Errorf(
			"[CartGoodsDelete]: no value in query error"))
		return
	}
	goodsId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		catchErrGin(ctx, http.StatusBadRequest, "Failed to parse int64 goods_id from query", fmt.Errorf(
			"[CartGoodsDelete]: %v", err,
		))
		return
	}

	err = h.validator.Var(cartId, "required,gt=0")
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("The cart_id validation failed. %v", translatedErr),
			fmt.Errorf("[CartGoodsDelete]: %v", err),
		)
		return
	}
	err = h.validator.Var(goodsId, "required,gt=0")
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("The goods_id validation failed. %v", translatedErr),
			fmt.Errorf("[CartGoodsDelete]: %v", err),
		)
		return
	}

	err = h.service.CartDeleteGoods(cartId, goodsId)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Request to DB doesn't succeed", fmt.Errorf(
			"[CartGoodsDelete]: %v",
			err,
		))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *ApiHandlers) CartDelete(ctx *gin.Context) {
	key := ctx.Request.URL.Query().Get("cart_id")
	if key == "" {
		catchErrGin(ctx, http.StatusBadRequest, "The cart_id in query required", fmt.Errorf(
			"[CartDelete]: no value in query error"))
		return
	}
	cartId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		catchErrGin(ctx, http.StatusBadRequest, "Failed to parse int64 cart_id from query", fmt.Errorf(
			"[CartDelete]: %v", err,
		))
		return
	}

	err = h.validator.Var(cartId, "required,gt=0")
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("The cart_id validation failed. %v", translatedErr),
			fmt.Errorf("[CartDelete]: %v", err),
		)
		return
	}

	err = h.service.CartDelete(cartId)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Request to DB doesn't succeed", fmt.Errorf(
			"[CartDelete]: %v",
			err,
		))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *ApiHandlers) OrderCreate(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Failed to read body", fmt.Errorf(
			"[OrderCreate]: %v",
			err,
		))
		return
	}
	cart := struct {
		CartId int64 `json:"cart_id" validate:"required,gt=0"`
	}{}
	err = jsoniter.Unmarshal(body, &cart)
	if err != nil {
		catchErrGin(ctx, http.StatusUnprocessableEntity, "Failed to unmarshal body", fmt.Errorf(
			"[OrderCreate]: %v",
			err,
		))
		return
	}

	err = h.validator.Struct(cart)
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("Body validation failed. %v", translatedErr),
			fmt.Errorf("[OrderCreate]: %v", err),
		)
		return
	}

	order, err := h.service.OrderCreate(cart.CartId)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Request to DB doesn't succeed", fmt.Errorf(
			"[OrderCreate]: %v",
			err,
		))
		return
	}

	respBody, err := jsoniter.Marshal(&order)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Failed to marshal response body", fmt.Errorf(
			"[OrderCreate]: %v",
			err,
		))
		return
	}
	_, err = ctx.Writer.Write(respBody)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Failed to write response body", fmt.Errorf(
			"[OrderCreate]: %v",
			err,
		))
		return
	}
}

func (h *ApiHandlers) OrderGet(ctx *gin.Context) {
	key := ctx.Request.URL.Query().Get("order_id")
	if key == "" {
		catchErrGin(ctx, http.StatusBadRequest, "The order_id in query required", fmt.Errorf(
			"[OrderGet]: no value in query error"))
		return
	}
	orderId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		catchErrGin(ctx, http.StatusBadRequest, "Failed to parse int64 order_id from query", fmt.Errorf(
			"[OrderGet]: %v", err,
		))
		return
	}

	err = h.validator.Var(orderId, "required,gt=0")
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("The order_id validation failed. %v", translatedErr),
			fmt.Errorf("[OrderGet]: %v", err),
		)
		return
	}

	order, err := h.service.OrderGet(orderId)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Request to DB doesn't succeed", fmt.Errorf(
			"[OrderGet]: %v",
			err,
		))
		return
	}

	respBody, err := jsoniter.Marshal(&order)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Failed to marshal response body", fmt.Errorf(
			"[OrderGet]: %v",
			err,
		))
		return
	}
	_, err = ctx.Writer.Write(respBody)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Failed to write response body", fmt.Errorf(
			"[OrderGet]: %v",
			err,
		))
		return
	}
}

func (h *ApiHandlers) OrderUpdate(ctx *gin.Context) {
	key := ctx.Request.URL.Query().Get("order_id")
	if key == "" {
		catchErrGin(ctx, http.StatusBadRequest, "The order_id in query required", fmt.Errorf(
			"[OrderUpdate]: no value in query error"))
		return
	}
	orderId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		catchErrGin(ctx, http.StatusBadRequest, "Failed to parse int64 order_id from query", fmt.Errorf(
			"[OrderUpdate]: %v", err,
		))
		return
	}

	err = h.validator.Var(orderId, "required,gt=0")
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("The OrderUpdate validation failed. %v", translatedErr),
			fmt.Errorf("[OrderUpdate]: %v", err),
		)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Failed to read body", fmt.Errorf(
			"[OrderUpdate]: %v",
			err,
		))
		return
	}
	order := dto.OrderUpdate{}
	err = jsoniter.Unmarshal(body, &order)
	if err != nil {
		catchErrGin(ctx, http.StatusUnprocessableEntity, "Failed to unmarshal body", fmt.Errorf(
			"[OrderUpdate]: %v",
			err,
		))
		return
	}

	err = h.validator.Struct(order)
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("Body validation failed. %v", translatedErr),
			fmt.Errorf("[OrderUpdate]: %v", err),
		)
		return
	}

	err = h.service.OrderUpdate(orderId, &order)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Request to DB doesn't succeed", fmt.Errorf(
			"[OrderUpdate]: %v",
			err,
		))
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (h *ApiHandlers) OrderDelete(ctx *gin.Context) {
	key := ctx.Request.URL.Query().Get("order_id")
	if key == "" {
		catchErrGin(ctx, http.StatusBadRequest, "The order_id in query required", fmt.Errorf(
			"[OrderDelete]: no value in query error"))
		return
	}
	orderId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		catchErrGin(ctx, http.StatusBadRequest, "Failed to parse int64 order_id from query", fmt.Errorf(
			"[OrderDelete]: %v", err,
		))
		return
	}

	err = h.validator.Var(orderId, "required,gt=0")
	if err != nil {
		translatedErr := translateError(err, h.ts)
		catchErrGin(
			ctx,
			http.StatusBadRequest,
			fmt.Sprintf("The order_id validation failed. %v", translatedErr),
			fmt.Errorf("[OrderDelete]: %v", err),
		)
		return
	}

	err = h.service.OrderDelete(orderId)
	if err != nil {
		catchErrGin(ctx, http.StatusInternalServerError, "Request to DB doesn't succeed", fmt.Errorf(
			"[OrderDelete]: %v",
			err,
		))
		return
	}

	ctx.Status(http.StatusNoContent)
}
