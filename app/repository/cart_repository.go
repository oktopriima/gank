/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 04/01/23, 16:45
 * Copyright (c) 2023
 */

package repository

import (
	"context"
	"github.com/oktopriima/gank/app/models"
	"github.com/oktopriima/gank/app/usecase/API/carts"
	"gorm.io/gorm"
)

type cartRepository struct {
}

func (c *cartRepository) Delete(db *gorm.DB, params models.CartParams, ctx context.Context) error {
	res := db.WithContext(ctx).
		Where("id", params.ID).
		Where("is_claimed", params.IsClaimed).
		Delete(&models.Cart{})

	return res.Error
}

func (c *cartRepository) FindBy(db *gorm.DB, params models.CartParams, ctx context.Context) (*models.Cart, error) {
	data := new(models.Cart)
	res := db.WithContext(ctx)

	if params.ID != 0 {
		res = res.Where("id", params.ID)
	}

	if params.UserID != 0 {
		res = res.Where("user_id", params.UserID)
	}

	if params.ProductID != 0 {
		res = res.Where("product_id", params.ProductID)
	}

	res = res.First(&data)
	if err := res.Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (c *cartRepository) Update(db *gorm.DB, cart *models.Cart, ctx context.Context) error {
	err := db.WithContext(ctx).Save(cart).Error
	return err
}

func (c *cartRepository) List(db *gorm.DB, cart models.CartParams, ctx context.Context) ([]*models.Cart, error) {
	var data []*models.Cart
	res := db.WithContext(ctx)

	offset := (cart.Page - 1) * cart.Size

	if cart.UserID != 0 {
		res = res.Where("user_id", cart.UserID)
	}

	if cart.ProductID != 0 {
		res = res.Where("product_id", cart.ProductID)
	}

	res = res.Where("is_claimed", cart.IsClaimed).
		Offset(offset).
		Limit(cart.Size)

	res = res.Preload("Products").Find(&data)
	if err := res.Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (c *cartRepository) Create(db *gorm.DB, cart *models.Cart, ctx context.Context) error {
	err := db.WithContext(ctx).Create(&cart).Error
	return err
}

func NewCartRepository() carts.Outport {
	return &cartRepository{}
}
