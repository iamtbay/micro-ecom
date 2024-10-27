package main

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func isCookieValid(c *gin.Context) (uuid.UUID, error) {
	tokenString, err := c.Cookie("accessToken")
	if err != nil {
		return uuid.UUID{}, err
	}
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return uuid.UUID{}, errors.New("invalid token, please login again")
		}
		return uuid.UUID{}, nil
	} else if claims, ok := token.Claims.(*jwtClaims); ok {
		return claims.UserID, nil
	} else {
		return uuid.UUID{}, errors.New("something went wrong while verifying user")
	}

	//

}

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
