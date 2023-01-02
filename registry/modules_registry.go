/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 02/01/23, 15:24
 * Copyright (c) 2023
 */

package registry

import (
	"github.com/oktopriima/gank/app/modules"
	"go.uber.org/dig"
)

func NewModulesRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(modules.NewBcryptPassword); err != nil {
		panic(err)
	}

	return container
}
