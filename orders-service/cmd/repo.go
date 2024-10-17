package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Repository struct{}

func initRepo() *Repository {
	return &Repository{}
}

// !
// GET ORDER
func (x *Repository) getOrder(orderUUID uuid.UUID) (Order, error) {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//query
	query := `SELECT * FROM orders WHERE id=$1 and is_active=TRUE`

	var orderInfo Order
	var productsJSON []byte
	//
	err := conn.QueryRow(ctx, query, orderUUID).Scan(
		&orderInfo.OrderID,
		&orderInfo.CustomerID,
		&productsJSON,
		&orderInfo.AddressID,
		&orderInfo.TotalPrice,
		&orderInfo.OrderDate,
		&orderInfo.IsActive,
	)
	if err != nil {
		return Order{}, err
	}

	err = json.Unmarshal(productsJSON, &orderInfo.Products)
	if err != nil {
		return Order{}, err
	}
	return orderInfo, nil
}

// !
// GET ORDERS BY USER ID
func (x *Repository) getAllOrdersByUserID(userID uuid.UUID) ([]*Order, error) {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//query
	query := `SELECT * FROM orders WHERE user_id=$1 AND is_active=TRUE`
	rows, err := conn.Query(ctx, query, userID)
	fmt.Println(rows)
	if err != nil {
		return []*Order{}, err
	}
	//scan to file
	var orders []*Order
	for rows.Next() {
		var order Order
		var products []byte
		err := rows.Scan(
			&order.OrderID,
			&order.CustomerID,
			&products,
			&order.AddressID,
			&order.TotalPrice,
			&order.OrderDate,
			&order.IsActive,
		)
		if err != nil {
			return []*Order{}, err
		}
		err = json.Unmarshal(products, &order.Products)
		if err != nil {
			return []*Order{}, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

// !
// NEW ORDER
func (x *Repository) newOrder(order Order, productsJson []byte) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//
	query := `INSERT INTO orders(user_id,products,shipping_adress_id,total_price)
			VALUES($1,$2,$3,$4)
			`
	_, err := conn.Exec(ctx, query, order.CustomerID, productsJson, order.AddressID, order.TotalPrice)
	if err != nil {
		fmt.Println("Err? here", err)
		return err
	}

	return nil
}

// !
// DELETE ORDER
func (x *Repository) deleteOrder(orderID, userID uuid.UUID) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//query
	query := `UPDATE orders SET is_active = FALSE
			WHERE id=$1 AND user_id=$2 `

	_, err := conn.Exec(ctx, query, orderID, userID)
	if err != nil {
		return err
	}
	return nil
}
