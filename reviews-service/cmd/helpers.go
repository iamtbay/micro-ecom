package main

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func reviewsToBSON(review NewReview) NewReviewBSON {
	return NewReviewBSON{
		ProductID: review.ProductID,
		UserID:    primitive.Binary{Subtype: 4, Data: review.UserID[:]},
		Name:      review.Name,
		Surname:   review.Surname,
		Point:     review.Point,
		Comment:   review.Comment,
		Date:      review.Date,
		IsDeleted: review.IsDeleted,
	}
}

func reviewsToNormal(review GetReviewBSON) (GetReview, error) {
	userID, err := uuid.FromBytes(review.UserID.Data)
	if err != nil {
		return GetReview{}, err
	}
	return GetReview{
		ReviewID:  review.ReviewID,
		ProductID: review.ProductID,
		UserID:    userID,
		Name:      review.Name,
		Surname:   review.Surname,
		Point:     review.Point,
		Comment:   review.Comment,
		Date:      review.Date,
		IsDeleted: review.IsDeleted,
	}, nil
}
