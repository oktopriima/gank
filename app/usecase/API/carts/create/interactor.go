/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 04/01/23, 16:25
 * Copyright (c) 2023
 */

package create

import (
	"context"
	"fmt"
	"github.com/oktopriima/gank/app/middleware"
	"github.com/oktopriima/gank/app/models"
	"github.com/oktopriima/gank/app/usecase/API/carts"
	"github.com/oktopriima/gank/app/usecase/API/products"
	"gorm.io/gorm"
)

type createUsecase struct {
	cart    carts.Outport
	product products.Outport
	db      *gorm.DB
}

func NewUsecase(outport carts.Outport, product products.Outport, db *gorm.DB) Inport {
	return &createUsecase{
		cart:    outport,
		product: product,
		db:      db,
	}
}

func (c *createUsecase) Interactor(request *InportRequest, ctx context.Context) *InportResponse {
	res := new(InportResponse)

	userID, err := middleware.GetAuthenticatedUser(ctx)
	if err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	// check product exist
	prd, err := c.product.FindBy(c.db, &models.ProductParams{
		ID: request.ProductID,
	}, ctx)
	if err != nil {
		res.ErrorMessage = fmt.Sprintf("product %s", err.Error())
		return res
	}

	cart := new(models.Cart)
	cart.UserID = userID
	cart.ProductID = prd.ID
	cart.Quantity = request.Quantity

	// save cart
	if err = c.cart.Create(c.db, cart, ctx); err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	res.Data = cart
	res.Success = true
	return res
}
