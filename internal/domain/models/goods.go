package models

// Goods - товар в магазине
type Goods struct {
	GoodsId  int64  `json:"goods_id" db:"goods_id" validate:"required,gt=0"`
	Name     string `json:"name" db:"name" validate:"required"`
	Price    string `json:"price" db:"price" validate:"required,gt=0"`
	Quantity int64  `json:"quantity" db:"quantity" validate:"required,gte=0"`
}
