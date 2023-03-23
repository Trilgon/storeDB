package models

// Cart - корзина в магазине
type Cart struct {
	CartId   int64 `json:"cart_id" db:"cart_id" validate:"required,gt=0"`
	GoodsId  int64 `json:"goods_id" db:"goods_id" validate:"required,gt=0"`
	Quantity int64 `json:"quantity" db:"quantity" validate:"required,gt=0"`
	Total    int64 `json:"total" db:"total" validate:"required,gt=0"`
}
