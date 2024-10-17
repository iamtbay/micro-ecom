package main

import (
	"errors"
	"net/http"
	"os"
	"regexp"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
var nameRegex = regexp.MustCompile(`^[a-zA-Z]+([ '-] [a-zA-Z]+)*$`)

// IS MAIL VALID?
func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// IS NAME VALID?
func isValidName(name string) bool {
	return nameRegex.MatchString(name)

}

// IS PASSWORD VALID?
func isValidPassword(password string) bool {
	var hasMinLen, hasUpper, hasLower, hasNumber, hasSpecial bool
	const minLen = 8

	if len(password) >= minLen {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

// HASH PASSWORD
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// CHECK PASSWORD
func isPasswordCorrect(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// CREATE JWT
func createJWT(id uuid.UUID, email string) (string, error) {
	claims := jwtClaims{
		UserID: id,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "tyr-Shopping",
			Subject:   "access",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//SIGN JWT
	ss, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return ss, nil
}

// PARSE JWT
func parseJWT(tokenString string) (uuid.UUID, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return uuid.UUID{}, errors.New("invalid token, please login again")
		}
		return uuid.UUID{}, err
	} else if claims, ok := token.Claims.(*jwtClaims); ok {
		return claims.UserID, nil
	} else {
		return uuid.UUID{}, errors.New("something went wrong while verifying user")
	}

}

// SET COOKIE
func setCookie(c *gin.Context, tokenName, token string) {
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie(tokenName, token, 3600, "/", "localhost", true, true)
}

// GET COOKIE
func getCookie(c *gin.Context) (string, error) {
	token, err := c.Cookie("accessToken")
	if err != nil {
		return "", err
	}
	return token, nil
}
