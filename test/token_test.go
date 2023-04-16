package test

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"testing"
	"time"
)

type MyCustomClaims struct {
	User string `json:"username"`
	jwt.RegisteredClaims
}

type T struct{}

// Generate  生成token
func (T) Generate(username string) (tokenStr string, err error) {
	claims := MyCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(20000 * time.Second)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",     //颁发者
			Subject:   "somebody", //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString([]byte("lgp"))
	if err != nil {
		return "", err
	}
	return tokenStr, err
}

func (T) Parse(tokenStr string) (userId string, err error) {
	token, _ := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("lgp"), nil
	})
	claims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {
		return claims.User, nil
	}
	err = errors.New("用户未认证")
	return "", err
}

func TestToken(t *testing.T) {
	var Token = T{}
	tokenStr, err := Token.Generate("test")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(tokenStr)

	userId, err := Token.Parse(tokenStr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(userId)
}
