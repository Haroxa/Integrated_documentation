package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"github.com/Haroxa/Integrated_documentation/common"
)

// token 的claim

type JwtClaim struct {
	UserId int
	jwt.RegisteredClaims
}

// this key is the most dangerous!!!! MUST BE DIFFICULT TO GUESS
// jwt 加密密匙
var myKey = []byte("fahkdslfhakldsjfklasdk321084710jfd")

func CreateToken(UserId int) (string, error) {
	claim := JwtClaim{
		// 自定义字段
		UserId: UserId,
		// 标准字段
		RegisteredClaims: jwt.RegisteredClaims{
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().In(common.ChinaTime).Add(168 * time.Hour)),
			// 发放时间
			IssuedAt: jwt.NewNumericDate(time.Now().In(common.ChinaTime)),
		}}
	// 使用jwt密匙生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func VerifyToken(token string) (int, error) {
	// 解析 token
	tempToken, err := jwt.ParseWithClaims(token, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return -1, err
	}
	// 获取 claim 部分，并检验
	claims, ok := tempToken.Claims.(*JwtClaim)
	if !ok {
		return -1, errors.New("claims error")
	}
	if err = tempToken.Claims.Valid(); err != nil {
		return -1, err
	}
	return claims.UserId, nil
}
