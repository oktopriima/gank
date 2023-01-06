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
	OrderCode    string  `json:"order_code"`
	FinalPrice   float64 `json:"final_price"`
	OrderDetails []struct {
		Quantity   float64 `json:"quantity"`
		Price      float64 `json:"price"`
		TotalPrice float64 `json:"total_price"`
		Product    struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"product"`
	} `json:"order_details"`
}

type Inport interface {
	Interactor(request *InportRequest, ctx context.Context) *InportResponse
}
