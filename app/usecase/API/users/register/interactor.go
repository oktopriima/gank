/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 02/01/23, 14:55
 * Copyright (c) 2023
 */

package register

import (
	"context"
	"github.com/oktopriima/gank/app/models"
	"github.com/oktopriima/gank/app/modules"
	"github.com/oktopriima/gank/app/usecase/API/users"
	"gorm.io/gorm"
)

type createInteractor struct {
	outport users.UserOutport
	bcrypt  modules.BcryptPassword
	db      *gorm.DB
}

func (c *createInteractor) Interactor(ctx context.Context, request InportRequest) *InportResponse {
	res := new(InportResponse)

	// generate password from request
	pass, err := c.bcrypt.Encrypt(request.Password)
	if err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	// assign request to models
	user := new(models.User)
	user.Email = request.Email
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Password = pass

	if err = c.outport.Create(c.db, user, ctx); err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	res.Success = true
	res.Data = user
	return res
}

func NewUsecase(outport users.UserOutport, bcrypt modules.BcryptPassword, db *gorm.DB) Inport {
	return &createInteractor{
		outport: outport,
		bcrypt:  bcrypt,
		db:      db,
	}
}
