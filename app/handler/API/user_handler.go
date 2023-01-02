/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 02/01/23, 15:30
 * Copyright (c) 2023
 */

package API

import (
	"github.com/labstack/echo"
	"github.com/oktopriima/gank/app/usecase/API/users/register"
	"net/http"
)

type UserHandler struct {
	register register.Inport
}

func NewUserHandler(register register.Inport) UserHandler {
	return UserHandler{register: register}
}

func (u *UserHandler) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := register.InportRequest{}
	res := new(register.InportResponse)

	err := c.Bind(&req)

	if err != nil {
		res.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res = u.register.Interactor(ctx, req)
	return c.JSON(http.StatusOK, res)
}
