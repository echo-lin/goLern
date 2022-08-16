package common

import (
	"github.com/dgrijalva/jwt-go"
	"golern/model"
	"time"
)

var jwtKey = []byte("lweruionxcvi*&^6323kjsdfdf")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// ReleaseToken
// @Description 发放token
// @Date 2022-07-27 05:31:19
// @param user
// @return string
// @return error
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "goLern.com",
			IssuedAt:  time.Now().Unix(),
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
