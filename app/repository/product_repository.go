/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 06/01/23, 11:03
 * Copyright (c) 2023
 */

package repository

import (
	"context"
	"fmt"
	"github.com/oktopriima/gank/app/models"
	"github.com/oktopriima/gank/app/usecase/API/products"
	"gorm.io/gorm"
)

type productRepository struct {
}

func (p *productRepository) List(db *gorm.DB, params *models.ProductParams, ctx context.Context) ([]*models.Product, error) {
	var data []*models.Product
	res := db.WithContext(ctx)

	offset := (params.Page - 1) * params.Size

	if params.Name != "" {
		res = res.Where("name like ?", fmt.Sprintf("%%%s%%", params.Name))
	}

	res = res.Offset(offset).
		Limit(params.Size).Find(&data)
	if err := res.Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (p *productRepository) Create(db *gorm.DB, product *models.Product, ctx context.Context) error {
	err := db.WithContext(ctx).Create(&product).Error
	return err
}

func (p *productRepository) FindBy(db *gorm.DB, params *models.ProductParams, ctx context.Context) (*models.Product, error) {
	data := new(models.Product)
	res := db.WithContext(ctx)

	if params.ID != 0 {
		res = res.Where("id", params.ID)
	}

	if params.Name != "" {
		res = res.Where("name like ?", fmt.Sprintf("%%%s%%", params.Name))
	}

	res = res.First(&data)
	if err := res.Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewProductRepository() products.Outport {
	return &productRepository{}
}
