package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository struct {
}

func initRepository() *Repository {
	return &Repository{}
}

// !
// GET REVIEWS BY PRODUCT ID
func (x *Repository) getProductReviewsByProductID(productID primitive.ObjectID) ([]*GetReview, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"product_id": productID})
	if err != nil {
		return []*GetReview{}, err
	}
	var reviews []*GetReview
	for cursor.Next(ctx) {
		var review GetReviewBSON
		if err := cursor.Decode(&review); err != nil {
			return reviews, err
		}
		userID, err := uuid.FromBytes(review.UserID.Data)
		if err != nil {
			return reviews, err
		}
		//check here
		reviews = append(reviews, &GetReview{
			ReviewID:  review.ReviewID,
			ProductID: review.ProductID,
			UserID:    userID, //
			Name:      review.Name,
			Surname:   review.Surname,
			Point:     review.Point,
			Comment:   review.Comment,
			Date:      review.Date,
			IsDeleted: review.IsDeleted,
		})
	}
	return reviews, nil
}

// !
// GET REVIEW BY ID
func (x *Repository) getReviewByID(reviewID primitive.ObjectID) (*GetReview, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	filter := bson.M{"_id": reviewID}
	var review GetReviewBSON
	err := collection.FindOne(ctx, filter).Decode(&review)
	if err != nil {
		return &GetReview{}, err
	}
	userID, err := uuid.FromBytes(review.UserID.Data)
	if err != nil {
		return &GetReview{}, err
	}

	return &GetReview{
		ReviewID:  review.ReviewID,
		ProductID: review.ProductID,
		UserID:    userID, //
		Name:      review.Name,
		Surname:   review.Surname,
		Point:     review.Point,
		Comment:   review.Comment,
		Date:      review.Date,
		IsDeleted: review.IsDeleted,
	}, nil
}

// !
// NEW REVIEW
func (x *Repository) newReview(newReview NewReviewBSON) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := collection.InsertOne(ctx, newReview)
	if err != nil {
		return err
	}
	return nil
}

// !
// EDIT REVIEW BY REVIEW ID
func (x *Repository) editReviewByReviewID(reviewID primitive.ObjectID, review NewReview) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	filter := bson.M{"_id": reviewID}
	update := bson.M{
		"$set": bson.M{
			"point":   review.Point,
			"comment": review.Comment,
			"date":    time.Now(),
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

// !
// DELETE REVIEW BY REVIEW ID
func (x *Repository) deleteReviewByReviewID(reviewID primitive.ObjectID, userID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	filter := bson.M{"_id": reviewID, "user_id": primitive.Binary{Subtype: 4, Data: userID[:]}}
	update := bson.M{
		"$set": bson.M{
			"is_deleted": true,
		}}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	fmt.Println("deleted")
	return nil
}
