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
	"fmt"
	"github.com/oktopriima/gank/app/models"
	"github.com/oktopriima/gank/app/usecase/API/users"
	"gorm.io/gorm"
)

type userRepository struct{}

func (u *userRepository) FindBy(db *gorm.DB, user models.UserParams, ctx context.Context) (*models.User, error) {
	result := new(models.User)
	res := db.WithContext(ctx)

	if user.Email != "" {
		res = res.Where("email = ?", user.Email)
	}

	if user.FirstName != "" {
		res = res.Where("first_name LIKE ?", fmt.Sprintf("%%%s%%", user.FirstName))
	}

	if user.LastName != "" {
		res = res.Where("last_name LIKE ?", fmt.Sprintf("%%%s%%", user.LastName))
	}

	res = res.First(result)

	if err := res.Error; err != nil {
		return nil, err
	}

	return result, nil
}

func NewUserRepository() users.UserOutport {
	return &userRepository{}
}

func (u *userRepository) Create(db *gorm.DB, user *models.User, ctx context.Context) error {
	err := db.WithContext(ctx).Create(&user).Error
	return err
}
