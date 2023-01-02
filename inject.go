/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46
 * Copyright (c) 2022
 */

package main

import (
	"github.com/oktopriima/gank/registry"
	"go.uber.org/dig"
)

func NewInjection() *dig.Container {
	c := dig.New()

	c = registry.NewConfigRegistry(c)
	c = registry.NewHandlerRegistry(c)
	c = registry.NewServiceRegistry(c)
	c = registry.NewUsecaseRegistry(c)
	c = registry.NewAddOnsRegistry(c)
	c = registry.NewRouteRegistry(c)

	return c
}
