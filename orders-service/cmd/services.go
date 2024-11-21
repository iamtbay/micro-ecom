package main

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type Services struct{}

func initServices() *Services { return &Services{} }

var repo = initRepo()

// !
// GET SINGLE ORDER BY ID
func (x *Services) getSingleOrder(orderID string) (*Order, error) {
	err := uuid.Validate(orderID)
	if err != nil {
		return &Order{}, errors.New("unvalid order uuid")
	}
	//convert it to uuid
	orderUUID, err := uuid.Parse(orderID)
	if err != nil {
		return &Order{}, err
	}
	//get response from database
	orderData, err := repo.getOrder(orderUUID)
	if err != nil {
		return &Order{}, err
	}

	return &orderData, nil
}

// !
// GET ALL ORDERS BY USER ID
func (x *Services) getAllOrdersByUserID(userID string) ([]*Order, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return []*Order{}, err
	}
	orders, err := repo.getAllOrdersByUserID(userUUID)
	if err != nil {
		return []*Order{}, err
	}
	return orders, nil
}

// !
// NEW ORDER
func (x *Services) newOrder(order Order) (uuid.UUID, error) {
	productsJson, err := json.Marshal(order.Products)
	if err != nil {
		return uuid.UUID{}, err
	}

	//repo
	orderID, err := repo.newOrder(order, productsJson)
	if err != nil {
		return uuid.UUID{}, err
	}

	return orderID, nil
}

// !
// get single order
func (x *Services) deleteOrder(orderID string, userID uuid.UUID) error {
	orderUUID, err := uuid.Parse(orderID)
	if err != nil {
		return err
	}

	err = repo.deleteOrder(orderUUID, userID)
	if err != nil {
		return err
	}
	return nil
}
