/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46
 * Copyright (c) 2022
 */

package response

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type httpResponse struct {
}

func (h *httpResponse) ErrorResponse(err error, ctx echo.Context, code ...int) error {
	statusCode := http.StatusUnprocessableEntity
	if len(code) > 0 {
		statusCode = code[0]
	}

	resp := struct {
		Message          string `json:"message"`
		DeveloperMessage string `json:"developer_message"`
	}{Message: http.StatusText(statusCode), DeveloperMessage: err.Error()}

	return ctx.JSON(statusCode, resp)
}

func (h *httpResponse) SuccessResponse(data interface{}, ctx echo.Context, code ...int) error {
	statusCode := http.StatusOK
	if len(code) > 0 {
		statusCode = code[0]
	}

	resp := struct {
		Message          string      `json:"message"`
		DeveloperMessage string      `json:"developer_message"`
		Data             interface{} `json:"data"`
	}{
		Message:          http.StatusText(statusCode),
		DeveloperMessage: fmt.Sprintf("OK"),
		Data:             data,
	}
	return ctx.JSON(statusCode, resp)
}

type HttpResponse interface {
	ErrorResponse(err error, ctx echo.Context, code ...int) error
	SuccessResponse(data interface{}, ctx echo.Context, code ...int) error
}

func NewHttpResponse() (response HttpResponse, err error) {
	return &httpResponse{}, nil
}
