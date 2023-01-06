/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 04/01/23, 14:51
 * Copyright (c) 2023
 */

package models

import "time"

type OrderDetail struct {
	ID         int       `json:"id"`
	OrderID    int       `json:"order_id"`
	ProductID  int       `json:"product_id"`
	Quantity   float64   `json:"quantity"`
	Price      float64   `json:"price"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Product    *Product  `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

func (d OrderDetail) TableName() string {
	return "order_detail"
}
