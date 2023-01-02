/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46
 * Copyright (c) 2022
 */

package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ServerInstance struct {
	Router *echo.Echo
	Env    Environment
}

func NewServerInstance(r *echo.Echo, env Environment) *ServerInstance {
	return &ServerInstance{
		Router: r,
		Env:    env,
	}
}

func (s *ServerInstance) RunWithGracefullyShutdown() {
	go func() {
		if err := s.runHttp(); err != nil && errors.Is(err, http.ErrServerClosed) {
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Router.Shutdown(ctx); err != nil {
		os.Exit(1)
	}
}

func (s *ServerInstance) runHttp() error {
	var err error
	port := fmt.Sprintf(":%s", s.Env.GetString("app.port"))

	if err = s.Router.Start(port); err != nil {
		return err
	}

	return nil
}
