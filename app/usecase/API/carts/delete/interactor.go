/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 06/01/23, 12:28
 * Copyright (c) 2023
 */

package delete

import (
	"context"
	"github.com/oktopriima/gank/app/models"
	"github.com/oktopriima/gank/app/usecase/API/carts"
	"gorm.io/gorm"
)

type deleteUsecase struct {
	cart carts.Outport
	db   *gorm.DB
}

func (d *deleteUsecase) Interactor(request *InportRequest, ctx context.Context) *InportResponse {
	res := new(InportResponse)

	if err := d.cart.Delete(d.db, models.CartParams{
		ID:        request.ID,
		IsClaimed: false,
	}, ctx); err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	res.Success = true
	return res
}

func NewUsecase(cart carts.Outport, db *gorm.DB) Inport {
	return &deleteUsecase{
		cart: cart,
		db:   db,
	}
}
