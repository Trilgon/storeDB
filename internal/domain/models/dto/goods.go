package dto

// GoodsUpdate - данные для обновления информации о товаре в магазине
type GoodsUpdate struct {
	Name     string `json:"name" db:"name" validate:"required"`
	Price    string `json:"price" db:"price" validate:"required,gt=0"`
	Quantity int64  `json:"quantity" db:"quantity" validate:"required,gte=0"`
}

// GoodsAdd - данные для добавления товара в корзину
type GoodsAdd struct {
	GoodsId  int64 `json:"goods_id" db:"goods_id" validate:"required,gt=0"`
	Quantity int64 `json:"quantity" db:"quantity" validate:"required,gt=0"`
}
