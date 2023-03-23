package service

import (
	"fmt"
	"store_api/internal/domain/models"
	"store_api/internal/domain/models/dto"
	"store_api/internal/repository"
	"store_api/internal/repository/postgresql"
)

// StoreService - сервисная логика взаимодействия с репозиторием приёмки
type StoreService interface {
	// GoodsAdd - добавление товара
	GoodsAdd(goods *models.Goods) error
	// GoodsGet - получение информации о товаре, возвращает товар
	GoodsGet(goodsId int64) (*models.Goods, error)
	// GoodsUpdate - обновление информации о товаре
	GoodsUpdate(goodsId int64, goods *dto.GoodsUpdate) error
	// GoodsDelete - удаление товара
	GoodsDelete(goodsId int64) error
	// CartCreate - создание корзины
	CartCreate(cart *models.Cart) error
	// CartAddGoods - добавление товара в корзину
	CartAddGoods(goods *dto.GoodsAdd) error
	// CartGetGoods - получение списка товаров в корзине
	CartGetGoods() ([]models.Goods, error)
	// CartGoodsUpdate - обновление информации о товаре в корзине
	CartGoodsUpdate(cartId, goodsId, quantity int64) error
	// CartDeleteGoods - удаление товара из корзины
	CartDeleteGoods(cartId, goodsId int64) error
	// CartDelete - удаление корзины
	CartDelete(cartId int64) error
	// OrderCreate - оформление заказа на основе корзины
	OrderCreate(cartId int64) (models.Order, error)
	// OrderGet - получение информации о заказе
	OrderGet(orderId int64) (models.Order, error)
	// OrderUpdate - обновление информации о заказе
	OrderUpdate(orderId int64, order *dto.OrderUpdate) error
	// OrderDelete - Получение списка товаров в корзине
	OrderDelete(orderId int64) error
}

type Store struct {
	rep repository.StoreRepository
}

func NewStore() (*Store, error) {
	storeRep, err := postgresql.NewStoreRepository()
	if err != nil {
		return nil, fmt.Errorf("[NewStore]: %v", err)
	}
	return &Store{rep: storeRep}, nil
}

func (s *Store) GoodsAdd(goods *models.Goods) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) GoodsGet(goodsId int64) (*models.Goods, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) GoodsUpdate(goodsId int64, goods *dto.GoodsUpdate) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) GoodsDelete(goodsId int64) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) CartCreate(cart *models.Cart) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) CartAddGoods(goods *dto.GoodsAdd) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) CartGetGoods() ([]models.Goods, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) CartGoodsUpdate(cartId, goodsId, quantity int64) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) CartDeleteGoods(cartId, goodsId int64) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) CartDelete(cartId int64) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) OrderCreate(cartId int64) (models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) OrderGet(orderId int64) (models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) OrderUpdate(orderId int64, order *dto.OrderUpdate) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) OrderDelete(orderId int64) error {
	//TODO implement me
	panic("implement me")
}
