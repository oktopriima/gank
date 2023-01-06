/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 06/01/23, 11:19
 * Copyright (c) 2023
 */

package API

import (
	"github.com/labstack/echo"
	"github.com/oktopriima/gank/app/usecase/API/products/create"
	"github.com/oktopriima/gank/app/usecase/API/products/list"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	create create.Inport
	list   list.Inport
}

func NewProductHandler(create create.Inport, list list.Inport) ProductHandler {
	return ProductHandler{
		create: create,
		list:   list,
	}
}

func (p *ProductHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	req := create.InportRequest{}
	res := new(create.InportResponse)

	err := c.Bind(&req)

	if err != nil {
		res.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res = p.create.Interactor(&req, ctx)
	if !res.Success {
		return c.JSON(http.StatusUnprocessableEntity, res)
	}
	return c.JSON(http.StatusOK, res)
}

func (p *ProductHandler) List(c echo.Context) error {
	ctx := c.Request().Context()

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
	req.Search = c.QueryParam("search")

	res = p.list.Interactor(req, ctx)
	if !res.Success {
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	return c.JSON(http.StatusOK, res)
}
