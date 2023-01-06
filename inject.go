/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46
 * Copyright (c) 2022
 */

package main

import (
	"github.com/oktopriima/gank/app/middleware"
	"github.com/oktopriima/gank/registry"
	"go.uber.org/dig"
)

func NewInjection() *dig.Container {
	c := dig.New()

	c = registry.NewConfigRegistry(c)
	c = registry.NewModulesRegistry(c)
	c = registry.NewHandlerRegistry(c)
	c = registry.NewRepositoryRegistry(c)
	c = registry.NewUsecaseRegistry(c)

	// custom class
	if err := c.Provide(middleware.NewAuthenticateMiddleware); err != nil {
		panic(err)
	}

	c = registry.NewRouteRegistry(c)

	return c
}
