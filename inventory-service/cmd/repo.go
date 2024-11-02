package main

import (
	"context"
	"time"
)

type Repository struct{}

func initRepository() *Repository {
	return &Repository{}
}

// !
// NEW PRODUCT STOCK
func (x *Repository) newProductStock(product Product) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//sql
	query := `INSERT INTO 
			inventory(product_id,properties,available_stock) 
			VALUES($1,$2,$3)`
	_, err := conn.Exec(ctx, query, product.ProductID, product.Properties, product.AvailableStock)
	if err != nil {
		return err
	}
	return nil

}

// !
// GET PRODUCT STOCK BY ID
func (x *Repository) getStock(productID string) (*Product, error) {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var product Product
	//sql
	query := `SELECT * FROM inventory WHERE product_id=$1`
	err := conn.QueryRow(ctx, query, productID).Scan(
		&product.ID,
		&product.ProductID,
		&product.Properties,
		&product.AvailableStock,
		&product.ReservedStock,
	)
	if err != nil {
		return &Product{}, err
	}
	return &product, nil
}

// !
// PRODUCT RE STOCK
func (x *Repository) productReStock(product Product) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//sql
	query := `UPDATE inventory
			  SET available_stock=available_stock+$1
			  WHERE product_id=$2`
	_, err := conn.Exec(ctx, query, product.AvailableStock, product.ProductID)
	if err != nil {
		return err
	}
	return nil
}

// !
// CANCEL RESERVATION
func (x *Repository) cancelReservation(product ProductData) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//
	query := `UPDATE inventory SET reserved_stock=reserved_stock - $1, available_stock=available_stock + $2 WHERE product_id=$3`
	_, err := conn.Exec(ctx, query, product.Quantity, product.Quantity, product.ProductID)
	if err != nil {
		return err
	}
	return nil
}

// !
// UPDATE STOCK VIA RESERVED
func (x *Repository) updateStockViaReserved(product ProductData) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//
	query := `UPDATE inventory SET reserved_stock=reserved_stock + $1, available_stock=available_stock - $2 WHERE product_id=$3`
	_, err := conn.Exec(ctx, query, product.Quantity, product.Quantity, product.ProductID)
	if err != nil {
		return err
	}
	return nil
}

// !
// UPDATE STOCK VIA RESERVED
func (x *Repository) updateStockViaSold(product ProductData) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//
	query := `UPDATE inventory SET reserved_stock=reserved_stock-$1 WHERE product_id=$2`
	_, err := conn.Exec(ctx, query, product.Quantity, product.ProductID)
	if err != nil {
		return err
	}
	return nil
}

//!
// CHECK IS AVAILABLE TO SELL

func (x *Repository) checkStock(productID string) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//
	var stock int
	query := `SELECT available_stock FROM inventory WHERE product_id=$1`
	err := conn.QueryRow(ctx, query, productID).Scan(&stock)
	if err != nil {
		return 0, err
	}
	return stock, nil
}
