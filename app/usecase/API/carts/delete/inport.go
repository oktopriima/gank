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
	"github.com/oktopriima/gank/app/response"
)

type InportRequest struct {
	ID int
}

type InportResponse struct {
	Success bool `json:"success"`
	response.BaseResponse
}

type Inport interface {
	Interactor(request *InportRequest, ctx context.Context) *InportResponse
}
