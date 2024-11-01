package main

import (
	"context"
	"errors"
	"log"
	"math"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
}

func initRepository() *Repository {
	return &Repository{}
}

var itemsPerPage int64 = 10

// GET PRODUCTS
func (x *Repository) getProducts(page int64) ([]*GetProduct, *PageInfo, error) {
	var products []*GetProduct

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"$expr": bson.M{
			"$gt": bson.A{
				bson.M{"$strLenCP": "$name"},
				0,
			},
		},
	}

	//check total count is equal to page
	totalCount, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return []*GetProduct{}, &PageInfo{}, err
	}
	//check page
	page = repo.checkPage(page, totalCount)

	//
	cursor, err := collection.Find(context.Background(), filter, options.Find().SetSkip(int64((page-1)*itemsPerPage)).SetLimit(int64(itemsPerPage)))
	if err != nil {
		return []*GetProduct{}, &PageInfo{}, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product *GetProductBSON
		if err := cursor.Decode(&product); err != nil {
			return products, &PageInfo{}, err
		}
		uuidFromBinary, err := uuid.FromBytes(product.AddedBy.Data)
		if err != nil {
			return []*GetProduct{}, &PageInfo{}, err
		}
		products = append(products, &GetProduct{
			ID:      product.ID,
			Name:    product.Name,
			Brand:   product.Brand,
			Content: product.Content,
			Price:   product.Price,
			AddedBy: uuidFromBinary,
		})
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	//pageInfos
	pageInfos := &PageInfo{
		TotalPage:         int(totalCount)/10 + 1,
		CurrentPage:       int(page),
		TotalProductCount: int(totalCount),
	}
	return products, pageInfos, nil

}

// !
// GET SINGLE PRODUCT
func (x *Repository) getSingleProduct(id primitive.ObjectID) (*GetProduct, error) {
	var product *GetProductBSON
	//context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//filter
	filter := bson.M{"_id": id}

	err := collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return &GetProduct{}, err
	}
	if err == mongo.ErrNoDocuments {
		return &GetProduct{}, errors.New("invalid id")
	}

	userIDFromBinary, err := uuid.FromBytes(product.AddedBy.Data)
	if err != nil {
		return &GetProduct{}, err
	}
	return &GetProduct{
		ID:      product.ID,
		Name:    product.Name,
		Brand:   product.Brand,
		Content: product.Content,
		Price:   product.Price,
		AddedBy: userIDFromBinary,
	}, nil

}

// ADD PRODUCT
func (x *Repository) addProduct(newProduct *NewProduct) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	productInfoBson := NewProductBSON{
		Name:    newProduct.Name,
		Brand:   newProduct.Brand,
		Content: newProduct.Content,
		Price:   newProduct.Price,
		AddedBy: primitive.Binary{Subtype: 4, Data: newProduct.AddedBy[:]},
	}

	//query
	result, err := collection.InsertOne(ctx, productInfoBson)

	if err != nil {
		return primitive.NilObjectID, err
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.New("couldn't get the insreted id")
	}

	return insertedID, nil
}

// !
// EDIT PRODUCT
func (x *Repository) editProduct(id primitive.ObjectID, newProductInfo *NewProduct) error {
	//ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{
		"_id":      id,
		"added_by": primitive.Binary{Subtype: 4, Data: newProductInfo.AddedBy[:]},
	}

	update := bson.M{
		"$set": bson.M{
			"name":    newProductInfo.Name,
			"brand":   newProductInfo.Brand,
			"content": newProductInfo.Content,
			"price":   newProductInfo.Price,
		},
	}
	res, err := collection.UpdateOne(ctx, filter, update)

	if res.MatchedCount == 0 {
		return errors.New("document not found or user not authorized")
	}

	if err != nil {
		return err
	}
	return nil

}

// !
// DELETE PRODUCT
func (x *Repository) deleteProduct(id primitive.ObjectID, userID uuid.UUID) error {
	//CTX
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"_id": id, "added_by": primitive.Binary{Subtype: 4, Data: userID[:]}}
	update := bson.M{
		"$set": bson.M{
			"name":    "",
			"brand":   "",
			"content": "",
			"price":   "",
		},
	}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("document not found or user not authorized")
	}
	return nil
}

// todo HELPERS
func (x *Repository) checkPage(page, totalCount int64) int64 {
	if page < 1 {
		page = 1
	}
	if totalCount < (page * itemsPerPage) {
		floatPa := math.Ceil(float64(totalCount) / 10)
		page = int64(floatPa)
	}
	if page == 0 {
		page = 1
	}
	return page
}
