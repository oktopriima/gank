/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 06/01/23, 12:12
 * Copyright (c) 2023
 */

package update

import (
	"context"
	"fmt"
	"github.com/oktopriima/gank/app/middleware"
	"github.com/oktopriima/gank/app/models"
	"github.com/oktopriima/gank/app/usecase/API/carts"
	"github.com/oktopriima/gank/app/usecase/API/products"
	"gorm.io/gorm"
)

type updateUsecase struct {
	cart    carts.Outport
	product products.Outport
	db      *gorm.DB
}

func (u *updateUsecase) Interactor(request *InportRequest, ctx context.Context) *InportResponse {
	res := new(InportResponse)

	userID, err := middleware.GetAuthenticatedUser(ctx)
	if err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	// check product exist
	prd, err := u.product.FindBy(u.db, &models.ProductParams{
		ID: request.ProductID,
	}, ctx)
	if err != nil {
		res.ErrorMessage = fmt.Sprintf("product %s", err.Error())
		return res
	}

	// find cart by id
	cart, err := u.cart.FindBy(u.db, models.CartParams{
		ID:     request.ID,
		UserID: userID,
	}, ctx)
	if err != nil {
		res.ErrorMessage = fmt.Sprintf("cart %s", err.Error())
		return res
	}

	// assign update
	cart.ProductID = prd.ID
	cart.Quantity = request.Quantity

	// update cart
	if err = u.cart.Update(u.db, cart, ctx); err != nil {
		res.ErrorMessage = fmt.Sprintf("cart %s", err.Error())
		return res
	}

	res.Data = cart
	res.Success = true
	return res
}

func NewUsecase(outport carts.Outport, product products.Outport, db *gorm.DB) Inport {
	return &updateUsecase{
		cart:    outport,
		product: product,
		db:      db,
	}
}
