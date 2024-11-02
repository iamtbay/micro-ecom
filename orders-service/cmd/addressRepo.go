package main

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// !
// GET ADDRESSES
func (x *Repository) getAddresses(userID uuid.UUID) ([]*GetAddresses, error) {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM addresses WHERE user_id=$1 AND is_deleted=FALSE`
	rows, err := conn.Query(ctx, query, userID)
	if err != nil {
		return []*GetAddresses{}, err
	}

	var addresses []*GetAddresses

	for rows.Next() {
		address, err := x.scanAddressToVariable(rows)
		if err != nil {
			return addresses, err
		}
		addresses = append(addresses, &address)
	}
	return addresses, nil
}

// !
//GET SINGLE ADRESS BY ID
func (x *Repository) getSingleAddressByID(userID, addressID uuid.UUID) (*GetAddresses, error) {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM addresses WHERE id=$1 AND user_id=$2 AND is_deleted=FALSE`
	var address GetAddresses
	rows := conn.QueryRow(ctx, query, addressID, userID)

	address, err := x.scanAddressToVariable(rows)
	if err != nil {
		return &address, err
	}
	return &address, nil
}

// !
func (x *Repository) addNewAddress(addressInfo NewAddress) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//
	query := `INSERT INTO addresses(user_id, address_name, street, city, state, postal_code, country)
			VALUES($1,$2,$3,$4,$5,$6,$7) `
	_, err := conn.Exec(ctx, query, addressInfo.UserID, addressInfo.AddressName, addressInfo.Street, addressInfo.City, addressInfo.State, addressInfo.PostalCode, addressInfo.Country)
	if err != nil {
		return err
	}
	return nil
}

// !
// EDIT ADDRESS BY ID
func (x *Repository) editAddressByID(addressInfo NewAddress, addressID uuid.UUID) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//sql
	query := `UPDATE addresses
			SET
			address_name = $1,
			street = $2,
			city = $3,
			state = $4,
			postal_code = $5,
			country =$6
			WHERE id=$7 AND user_id=$8 AND is_deleted=FALSE
			`
	_, err := conn.Exec(ctx, query,
		addressInfo.AddressName,
		addressInfo.Street,
		addressInfo.City,
		addressInfo.State,
		addressInfo.PostalCode,
		addressInfo.Country,
		addressID,
		addressInfo.UserID,
	)
	if err != nil {
		return err
	}
	return nil
}

// !
// DELETE ADDRESS BY ID
func (x *Repository) deleteAddressByID(userID, addressID uuid.UUID) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//sql
	query := `UPDATE addresses
			SET
			is_deleted=TRUE
			WHERE id=$1 AND user_id=$2
			`
	_, err := conn.Exec(ctx, query, addressID, userID)
	if err != nil {
		return err
	}
	return nil
}

//HELPER

func (x *Repository) scanAddressToVariable(rows pgx.Row) (GetAddresses, error) {
	var address GetAddresses
	err := rows.Scan(
		&address.ID,
		&address.UserID,
		&address.AddressName,
		&address.Street,
		&address.City,
		&address.State,
		&address.PostalCode,
		&address.Country,
		&address.IsDeleted,
	)
	if err != nil {
		return GetAddresses{}, err
	}
	return address, nil
}
