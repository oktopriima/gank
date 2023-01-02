/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46
 * Copyright (c) 2022
 */

package registry

import (
	"github.com/labstack/echo"
	"github.com/oktopriima/gank/app/response"
	"github.com/oktopriima/gank/config"
	"go.uber.org/dig"
)

func NewConfigRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(config.NewEnvironment); err != nil {
		panic(err)
	}

	if err = container.Provide(func() *echo.Echo {
		return echo.New()
	}); err != nil {
		panic(err)
	}

	if err = container.Provide(config.NewServerInstance); err != nil {
		panic(err)
	}

	if err = container.Provide(config.NewDatabaseInstance); err != nil {
		panic(err)
	}

	if err = container.Provide(response.NewHttpResponse); err != nil {
		panic(err)
	}

	return container
}
