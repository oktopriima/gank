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
	"github.com/oktopriima/gank/app/response"
)

type InportRequest struct {
	ID        int     `json:"id"`
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
