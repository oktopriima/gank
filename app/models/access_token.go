/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46 : 29/12/22, 10:47
 * Copyright (c) 2022
 */

package models

import (
	"fmt"
	"time"
)

func (a AccessToken) TableName() string {
	return fmt.Sprintf("access_tokens")
}

type AccessToken struct {
	ID          int64     `json:"id"`
	DisabledBy  int64     `json:"disabled_by"`
	GeneratedBy int64     `json:"generated_by"`
	ResetBy     int64     `json:"reset_by"`
	AuthToken   string    `json:"auth_token"`
	ClientID    string    `json:"client_id"`
	ClientName  string    `json:"client_name"`
	IsActive    bool      `json:"is_active"`
	DisabledAt  time.Time `json:"disabled_at"`
	ResetAt     time.Time `json:"reset_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Scopes      string    `json:"scopes"`
}
