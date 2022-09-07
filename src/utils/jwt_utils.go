package utils

import (
	"time"

	"github.com/aivot-digital/preppy-backend/src/cfg"
	"github.com/aivot-digital/preppy-backend/src/models"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtClaims struct {
	UserId    primitive.ObjectID
	IsAdmin   bool
	IsRefresh bool
	jwt.StandardClaims
}

func CreateJwt(user *models.User, refresh bool) (string, error) {
	expires := 2 * time.Hour
	if refresh {
		expires = 7 * 24 * time.Hour
	}
	claims := JwtClaims{
		UserId:    user.Id,
		IsAdmin:   user.IsAdmin,
		IsRefresh: refresh,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(time.Duration(expires)).Unix(),
			Id:        "",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "",
			NotBefore: 0,
			Subject:   "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(cfg.SECURE_KEY)
}

func ValidateJwt(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JwtClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return cfg.SECURE_KEY, nil
		})

	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return false, err
	}

	return claims.ExpiresAt > time.Now().Unix(), nil
}
