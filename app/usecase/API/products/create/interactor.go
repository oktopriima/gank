/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 06/01/23, 11:12
 * Copyright (c) 2023
 */

package create

import (
	"context"
	"github.com/oktopriima/gank/app/models"
	"github.com/oktopriima/gank/app/usecase/API/products"
	"gorm.io/gorm"
)

type createUsecase struct {
	product products.Outport
	db      *gorm.DB
}

func (c *createUsecase) Interactor(request *InportRequest, ctx context.Context) *InportResponse {
	res := new(InportResponse)

	prd := new(models.Product)
	prd.Name = request.Name
	prd.Description = request.Description
	prd.Price = request.Price

	// create product
	if err := c.product.Create(c.db, prd, ctx); err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	res.Success = true
	res.Data = prd
	return res
}

func NewUsecase(outport products.Outport, db *gorm.DB) Inport {
	return &createUsecase{
		product: outport,
		db:      db,
	}
}
