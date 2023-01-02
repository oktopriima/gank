/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 02/01/23, 14:54
 * Copyright (c) 2023
 */

package register

import (
	"context"
	"github.com/oktopriima/gank/app/response"
)

type InportRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type InportResponse struct {
	Success bool `json:"success"`
	response.BaseResponse
}

type Inport interface {
	Interactor(ctx context.Context, request InportRequest) *InportResponse
}
