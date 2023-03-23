package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"store_api/internal/domain/models"
	"store_api/internal/domain/models/dto"
)

func NewStoreRepository() (*StoreRepository, error) {
	db, err := getDB()
	if err != nil {
		return nil, fmt.Errorf("[NewStoreRepository]: %v", err)
	}
	return &StoreRepository{db: db}, nil
}

type StoreRepository struct {
	db *sqlx.DB
}

func (r *StoreRepository) GoodsAdd(goods *models.Goods) error {
	_, err := r.db.NamedExec(`INSERT INTO goods (goods_id, name, price, quantity) VALUES (:goods_id, :name, :price, :quantity)`, goods)
	if err != nil {
		return errors.Wrap(err, "failed to add goods")
	}
	return nil
}

func (r *StoreRepository) GoodsGet(goodsId int64) (*models.Goods, error) {
	goods := &models.Goods{}
	err := r.db.Get(goods, `SELECT * FROM goods WHERE goods_id=$1`, goodsId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrapf(err, "failed to get goods with id %d", goodsId)
		}
		return nil, errors.Wrap(err, "failed to get goods")
	}
	return goods, nil
}

func (r *StoreRepository) GoodsUpdate(goodsId int64, goods *dto.GoodsUpdate) error {
	_, err := r.db.Exec(
		`UPDATE goods SET name=:$1, price=$2, quantity=$3 WHERE goods_id=$4`,
		goods.Name,
		goods.Price,
		goods.Quantity,
		goodsId,
	)
	if err != nil {
		return errors.Wrap(err, "failed to update goods")
	}
	return nil
}

func (r *StoreRepository) GoodsDelete(goodsId int64) error {
	_, err := r.db.Exec(`DELETE FROM goods WHERE goods_id=$1`, goodsId)
	if err != nil {
		return errors.Wrapf(err, "failed to delete goods with id %d", goodsId)
	}
	return nil
}

func (r *StoreRepository) CartCreate(cart *models.Cart) error {
	_, err := r.db.NamedExec(`INSERT INTO carts (cart_id, goods_id, quantity, total) VALUES (:cart_id, :goods_id, :quantity, :total)`, cart)
	if err != nil {
		return errors.Wrap(err, "failed to create cart")
	}
	return nil
}

func (r *StoreRepository) CartAddGoods(goods *dto.GoodsAdd) error {
	_, err := r.db.NamedExec(`INSERT INTO carts (cart_id, goods_id, quantity, total) VALUES (:cart_id, :goods_id, :quantity, :total) ON CONFLICT ON CONSTRAINT carts_pkey DO UPDATE SET quantity = carts.quantity + :quantity, total = carts.total + :total`, goods)
	if err != nil {
		return errors.Wrap(err, "failed to add goods to cart")
	}
	return nil
}

func (r *StoreRepository) CartGetGoods() ([]models.Goods, error) {
	goods := make([]models.Goods, 0)
	err := r.db.Select(&goods, `SELECT g.* FROM goods g JOIN carts c ON g.goods_id=c.goods_id`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get goods from cart")
	}
	return goods, nil
}

func (r *StoreRepository) CartGoodsUpdate(cartId, goodsId, quantity int64) error {
	query := `UPDATE cart_goods SET quantity = $1 WHERE cart_id = $2 AND goods_id = $3`
	_, err := r.db.Exec(query, quantity, cartId, goodsId)
	if err != nil {
		return err
	}

	return nil
}

func (r *StoreRepository) CartDeleteGoods(cartId, goodsId int64) error {
	_, err := r.db.Exec(`DELETE FROM carts WHERE cart_id = $1 AND goods_id = $2`, cartId, goodsId)
	if err != nil {
		return fmt.Errorf("failed to delete goods from cart: %w", err)
	}

	return nil
}

func (r *StoreRepository) CartDelete(cartId int64) error {
	_, err := r.db.Exec(`DELETE FROM carts WHERE cart_id = $1`, cartId)
	if err != nil {
		return fmt.Errorf("failed to delete cart: %w", err)
	}

	return nil
}

func (r *StoreRepository) OrderCreate(cartId int64) (models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (r *StoreRepository) OrderGet(orderId int64) (models.Order, error) {
	order := models.Order{}
	err := r.db.Get(&order, `
		SELECT order_id, goods_id, quantity, total, order_time, finish_time
		FROM orders WHERE order_id = $1
	`, orderId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Order{}, fmt.Errorf("order with ID %d not found in DB", orderId)
		}
		return models.Order{}, fmt.Errorf("failed to get order: %w", err)
	}

	return order, nil
}

func (r *StoreRepository) OrderUpdate(orderId int64, order *dto.OrderUpdate) error {
	//TODO implement me
	panic("implement me")
}

func (r *StoreRepository) OrderDelete(orderId int64) error {
	// Check if order with provided ID exists
	var count int64
	err := r.db.Get(&count, `SELECT count(*) FROM orders WHERE order_id = $1`, orderId)
	if err != nil {
		return fmt.Errorf("failed to check order in DB: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("order with ID %d not found in DB", orderId)
	}

	_, err = r.db.Exec(`DELETE FROM orders WHERE order_id = $1`, orderId)
	if err != nil {
		return fmt.Errorf("failed to delete order from DB: %w", err)
	}

	return nil
}
