/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 04/01/23, 14:50
 * Copyright (c) 2023
 */

package models

import "time"

type Order struct {
	ID           int            `json:"id"`
	UserID       int            `json:"user_id"`
	OrderCode    string         `json:"order_code"`
	FinalPrice   float64        `json:"final_price"`
	Status       int            `json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	OrderDetails []*OrderDetail `json:"order_details,omitempty" gorm:"foreignKey:OrderID"`
}

type OrderParams struct {
	ID        int
	UserID    int
	OrderCode string
	Page      int
	Size      int
}

func (o Order) TableName() string {
	return "order"
}
