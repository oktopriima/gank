/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 02/01/23, 14:55
 * Copyright (c) 2023
 */

package users

import (
	"context"
	"github.com/oktopriima/gank/app/models"
	"gorm.io/gorm"
)

type UserOutport interface {
	Create(db *gorm.DB, user *models.User, ctx context.Context) error
	FindBy(db *gorm.DB, user models.UserParams, ctx context.Context) (*models.User, error)
}
