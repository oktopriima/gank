/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 06/01/23, 11:26
 * Copyright (c) 2023
 */

package list

import (
	"context"
	"github.com/oktopriima/gank/app/middleware"
	"github.com/oktopriima/gank/app/models"
	"github.com/oktopriima/gank/app/usecase/API/carts"
	"gorm.io/gorm"
)

type listUsecase struct {
	cart carts.Outport
	db   *gorm.DB
}

func (l *listUsecase) Interactor(request *InportRequest, ctx context.Context) *InportResponse {
	res := new(InportResponse)

	userID, err := middleware.GetAuthenticatedUser(ctx)
	if err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	data, err := l.cart.List(l.db, models.CartParams{
		UserID:    userID,
		IsClaimed: false,
		Page:      request.Page,
		Size:      request.Size,
	}, ctx)
	if err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	var list []*List
	for _, datum := range data {
		list = append(list, &List{
			ID:          datum.ID,
			ProductID:   datum.ProductID,
			ProductName: datum.Products.Name,
			Quantity:    datum.Quantity,
			Price:       datum.Products.Price,
			TotalPrice:  datum.Quantity * datum.Products.Price,
		})
	}

	res.Success = true
	res.Data = list
	return res
}

func NewUsecase(cart carts.Outport, db *gorm.DB) Inport {
	return &listUsecase{
		cart: cart,
		db:   db,
	}
}
