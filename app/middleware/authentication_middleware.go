/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46 : 29/12/22, 10:37
 * Copyright (c) 2022
 */

package middleware

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"github.com/oktopriima/gank/app/repository"
	"github.com/oktopriima/gank/app/response"
	"net/http"
)

type AuthenticationMiddlewareConfig struct {
	tokenRepository repository.AccessTokenRepository
	httpResponse    response.HttpResponse
}

func NewAuthenticationMiddlewareConfig(
	tokenRepository repository.AccessTokenRepository,
	httpResponse response.HttpResponse,
) *AuthenticationMiddlewareConfig {
	return &AuthenticationMiddlewareConfig{
		tokenRepository: tokenRepository,
		httpResponse:    httpResponse,
	}
}

func AuthenticationMiddleware(a *AuthenticationMiddlewareConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()

			user, pass, ok := c.Request().BasicAuth()

			if !ok {
				return a.httpResponse.ErrorResponse(
					fmt.Errorf("something went wrong"),
					c,
					http.StatusUnauthorized,
				)
			}

			if err := a.checkAccess(user, pass, ctx); err != nil {
				return a.httpResponse.ErrorResponse(err, c, http.StatusUnauthorized)
			}

			return next(c)
		}
	}
}

func (a *AuthenticationMiddlewareConfig) checkAccess(user, pass string, ctx context.Context) error {
	criteria := map[string]interface{}{
		"client_id":  user,
		"auth_token": pass,
	}

	token, err := a.tokenRepository.FindBy(criteria, ctx)
	if err != nil {
		return err
	}

	if token == nil {
		return fmt.Errorf("access token not found")
	}

	return nil
}
