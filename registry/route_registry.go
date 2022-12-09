/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 17/10/22, 14:31
 * Copyright (c) 2022
 */

package registry

import (
	"github.com/oktopriima/ninja/app/routes/api"
	"github.com/oktopriima/ninja/app/routes/web"
	"go.uber.org/dig"
)

func NewRouteRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Invoke(api.NewAPIRoutes); err != nil {
		panic(err)
	}

	if err = container.Invoke(web.NewWebRoutes); err != nil {
		panic(err)
	}

	return container
}
