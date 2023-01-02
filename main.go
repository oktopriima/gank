/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46
 * Copyright (c) 2022
 */

package main

import (
	"github.com/oktopriima/gank/config"
)

func main() {
	i := NewInjection()

	// run the webserver
	if err := i.Invoke(func(instance *config.ServerInstance) {
		instance.RunWithGracefullyShutdown()
	}); err != nil {
		panic(err)
	}
}
