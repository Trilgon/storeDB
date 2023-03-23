package models

import "time"

// Order - заказ в магазине
type Order struct {
	OrderId    int64      `json:"order_id" db:"order_id" validate:"required,gt=0"`
	GoodsId    int64      `json:"goods_id" db:"goods_id" validate:"required,gt=0"`
	Quantity   int64      `json:"quantity" db:"quantity" validate:"required,gt=0"`
	Total      int64      `json:"total" db:"total" validate:"required,gt=0"`
	OrderTime  *time.Time `json:"order_time" db:"order_time" validate:"omitempty"`
	FinishTime time.Time  `json:"finish_time" db:"finish_time" validate:"required"`
}
