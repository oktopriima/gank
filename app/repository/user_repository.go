/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 02/01/23, 14:51
 * Copyright (c) 2023
 */

package repository

import (
	"context"
	"github.com/oktopriima/gank/app/models"
	"github.com/oktopriima/gank/app/usecase/API/users"
	"gorm.io/gorm"
)

type userRepository struct{}

func NewUserRepository() users.UserOutport {
	return &userRepository{}
}

func (u *userRepository) Create(db *gorm.DB, user *models.User, ctx context.Context) error {
	err := db.WithContext(ctx).Create(&user).Error
	return err
}
