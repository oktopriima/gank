/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 04/01/23, 16:24
 * Copyright (c) 2023
 */

package create

import (
	"context"
	"github.com/oktopriima/gank/app/response"
)

type InportRequest struct {
	ProductID int     `json:"product_id"`
	Quantity  float64 `json:"quantity"`
}

type InportResponse struct {
	Success bool `json:"success"`
	response.BaseResponse
}

type Inport interface {
	Interactor(request *InportRequest, ctx context.Context) *InportResponse
}
