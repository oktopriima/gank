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
	"github.com/oktopriima/gank/app/handler/API"
	"net/http"
)

func NewAPIRoutes(
	e *echo.Echo,
	user API.UserHandler,
) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:       middleware.DefaultSkipper,
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Authorization"},
	}))

	route := e.Group("api")

	route.GET("/ping", func(context echo.Context) error {
		return context.JSON(http.StatusOK, struct {
			Message string
		}{Message: "all is good"})
	})

	{
		userRoute := route.Group("/user")
		userRoute.POST("/register", user.Register)
	}

}
