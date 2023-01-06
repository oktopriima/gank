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
	"github.com/oktopriima/gank/app/response"
)

type InportRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type InportResponse struct {
	Success bool `json:"success"`
	response.BaseResponse
}

type Inport interface {
	Interactor(request *InportRequest, ctx context.Context) *InportResponse
}
