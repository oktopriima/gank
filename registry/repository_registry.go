/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46
 * Copyright (c) 2022
 */

package registry

import (
	"github.com/oktopriima/gank/app/repository"
	"go.uber.org/dig"
)

func NewRepositoryRegistry(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(repository.NewUserRepository); err != nil {
		panic(err)
	}

	if err = container.Provide(repository.NewCartRepository); err != nil {
		panic(err)
	}

	if err = container.Provide(repository.NewProductRepository); err != nil {
		panic(err)
	}

	if err = container.Provide(repository.NewOrderRepository); err != nil {
		panic(err)
	}
	return container
}
