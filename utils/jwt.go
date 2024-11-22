package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "superSecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":          email,
		"userId":         userId,
		"expirationTime": time.Now().Add(time.Hour * 2).Unix(), //현재 시간에서 2시간 후의 시간을 Unix 시간으로 변환
	})
	//jwt 서명 전체를 완성해서 리턴함
	return token.SignedString([]byte(secretKey)) //JWT 암호화 카를 평문으로 리턴하면 토큰의 내용이 노출되어 위변조될 위험
}
