/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46
 * Copyright (c) 2022
 */

package service

import (
	"context"
	"github.com/oktopriima/gank/app/models"
	"github.com/oktopriima/gank/app/repository"
	"github.com/oktopriima/gank/config"
	"gorm.io/gorm"
)

type accessTokenServices struct {
	db *gorm.DB
}

func NewAccessTokenServices(dbConfig config.DatabaseInstance) repository.AccessTokenRepository {
	return &accessTokenServices{db: dbConfig.MySql()}
}

func (a *accessTokenServices) FindBy(criteria map[string]interface{}, ctx context.Context) (*models.AccessToken, error) {
	result := new(models.AccessToken)
	if err := a.db.WithContext(ctx).Where(criteria).First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
