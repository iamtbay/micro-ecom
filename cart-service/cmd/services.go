package main

import (
	"errors"
	"log"
	"strconv"

	"github.com/google/uuid"
)

type Services struct{}

func initServices() *Services {
	return &Services{}
}

var repo = initRepo()

// !
// GET CART
func (x *Services) getCart(userID uuid.UUID) (CartOrder, error) {

	//get items on user's cart
	cart, err := repo.getCart(userID)
	if err != nil {
		return CartOrder{}, err
	}

	return cart, nil
}

// !
// GET CART
func (x *Services) checkOut(userID uuid.UUID, addressID uuid.UUID) error {

	//check prices are there any price changed or not?

	//get items on user's cart
	cart, err := repo.getCart(userID)
	if err != nil {
		return err
	}

	if len(cart.Products) == 0 {
		return errors.New("no product in your cart")
	}

	cart.AddressID = addressID
	//publish ordered message
	err = publishMessage(ch, cart)
	if err != nil {
		log.Fatalf("Failed to publish message %v", err)
	}
	//clear the cart
	err = repo.deleteUserCart(userID)
	if err != nil {
		return err
	}

	return nil
}

// !
// ADD TO CART
func (x *Services) addToCart(userID uuid.UUID, product CartItem) error {
	err := repo.addToCart(userID, product)
	if err != nil {
		return err
	}
	return nil
}

func (x *Services) updateQuantityOfProduct(userID uuid.UUID, productID string, quantity string, isExact bool) (string, error) {
	productQuantity, err := strconv.Atoi(quantity)
	if err != nil {
		return "", err
	}

	msg, err := repo.updateQuantityOfProduct(userID, productID, productQuantity, isExact)
	if err != nil {
		return "", err
	}
	return msg, nil
}

// !
// DELETE ITEM ON CART
func (x *Services) deleteProductOnCart(userID uuid.UUID, productID string) error {
	err := repo.deleteProductOnCart(userID, productID)
	if err != nil {
		return err
	}

	return nil
}

//!
//UPDATE PRODUCT ON CART
func (x *Services) updateProduct(product UpdateProductType) error{
	err := repo.updateProduct(product)
	if err!=nil {
		return err
	}
	return nil
}