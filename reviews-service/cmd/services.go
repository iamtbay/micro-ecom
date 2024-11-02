package main

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Services struct{}

func initServices() *Services {
	return &Services{}
}

var repo = initRepository()

// !
// GET REVIEWS BY PRODUCT ID
func (x *Services) getProductReviewsByProductID(productIDStr string) ([]*GetReview, error) {
	productID, err := primitive.ObjectIDFromHex(productIDStr)
	if err != nil {
		return []*GetReview{}, err
	}
	reviews, err := repo.getProductReviewsByProductID(productID)
	if err != nil {
		return  []*GetReview{}, err
	}
	return reviews, nil
}

// !
// GET REVIEW BY ID
func (x *Services) getReviewByID(reviewIDStr string) (*GetReview, error) {
	reviewID, err := primitive.ObjectIDFromHex(reviewIDStr)
	if err != nil {
		return &GetReview{}, err
	}
	//
	review, err := repo.getReviewByID(reviewID)
	if err != nil {
		return review, err
	}
	return review, nil
}

// !
// NEW REVIEW
func (x *Services) newReview(review NewReview, productIDStr string) error {
	productID, err := primitive.ObjectIDFromHex(productIDStr)
	if err != nil {
		return err
	}
	review.ProductID = productID
	review.Date = time.Now()

	reviewBSON := reviewsToBSON(review)
	err = repo.newReview(reviewBSON)
	if err != nil {
		return err
	}
	return nil
}

// !
// EDIT REVIEW BY ID
func (x *Services) editReviewByReviewID(reviewIDStr string, review NewReview) error {
	reviewID, err := primitive.ObjectIDFromHex(reviewIDStr)
	if err != nil {
		return err
	}
	err = repo.editReviewByReviewID(reviewID, review)
	if err != nil {
		return err
	}
	return nil

}

// !
// DELETE REVIEW BY ID
func (x *Services) deleteReviewByReviewID(reviewIDStr string, userID uuid.UUID) error {
	reviewID, err := primitive.ObjectIDFromHex(reviewIDStr)
	if err != nil {
		return err
	}
	err = repo.deleteReviewByReviewID(reviewID, userID)
	if err != nil {
		return err
	}
	return nil

}
