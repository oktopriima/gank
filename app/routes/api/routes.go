/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46
 * Copyright (c) 2022
 */

package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	customMiddleware "github.com/oktopriima/gank/app/middleware"

	"net/http"
)

func NewAPIRoutes(
	e *echo.Echo,
	authConfig *customMiddleware.AuthenticationMiddlewareConfig,
) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	route := e.Group("api")

	route.GET("/ping", func(context echo.Context) error {
		return context.JSON(http.StatusOK, struct {
			Message string
		}{Message: "hello from motherffffff!!!!"})
	})
}
