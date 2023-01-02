/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46 : 29/12/22, 10:53
 * Copyright (c) 2022
 */

package repository

import (
	"context"
	"github.com/oktopriima/gank/app/models"
)

type AccessTokenRepository interface {
	FindBy(criteria map[string]interface{}, ctx context.Context) (*models.AccessToken, error)
}
