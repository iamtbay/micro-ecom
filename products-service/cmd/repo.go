package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
}

func initRepository() *Repository {
	return &Repository{}
}

// GET PRODUCTS
func (x *Repository) getProducts() ([]*GetProduct, error) {
	var products []*GetProduct

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return []*GetProduct{}, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product *GetProduct
		if err := cursor.Decode(&product); err != nil {
			fmt.Println(cursor)
			fmt.Println(err)
			return products, errors.New("something went wrong while cursor the item")
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	//filters?
	//pagination?
	return products, nil

}

// GET SINGLE PRODUCT
func (x *Repository) getSingleProduct(id primitive.ObjectID) (*GetProduct, error) {
	var product *GetProduct
	//context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//filter
	filter := bson.M{"_id": id}

	err := collection.FindOne(ctx, filter).Decode(&product)
	fmt.Println("product is", product)
	if err != nil {
		return product, err
	}
	if err == mongo.ErrNoDocuments {
		return product, errors.New("Invalid id")
	}
	return product, nil

}

// ADD PRODUCT
func (x *Repository) addProduct(newProduct *NewProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//query
	result, err := collection.InsertOne(ctx, newProduct)
	if err != nil {
		return err
	}
	fmt.Println("inserted with", result.InsertedID)
	return nil
}

// EDIT PRODUCT
func (x *Repository) editProduct(id primitive.ObjectID, newProductInfo *NewProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name":    newProductInfo.Name,
			"brand":   newProductInfo.Brand,
			"content": newProductInfo.Content,
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	fmt.Println("updated with", result.UpsertedID)
	return nil

}

// DELETE PRODUCT
func (x *Repository) deleteProduct(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name":     "",
			"brand":    "",
			"content":  "",
			"added_by": "",
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	fmt.Println("deleted id", result.UpsertedID)
	return nil
}
