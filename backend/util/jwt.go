package util

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("little_lion")

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	UserId   string `form:"userId" json:"userId" binding:"required"`
}

type UserClaims struct {
	User
	jwt.StandardClaims
}

func GenerateJWT(user User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := UserClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "hello",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := tokenClaims.SignedString(jwtSecret)

	return jwt, err
}

func ParseJWT(JWT string) (*User, error) {
	tokenClaims, err := jwt.ParseWithClaims(JWT, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*UserClaims); ok && tokenClaims.Valid {
			return &claims.User, nil
		}
	}

	return nil, err
}

func IsJwtValid(JWT string) (bool, error) {
	tokenClaims, err := jwt.ParseWithClaims(JWT, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	// 判断是否解析成功
	if err != nil || tokenClaims == nil {
		return false, err
	}

	claims, ok := tokenClaims.Claims.(*UserClaims)

	// 是否断言成功
	if !ok {
		return false, errors.New("claims trans failed")
	}
	// 是否合法
	if !tokenClaims.Valid {
		return false, errors.New("tokenClaims is invalid")
	}
	// 是否超时
	if time.Now().Unix() > claims.ExpiresAt {
		return false, errors.New("jwt expire time limit")
	}
	return true, nil
}

func SplitJwt(token string) string {
	slic := strings.Split(token, " ")
	if len(slic) != 2 {
		return ""
	} else if slic[0] != "Bearer" {
		return ""
	} else {
		return slic[1]
	}
}
