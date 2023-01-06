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
	cm "github.com/oktopriima/gank/app/middleware"
	"net/http"
)

func NewAPIRoutes(
	e *echo.Echo,
	authenticateMiddleware *cm.AuthenticateMiddleware,
	user API.UserHandler,
	cart API.CartHandler,
	product API.ProductHandler,
	order API.OrderHandler,
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
		userRoute.POST("/login", user.Login)
	}

	{
		productRoute := route.Group("/product")
		productRoute.Use(cm.Authenticated(authenticateMiddleware))
		productRoute.POST("", product.Create)
		productRoute.GET("", product.List)
	}

	{
		cartRoute := route.Group("/cart")
		cartRoute.Use(cm.Authenticated(authenticateMiddleware))
		cartRoute.POST("", cart.Create)
		cartRoute.GET("", cart.List)
		cartRoute.PUT("/:id", cart.Update)
		cartRoute.DELETE("/:id", cart.Delete)
	}

	{
		orderRoute := route.Group("/order")
		orderRoute.Use(cm.Authenticated(authenticateMiddleware))
		orderRoute.POST("", order.Create)
		orderRoute.GET("", order.List)
	}
}
