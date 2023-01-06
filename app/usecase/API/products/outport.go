/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 06/01/23, 11:02
 * Copyright (c) 2023
 */

package products

import (
	"context"
	"github.com/oktopriima/gank/app/models"
	"gorm.io/gorm"
)

type Outport interface {
	FindBy(db *gorm.DB, params *models.ProductParams, ctx context.Context) (*models.Product, error)
	List(db *gorm.DB, params *models.ProductParams, ctx context.Context) ([]*models.Product, error)
	Create(db *gorm.DB, product *models.Product, ctx context.Context) error
}
