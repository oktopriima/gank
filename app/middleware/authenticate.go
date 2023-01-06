/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 04/01/23, 15:34
 * Copyright (c) 2023
 */

package middleware

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/oktopriima/gank/app/helpers"
	"github.com/oktopriima/gank/app/modules"
	"github.com/oktopriima/gank/app/response"
	"github.com/oktopriima/gank/app/usecase/API/users"
	"gorm.io/gorm"
	"net/http"
)

const AuthToken = "AUTH_TOKEN"
const ExtractedValue = "EXTRACTED_TOKEN"

type AuthenticateMiddleware struct {
	user     users.UserOutport
	jwtToken *modules.JwtModules
	db       *gorm.DB
}

func NewAuthenticateMiddleware(outport users.UserOutport, jwtToken *modules.JwtModules, db *gorm.DB) (*AuthenticateMiddleware, error) {
	return &AuthenticateMiddleware{
		user:     outport,
		jwtToken: jwtToken,
		db:       db,
	}, nil
}

func Authenticated(middleware *AuthenticateMiddleware) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			oldCtx := c.Request().Context()

			res := new(response.BaseResponse)
			token, err := helpers.HeaderExtractor("Authorization", c.Request())
			if err != nil {
				res.ErrorMessage = err.Error()
				return c.JSON(http.StatusForbidden, res)
			}

			extracted, err := middleware.jwtToken.ExtractToken(token)
			if err != nil {
				res.ErrorMessage = err.Error()
				return c.JSON(http.StatusForbidden, res)
			}

			// set auth info into the context
			ctx := context.WithValue(oldCtx, AuthToken, token)
			ctx = context.WithValue(ctx, ExtractedValue, extracted)

			request := c.Request().WithContext(ctx)
			c.SetRequest(request)

			return next(c)
		}
	}
}

func GetAuthenticatedUser(ctx context.Context) (int, error) {
	value := ctx.Value(ExtractedValue)

	extracted := modules.JwtExtracted{}

	valBytes, err := json.Marshal(value)
	if err != nil {
		return 0, err
	}
	if err := json.Unmarshal(valBytes, &extracted); err != nil {
		return 0, err
	}

	return extracted.UserID, nil
}
