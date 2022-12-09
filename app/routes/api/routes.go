/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 14/10/22, 14:35
 * Copyright (c) 2022
 */

package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/oktopriima/ninja/app/handler/API/users"

	"net/http"
)

func NewAPIRoutes(
	e *echo.Echo,
	userHandler users.BaseHandler,
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

	{
		userRoutes := route.Group("/user")
		{
			userInternal := userRoutes.Group("/internal")
			userInternal.GET("", userHandler.SelfHandler)
		}
	}
}
