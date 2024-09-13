package main

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Services struct {
}

func initServices() *Services {
	return &Services{}
}

//!SERVICES

var repo = initRepository()

// GET ALL PRODUCTS
func (x *Services) getProducts() ([]*GetProduct, error) {
	products, err := repo.getProducts()
	if err != nil {
		return []*GetProduct{}, err
	}
	return products, nil
}

// GET SINGLE PRODUCT
func (x *Services) getSingleProduct(id string) (*GetProduct, error) {
	objID := turnIdToObjID(id)
	product, err := repo.getSingleProduct(objID)
	if err != nil {
		return &GetProduct{}, err
	}
	return product, nil
}

func (x *Services) addProduct(newProduct *NewProduct) error {

	err := repo.addProduct(newProduct) // user id because of added by?
	if err != nil {
		return err
	}

	return nil
}

func (x *Services) editProduct(id string, newProduct *NewProduct) (*NewProduct, error) {
	objID := turnIdToObjID(id)
	err := repo.editProduct(objID, newProduct)
	if err != nil {
		return &NewProduct{}, err
	}

	return newProduct, nil
}

func (x *Services) deleteProduct(id string) error {

	objID := turnIdToObjID(id)
	//also user id needed to verify user has authorized or not
	err := repo.deleteProduct(objID)
	if err != nil {
		return err
	}

	return nil
}

// HELPER
func turnIdToObjID(id string) (objID primitive.ObjectID) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("something bad while turning ID to object ID")
	}
	return
}
