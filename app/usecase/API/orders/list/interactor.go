/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 06/01/23, 13:21
 * Copyright (c) 2023
 */

package list

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/oktopriima/gank/app/middleware"
	"github.com/oktopriima/gank/app/models"
	"github.com/oktopriima/gank/app/usecase/API/orders"
	"github.com/oktopriima/gank/app/usecase/API/products"
	"gorm.io/gorm"
)

type listUsecase struct {
	order   orders.Outport
	product products.Outport
	db      *gorm.DB
}

func (l *listUsecase) Interactor(request *InportRequest, ctx context.Context) *InportResponse {
	res := new(InportResponse)
	userID, err := middleware.GetAuthenticatedUser(ctx)
	if err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	var list []*List
	data, err := l.order.List(l.db, models.OrderParams{
		UserID: userID,
		Page:   request.Page,
		Size:   request.Size,
	}, ctx)
	if err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	if err := copier.Copy(&list, &data); err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	res.Data = list
	res.Success = true
	return res
}

func NewUsecase(
	order orders.Outport,
	product products.Outport,
	db *gorm.DB,
) Inport {
	return &listUsecase{
		order:   order,
		product: product,
		db:      db,
	}
}
