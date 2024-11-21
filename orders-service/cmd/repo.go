package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

	//
	row := conn.QueryRow(ctx, query, orderUUID)
	return x.scanOrderToVariable(row)

}

// !
// GET ORDERS BY USER ID
func (x *Repository) getAllOrdersByUserID(userID uuid.UUID) ([]*Order, error) {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//query
	query := `SELECT * FROM orders WHERE user_id=$1 AND is_active=TRUE ORDER BY order_date DESC`
	rows, err := conn.Query(ctx, query, userID)

	if err != nil {
		return []*Order{}, err
	}
	//scan to file
	var orders []*Order
	for rows.Next() {
		order, err := x.scanOrderToVariable(rows)
		if err != nil {
			return []*Order{}, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

// !
// NEW ORDER
func (x *Repository) newOrder(order Order, productsJson []byte) (uuid.UUID, error) {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//
	var newOrderID uuid.UUID
	query := `INSERT INTO orders(user_id,products,shipping_adress_id,total_price)
			VALUES($1,$2,$3,$4) RETURNING id
			`
	err := conn.QueryRow(ctx, query, order.CustomerID, productsJson, order.AddressID, order.TotalPrice).Scan(&newOrderID)
	if err != nil {
		return uuid.UUID{}, err
	}

	return newOrderID, nil
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

// !
// SCAN ORDER TO VARIABLE
func (x *Repository) scanOrderToVariable(row pgx.Row) (Order, error) {
	var orderInfo Order
	var productsJSON []byte
	err := row.Scan(
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
