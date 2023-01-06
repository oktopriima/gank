/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 06/01/23, 12:41
 * Copyright (c) 2023
 */

package orders

import (
	"context"
	"github.com/oktopriima/gank/app/models"
	"gorm.io/gorm"
)

type Outport interface {
	Create(db *gorm.DB, order *models.Order, ctx context.Context) error
	List(db *gorm.DB, params models.OrderParams, ctx context.Context) ([]*models.Order, error)
	FindBy(db *gorm.DB, params models.OrderParams, ctx context.Context) (*models.Order, error)
}
