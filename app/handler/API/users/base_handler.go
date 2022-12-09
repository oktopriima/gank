/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 14/10/22, 14:31
 * Copyright (c) 2022
 */

package users

import (
	"github.com/labstack/echo"
)

type baseHandler struct {
}

type BaseHandler interface {
	SelfHandler(ctx echo.Context) error
}

func NewBaseHandler() BaseHandler {
	return &baseHandler{}
}
