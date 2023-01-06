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
	"github.com/oktopriima/gank/app/response"
)

type InportRequest struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type InportResponse struct {
	Success bool `json:"success"`
	response.BaseResponse
}

type List struct {
	ID          int     `json:"id"`
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
	TotalPrice  float64 `json:"total_price"`
}

type Inport interface {
	Interactor(request *InportRequest, ctx context.Context) *InportResponse
}
