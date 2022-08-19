package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"time"
	error2 "wtv-online/utils/error"
)

var jwtSecret = []byte("bcm-9c0b56e0-0f8a-43ab-856b-9bc5f37e2e7e")
var TokenMap = map[string]string{} // 后端临时存储token位置

type Claims struct {
	Username string `json:"username"`
	UserId   int    `json:"user_id"`
	jwt.StandardClaims
}

func GeneratePassword(s, salt string) string {
	b := []byte(s)
	h := md5.New()
	h.Write(b)
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateToken(username string, userId int) (string, error) {
	nowTime := time.Now()                     //当前时间
	expireTime := nowTime.Add(24 * time.Hour) //有效时间

	claims := Claims{
		Username: username,
		UserId:   userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "YiGeChengZi",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	m := Md5(username)
	TokenMap[m] = token
	return m, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, _ := jwt.ParseWithClaims(TokenMap[token], &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, error2.NotFound("登录已经失效")
}
