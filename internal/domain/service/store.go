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
	OrderCreate(cartId int64) (*models.Order, error)
	// OrderGet - получение информации о заказе
	OrderGet(orderId int64) (*models.Order, error)
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
	err := s.rep.GoodsAdd(goods)
	if err != nil {
		return fmt.Errorf("[GoodsAdd]: %s", err)
	}
	return nil
}

func (s *Store) GoodsGet(goodsId int64) (*models.Goods, error) {
	goods, err := s.rep.GoodsGet(goodsId)
	if err != nil {
		return nil, fmt.Errorf("[GoodsGet]: %s", err)
	}
	return goods, nil
}

func (s *Store) GoodsUpdate(goodsId int64, goods *dto.GoodsUpdate) error {
	err := s.rep.GoodsUpdate(goodsId, goods)
	if err != nil {
		return fmt.Errorf("[GoodsUpdate]: %s", err)
	}
	return nil
}

func (s *Store) GoodsDelete(goodsId int64) error {
	err := s.rep.GoodsDelete(goodsId)
	if err != nil {
		return fmt.Errorf("[GoodsDelete]: %s", err)
	}
	return nil
}

func (s *Store) CartCreate(cart *models.Cart) error {
	err := s.rep.CartCreate(cart)
	if err != nil {
		return fmt.Errorf("[CartCreate]: %s", err)
	}
	return nil
}

func (s *Store) CartAddGoods(goods *dto.GoodsAdd) error {
	err := s.rep.CartAddGoods(goods)
	if err != nil {
		return fmt.Errorf("[CartAddGoods]: %s", err)
	}
	return nil
}

func (s *Store) CartGetGoods() ([]models.Goods, error) {
	goods, err := s.rep.CartGetGoods()
	if err != nil {
		return nil, fmt.Errorf("[CartGetGoods]: %s", err)
	}
	return goods, nil
}

func (s *Store) CartGoodsUpdate(cartId, goodsId, quantity int64) error {
	err := s.rep.CartGoodsUpdate(cartId, goodsId, quantity)
	if err != nil {
		return fmt.Errorf("[CartGoodsUpdate]: %s", err)
	}
	return nil
}

func (s *Store) CartDeleteGoods(cartId, goodsId int64) error {
	err := s.rep.CartDeleteGoods(cartId, goodsId)
	if err != nil {
		return fmt.Errorf("[CartDeleteGoods]: %s", err)
	}
	return nil
}

func (s *Store) CartDelete(cartId int64) error {
	err := s.rep.CartDelete(cartId)
	if err != nil {
		return fmt.Errorf("[CartDelete]: %s", err)
	}
	return nil
}

func (s *Store) OrderCreate(cartId int64) (*models.Order, error) {
	order, err := s.rep.OrderCreate(cartId)
	if err != nil {
		return nil, fmt.Errorf("[OrderCreate]: %s", err)
	}
	return order, nil
}

func (s *Store) OrderGet(orderId int64) (*models.Order, error) {
	order, err := s.rep.OrderGet(orderId)
	if err != nil {
		return nil, fmt.Errorf("[OrderGet]: %s", err)
	}
	return order, nil
}

func (s *Store) OrderUpdate(orderId int64, order *dto.OrderUpdate) error {
	err := s.rep.OrderUpdate(orderId, order)
	if err != nil {
		return fmt.Errorf("[OrderUpdate]: %s", err)
	}
	return nil
}

func (s *Store) OrderDelete(orderId int64) error {
	err := s.rep.OrderDelete(orderId)
	if err != nil {
		return fmt.Errorf("[OrderDelete]: %s", err)
	}
	return nil
}
