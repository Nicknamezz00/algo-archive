package app

import (
	"algo-archive/internal/conf"
	"algo-archive/internal/model"
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	UID      int64  `json:"uid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func GenerateToken(User *model.User) (string, error) {
	now := time.Now()
	expire := now.Add(conf.JWTSetting.Expire)
	claims := Claims{
		UID:      User.ID,
		Username: User.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			Issuer:    conf.JWTSetting.Issuer + ":" + User.Salt,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func GetJWTSecret() []byte {
	return []byte(conf.JWTSetting.Secret)
}
