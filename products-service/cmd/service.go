package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Services struct {
}

func initServices() *Services {
	return &Services{}
}

//!SERVICES

var repo = initRepository()

// !
// GET ALL PRODUCTS
func (x *Services) getProducts(page int64) ([]*GetProduct, *PageInfo, error) {
	products, pageInfos, err := repo.getProducts(page)
	if err != nil {
		fmt.Println(page)
		return []*GetProduct{}, &PageInfo{}, err
	}
	return products, pageInfos, nil
}

// !
// GET SINGLE PRODUCT
func (x *Services) getSingleProduct(id string) (*GetProduct, error) {
	objID, err := turnIdToObjID(id)
	if err != nil {
		return &GetProduct{}, err
	}
	product, err := repo.getSingleProduct(objID)
	if err != nil {
		return &GetProduct{}, err
	}
	return product, nil
}

// !
// ADD PRODUCT
func (x *Services) addProduct(newProduct *NewProduct) error {

	err := x.checkCredentials(newProduct)
	if err != nil {
		return err
	}
	productID, err := repo.addProduct(newProduct)

	if err != nil {
		return err
	}
	err = publishNewProduct(ProductInventoryType{
		ProductID: productID.Hex(),
		Quantity:  newProduct.Stock,
	})
	if err != nil {
		return err
	}

	err = publishNewProductIndex(productID, newProduct)
	if err != nil {
		return err
	}

	return nil
}

// !
// EDIT PRODUCT
func (x *Services) editProduct(id string, newProduct *NewProduct) (*NewProduct, error) {
	objID, err := turnIdToObjID(id)
	if err != nil {
		return &NewProduct{}, err
	}
	err = x.checkCredentials(newProduct)
	if err != nil {
		return &NewProduct{}, err
	}
	err = repo.editProduct(objID, newProduct)
	if err != nil {
		return &NewProduct{}, err
	}

	//publish changes
	product := GetProduct{
		ID:    objID,
		Name:  newProduct.Name,
		Price: newProduct.Price,
	}
	err = publishPrice(product)
	if err != nil {
		log.Println("Error while publishig price", err)
	}

	return newProduct, nil
}

// !
// DELETE PRODUCT
func (x *Services) deleteProduct(id string, userID uuid.UUID) error {

	objID, err := turnIdToObjID(id)
	if err != nil {
		return err
	}
	//also user id needed to verify user has authorized or not
	err = repo.deleteProduct(objID, userID)
	if err != nil {
		return err
	}

	return nil
}
