package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"img/server/global"
	"time"
)

type MyCustomClaims struct {
	User string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(username string) (tokenStr string, err error) {
	tokenConf := global.Config.Token
	claims := MyCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenConf.ExpiresTime), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",     //颁发者
			Subject:   "somebody", //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
	tokenStr, err = token.SignedString([]byte(tokenConf.SigningKey))
	if err != nil {
		return "", err
	}
	return tokenStr, err
}

func ParseToken(tokenStr string) (userId string, err error) {
	tokenConf := global.Config.Token
	token, _ := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenConf.SigningKey), nil
	})
	claims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {
		return claims.User, nil
	}
	err = errors.New("用户未认证")
	return "", err
}
