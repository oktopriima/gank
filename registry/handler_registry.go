/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 14/10/22, 14:29
 * Copyright (c) 2022
 */

package registry

import (
	"github.com/oktopriima/ninja/app/handler/API/users"
	"go.uber.org/dig"
)

func NewHandlerRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(users.NewBaseHandler); err != nil {
		panic(err)
	}

	return container
}
