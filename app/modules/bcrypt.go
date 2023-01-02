/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 02/01/23, 15:17
 * Copyright (c) 2023
 */

package modules

import (
	"github.com/oktopriima/gank/config"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type bcryptPassword struct {
	Salt int
}

func (b *bcryptPassword) Encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), b.Salt)
	return string(bytes), err
}

func (b *bcryptPassword) Decrypt(hashedString, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(password))
	return err == nil
}

type BcryptPassword interface {
	Encrypt(password string) (string, error)
	Decrypt(hashedString, password string) bool
}

func NewBcryptPassword(env config.Environment) (BcryptPassword, error) {
	salt, err := strconv.Atoi(env.GetString("app.salt"))
	if err != nil {
		return nil, err
	}
	return &bcryptPassword{Salt: salt}, nil
}
