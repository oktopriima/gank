/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 04/01/23, 14:49
 * Copyright (c) 2023
 */

package models

import "time"

type Cart struct {
	ID           int       `json:"id"`
	ProductID    int       `json:"product_id"`
	UserID       int       `json:"user_id"`
	Quantity     float64   `json:"quantity"`
	IsClaimed    bool      `json:"is_claimed"`
	OrderClaimed string    `json:"order_claimed"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Products     *Product  `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

type CartParams struct {
	ID        int
	ProductID int
	UserID    int
	IsClaimed bool
	Page      int
	Size      int
}

func (c Cart) TableName() string {
	return "cart"
}
