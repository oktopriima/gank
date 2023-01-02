/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46
 * Copyright (c) 2022
 */

package registry

import (
	"github.com/oktopriima/gank/app/middleware"
	"go.uber.org/dig"
)

func NewAddOnsRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(middleware.NewAuthenticationMiddlewareConfig); err != nil {
		panic(err)
	}

	return container
}