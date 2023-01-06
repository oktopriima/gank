/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 06/01/23, 11:57
 * Copyright (c) 2023
 */

package list

import (
	"context"
	"github.com/oktopriima/gank/app/models"
	"github.com/oktopriima/gank/app/usecase/API/products"
	"gorm.io/gorm"
)

type listUsecase struct {
	product products.Outport
	db      *gorm.DB
}

func (l *listUsecase) Interactor(request *InportRequest, ctx context.Context) *InportResponse {
	res := new(InportResponse)

	data, err := l.product.List(l.db, &models.ProductParams{
		Name: request.Search,
		Page: request.Page,
		Size: request.Size,
	}, ctx)

	if err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	res.Data = data
	res.Success = true
	return res
}

func NewUsecase(product products.Outport,
	db *gorm.DB) Inport {
	return &listUsecase{
		product: product,
		db:      db,
	}
}
