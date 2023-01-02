/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 02/01/23, 14:48
 * Copyright (c) 2023
 */

package models

import "time"

func (u User) TableName() string {
	return "users"
}

type User struct {
	ID         int       `json:"id"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Password   string    `json:"-"`
	IsDeleted  bool      `json:"-"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
