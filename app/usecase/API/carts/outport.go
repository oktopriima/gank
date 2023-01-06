/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 04/01/23, 16:23
 * Copyright (c) 2023
 */

package carts

import (
	"context"
	"github.com/oktopriima/gank/app/models"
	"gorm.io/gorm"
)

type Outport interface {
	Create(db *gorm.DB, cart *models.Cart, ctx context.Context) error
	List(db *gorm.DB, cart models.CartParams, ctx context.Context) ([]*models.Cart, error)
	FindBy(db *gorm.DB, params models.CartParams, ctx context.Context) (*models.Cart, error)
	Update(db *gorm.DB, cart *models.Cart, ctx context.Context) error
	Delete(db *gorm.DB, params models.CartParams, ctx context.Context) error
}
