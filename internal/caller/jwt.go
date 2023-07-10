package caller

import (
	"github.com/golang-jwt/jwt"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"time"
)

var jwtKey = []byte("a_secret_create")

type MyClaims struct {
	UserID int64
	jwt.StandardClaims
}

func ReleaseToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &MyClaims{
		UserID: user.UID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "lyh",
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

func ParseToken(tokenString string) (*jwt.Token, *MyClaims, error) {
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
