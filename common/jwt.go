package common

import (
	"go_gin_second/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//定义jwt加密的密钥
var jwtKey = []byte("gin_second_key")

/*
Claims 定义Token的Claims
*/
type Claims struct {
	UserID uint
	jwt.StandardClaims
}

/*
ReleaseTOken 发放token
*/
func ReleaseTOken(user model.User) (string, error) {

	expitationTime := time.Now().Add(7 * 24 * time.Hour) //token有效期 7 * 24h

	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expitationTime.Unix(), //token过期时间
			IssuedAt:  time.Now().Unix(),     //token 发放时间
			Issuer:    "gin_second",          //token的发布者
			Subject:   "user token",          //主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey) //使用自己的jetKey作为密钥加密
	if err != nil {                                //token生成失败，返回错误
		return "", err
	}
	//返回生成的token字符串
	return tokenString, err
}

//ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
