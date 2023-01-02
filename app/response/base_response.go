/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 02/01/23, 15:00
 * Copyright (c) 2023
 */

package response

type BaseResponse struct {
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"error_message,omitempty"`
}
