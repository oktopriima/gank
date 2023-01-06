/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 06/01/23, 13:04
 * Copyright (c) 2023
 */

package API

import (
	"github.com/labstack/echo"
	"github.com/oktopriima/gank/app/usecase/API/orders/create"
	"github.com/oktopriima/gank/app/usecase/API/orders/list"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	create create.Inport
	list   list.Inport
}

func NewOrderHandler(create create.Inport, list list.Inport) OrderHandler {
	return OrderHandler{create: create, list: list}
}

func (r *OrderHandler) Create(c echo.Context) error {
	req := new(create.InportRequest)
	res := new(create.InportResponse)

	if err := c.Bind(&req); err != nil {
		res.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res = r.create.Interactor(req, c.Request().Context())

	if !res.Success {
		return c.JSON(http.StatusUnprocessableEntity, res)
	}
	return c.JSON(http.StatusOK, res)
}

func (r *OrderHandler) List(c echo.Context) error {
	req := new(list.InportRequest)
	res := new(list.InportResponse)

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		res.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	req.Page = page

	size, err := strconv.Atoi(c.QueryParam("size"))
	if err != nil {
		res.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	req.Size = size

	res = r.list.Interactor(req, c.Request().Context())
	if !res.Success {
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	return c.JSON(http.StatusOK, res)
}
