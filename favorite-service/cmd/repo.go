package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct{}

func initRepository() *Repository {
	return &Repository{}
}

func (x *Repository) newFavorite(userID uuid.UUID, favoriteProduct FavoriteProductBSON) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	filter := bson.M{
		"user_id": primitive.Binary{Subtype: 4, Data: userID[:]},
	}
	update := bson.M{
		"$addToSet": bson.M{
			"favorites": bson.M{
				"product_id": favoriteProduct.ProductID,
			},
		},
	}
	opts := options.Update().SetUpsert(true)
	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		panic(err)
	}
	if result.UpsertedCount > 0 {
		fmt.Printf("A new document was created with ID: %v\n", result.UpsertedID)
	} else {
		fmt.Printf("Modified %v document(s)\n", result.ModifiedCount)
	}
	//collection.InsertOne(ctx, favoriteProduct)

	return nil
}

func (x *Repository) getAllFavorites(userID uuid.UUID) (AllFavorites, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	filter := bson.M{
		"user_id": primitive.Binary{Subtype: 4, Data: userID[:]},
	}
	var favProducts AllFavorites
	err := collection.FindOne(ctx, filter).Decode(&favProducts)
	if err != nil {
		return AllFavorites{}, err
	}

	return favProducts, nil
}

func (x *Repository) removeFavorite(userID uuid.UUID, productID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	filter := bson.M{
		"user_id": primitive.Binary{Subtype: 4, Data: userID[:]},
	}
	update := bson.M{
		"$pull": bson.M{
			"favorites": bson.M{
				"product_id": productID,
			},
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
