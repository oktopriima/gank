/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 14/10/22, 14:11
 * Copyright (c) 2022
 */

package main

import (
	"github.com/oktopriima/ninja/registry"
	"go.uber.org/dig"
)

func NewInjection() *dig.Container {
	c := dig.New()

	c = registry.NewConfigRegistry(c)
	c = registry.NewHandlerRegistry(c)
	c = registry.NewRouteRegistry(c)

	return c
}
