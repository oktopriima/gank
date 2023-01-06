/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46
 * Copyright (c) 2022
 */

package registry

import (
	"github.com/oktopriima/gank/app/handler/API"
	"go.uber.org/dig"
)

func NewHandlerRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(API.NewUserHandler); err != nil {
		panic(err)
	}

	if err = container.Provide(API.NewCartHandler); err != nil {
		panic(err)
	}

	if err = container.Provide(API.NewProductHandler); err != nil {
		panic(err)
	}

	if err = container.Provide(API.NewOrderHandler); err != nil {
		panic(err)
	}

	return container
}
