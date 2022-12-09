/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 14/10/22, 14:06
 * Copyright (c) 2022
 */

package main

import (
	"github.com/oktopriima/ninja/config"
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
