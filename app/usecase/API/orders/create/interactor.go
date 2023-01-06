/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 06/01/23, 12:44
 * Copyright (c) 2023
 */

package create

import (
	"context"
	"fmt"
	"github.com/oktopriima/gank/app/helpers"
	"github.com/oktopriima/gank/app/middleware"
	"github.com/oktopriima/gank/app/models"
	"github.com/oktopriima/gank/app/usecase/API/carts"
	"github.com/oktopriima/gank/app/usecase/API/orders"
	"github.com/oktopriima/gank/app/usecase/API/products"
	"gorm.io/gorm"
)

type createUsecase struct {
	order   orders.Outport
	cart    carts.Outport
	product products.Outport
	db      *gorm.DB
}

func (c *createUsecase) Interactor(request *InportRequest, ctx context.Context) *InportResponse {
	res := new(InportResponse)

	userID, err := middleware.GetAuthenticatedUser(ctx)
	if err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	tx := c.db.Begin()
	defer func() {
		tx.Rollback()
	}()

	var finalPrice float64
	var orderDetails []*models.OrderDetail

	orderCode := helpers.RandString(10)
	for _, cartID := range request.CartID {
		cart, err := c.cart.FindBy(tx, models.CartParams{
			ID:     cartID,
			UserID: userID,
		}, ctx)
		if err != nil {
			res.ErrorMessage = fmt.Sprintf("cart %s", err.Error())
			return res
		}

		product, err := c.product.FindBy(tx, &models.ProductParams{
			ID: cart.ProductID,
		}, ctx)
		if err != nil {
			res.ErrorMessage = fmt.Sprintf("product %s", err.Error())
			return res
		}

		finalPrice = finalPrice + (cart.Quantity * product.Price)

		// assign order details
		orderDetail := new(models.OrderDetail)
		orderDetail.Quantity = cart.Quantity
		orderDetail.ProductID = product.ID
		orderDetail.Price = product.Price
		orderDetail.TotalPrice = product.Price * cart.Quantity

		orderDetails = append(orderDetails, orderDetail)

		// update cart to claimed
		cart.IsClaimed = true
		cart.OrderClaimed = orderCode

		if err := c.cart.Update(tx, cart, ctx); err != nil {
			res.ErrorMessage = fmt.Sprintf("cart update %s", err.Error())
			return res
		}
	}

	order := new(models.Order)
	order.UserID = userID
	order.OrderCode = orderCode
	order.FinalPrice = finalPrice
	order.OrderDetails = orderDetails

	if err := c.order.Create(tx, order, ctx); err != nil {
		res.ErrorMessage = err.Error()
		return res
	}

	tx.Commit()
	res.Success = true
	return res
}

func NewUsecase(
	order orders.Outport,
	cart carts.Outport,
	product products.Outport,
	db *gorm.DB,
) Inport {
	return &createUsecase{
		order:   order,
		cart:    cart,
		product: product,
		db:      db,
	}
}
