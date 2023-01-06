/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 04/01/23, 14:54
 * Copyright (c) 2023
 */

package login

import (
	"context"
	"fmt"
	"github.com/oktopriima/gank/app/models"
	"github.com/oktopriima/gank/app/modules"
	"github.com/oktopriima/gank/app/usecase/API/users"
	"gorm.io/gorm"
)

type loginInteractor struct {
	outport users.UserOutport
	bcrypt  modules.BcryptPassword
	token   *modules.JwtModules
	db      *gorm.DB
}

func (l *loginInteractor) Interactor(ctx context.Context, request InportRequest) *InportResponse {
	res := new(InportResponse)

	user, err := l.outport.FindBy(l.db, models.UserParams{
		Email: request.Email,
	}, ctx)
	if err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	if !l.bcrypt.Decrypt(user.Password, request.Password) {
		res.ErrorMessage = fmt.Sprintf("password not match")
		return res
	}

	// generate JWT token
	jwtToken, err := l.token.GenerateToken(modules.JwtParams{
		UserID: user.ID,
		Email:  user.Email,
	})
	if err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	res.Success = true
	res.Data = jwtToken
	return res
}

func NewUsecase(outport users.UserOutport, password modules.BcryptPassword, token *modules.JwtModules, db *gorm.DB) Inport {
	return &loginInteractor{
		outport: outport,
		bcrypt:  password,
		token:   token,
		db:      db,
	}
}
