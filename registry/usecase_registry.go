/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46
 * Copyright (c) 2022
 */

package registry

import (
	"github.com/oktopriima/gank/app/usecase/API/carts/create"
	deleteCart "github.com/oktopriima/gank/app/usecase/API/carts/delete"
	"github.com/oktopriima/gank/app/usecase/API/carts/list"
	"github.com/oktopriima/gank/app/usecase/API/carts/update"
	createOrder "github.com/oktopriima/gank/app/usecase/API/orders/create"
	listOrder "github.com/oktopriima/gank/app/usecase/API/orders/list"
	createProduct "github.com/oktopriima/gank/app/usecase/API/products/create"
	listProduct "github.com/oktopriima/gank/app/usecase/API/products/list"
	"github.com/oktopriima/gank/app/usecase/API/users/login"
	"github.com/oktopriima/gank/app/usecase/API/users/register"
	"go.uber.org/dig"
)

func NewUsecaseRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(register.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(login.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(create.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(createProduct.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(list.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(listProduct.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(update.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(deleteCart.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(createOrder.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(listOrder.NewUsecase); err != nil {
		panic(err)
	}

	return container
}
