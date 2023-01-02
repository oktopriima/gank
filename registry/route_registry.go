/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46 : 17/10/22, 14:31
 * Copyright (c) 2022
 */

package registry

import (
	"github.com/oktopriima/gank/app/routes/api"
	"go.uber.org/dig"
)

func NewRouteRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Invoke(api.NewAPIRoutes); err != nil {
		panic(err)
	}

	return container
}
