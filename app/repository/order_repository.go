/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 06/01/23, 13:02
 * Copyright (c) 2023
 */

package repository

import (
	"context"
	"github.com/oktopriima/gank/app/models"
	"github.com/oktopriima/gank/app/usecase/API/orders"
	"gorm.io/gorm"
)

type orderRepository struct {
}

func (o *orderRepository) Create(db *gorm.DB, order *models.Order, ctx context.Context) error {
	err := db.WithContext(ctx).Create(&order).Error
	return err
}

func (o *orderRepository) List(db *gorm.DB, params models.OrderParams, ctx context.Context) ([]*models.Order, error) {
	var data []*models.Order
	offset := (params.Page - 1) * params.Size

	res := db.WithContext(ctx)

	if params.ID != 0 {
		res = res.Where("id", params.ID)
	}

	if params.UserID != 0 {
		res = res.Where("user_id", params.UserID)
	}

	if params.OrderCode != "" {
		res = res.Where("order_code", params.OrderCode)
	}

	res = res.Offset(offset).
		Limit(params.Size).
		Preload("OrderDetails", func(db2 *gorm.DB) *gorm.DB {
			return db2.Preload("Product")
		}).
		Find(&data)
	if err := res.Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (o *orderRepository) FindBy(db *gorm.DB, params models.OrderParams, ctx context.Context) (*models.Order, error) {
	data := new(models.Order)
	offset := (params.Page - 1) * params.Size

	res := db.WithContext(ctx)

	if params.ID != 0 {
		res = res.Where("id", params.ID)
	}

	if params.UserID != 0 {
		res = res.Where("user_id", params.UserID)
	}

	if params.OrderCode != "" {
		res = res.Where("order_code", params.OrderCode)
	}

	res = res.Offset(offset).
		Limit(params.Size).
		Preload("OrderDetails", func(db2 *gorm.DB) *gorm.DB {
			return db2.Preload("Product")
		}).
		First(&data)
	if err := res.Error; err != nil {
		return nil, err
	}

	return data, nil
}

func NewOrderRepository() orders.Outport {
	return &orderRepository{}
}
