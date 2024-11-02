package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HELPER

// CHECK OBJECT ID IS VALID
func turnIdToObjID(id string) (primitive.ObjectID, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return objID, err
	}
	return objID, nil
}

// CHECK CREDENTIAL IS VALID
func (x *Services) checkCredentials(productInfo *NewProduct) error {
	if productInfo.Name == "" || len(productInfo.Name) < 2 {
		return errors.New("name can't be empty or less than 2 char")
	}
	if productInfo.Brand == "" || len(productInfo.Brand) < 2 {
		return errors.New("brand can't be empty or less than 2 char")
	}
	if productInfo.Content == "" || len(productInfo.Content) < 20 {
		return errors.New("content can't be empty or less than 20 char")
	}

	return nil
}

//CHECK COOKIE IS VALID

func isCookieValid(c *gin.Context) (uuid.UUID, error) {
	tokenString, err := c.Cookie("accessToken")
	if err != nil {
		errJSON(http.StatusUnauthorized, err, "Access error", c)
		c.Abort()
		return uuid.UUID{}, err

	}

	return parseJWT(tokenString)
}

func parseJWT(tokenString string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return uuid.UUID{}, errors.New("invalid token, please login again")
		}
		return uuid.UUID{}, err
	}

	if claims, ok := token.Claims.(*jwtClaims); ok {
		return claims.UserID, nil
	}
	return uuid.UUID{}, errors.New("something went wrong while verifying the user")

}
