/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 04/01/23, 16:35
 * Copyright (c) 2023
 */

package API

import (
	"github.com/labstack/echo"
	"github.com/oktopriima/gank/app/usecase/API/carts/create"
	delete2 "github.com/oktopriima/gank/app/usecase/API/carts/delete"
	"github.com/oktopriima/gank/app/usecase/API/carts/list"
	"github.com/oktopriima/gank/app/usecase/API/carts/update"
	"net/http"
	"strconv"
)

type CartHandler struct {
	create create.Inport
	list   list.Inport
	update update.Inport
	delete delete2.Inport
}

func NewCartHandler(create create.Inport,
	list list.Inport,
	update update.Inport,
	delete delete2.Inport) CartHandler {
	return CartHandler{
		create: create,
		list:   list,
		update: update,
		delete: delete,
	}
}

func (b *CartHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	req := create.InportRequest{}
	res := new(create.InportResponse)

	err := c.Bind(&req)

	if err != nil {
		res.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res = b.create.Interactor(&req, ctx)
	if !res.Success {
		return c.JSON(http.StatusUnprocessableEntity, res)
	}
	return c.JSON(http.StatusOK, res)
}

func (b *CartHandler) List(c echo.Context) error {
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

	res = b.list.Interactor(req, ctx)
	if !res.Success {
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	return c.JSON(http.StatusOK, res)
}

func (b *CartHandler) Update(c echo.Context) error {
	req := new(update.InportRequest)
	res := new(update.InportResponse)

	if err := c.Bind(req); err != nil {
		res.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, err)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, err)
	}
	req.ID = id

	res = b.update.Interactor(req, c.Request().Context())
	if !res.Success {
		res.ErrorMessage = err.Error()
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, res)
}

func (b *CartHandler) Delete(c echo.Context) error {
	req := new(delete2.InportRequest)
	res := new(delete2.InportResponse)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	req.ID = id

	res = b.delete.Interactor(req, c.Request().Context())
	if !res.Success {
		return c.JSON(http.StatusUnprocessableEntity, res)
	}
	return c.JSON(http.StatusOK, res)
}
