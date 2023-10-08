package pkg

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Token struct {
}

const Key = "下雨的星星"

func NewToken() *Token {
	return &Token{}
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (t *Token) SigningToken(username string) (string, error) {
	nowTime := time.Now()
	claims := MyClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "test",
			ExpiresAt: jwt.NewNumericDate(nowTime.Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(nowTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(Key))
	return tokenString, err
}

func (t *Token) ValidatingToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
	})
	if err != nil {
		return nil, errors.New("token验证失败:" + err.Error())
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token验证失败")
	}
}
