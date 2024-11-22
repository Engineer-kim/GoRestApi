package utils

import (
	"errors"
	"fmt"
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

func VerifyToken(token string) (int64, error) {
	// JWT 토큰을 파싱하고 서명을 검증합니다.
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("Could Not Parse Token")
	}

	tokenIsValid := parseToken.Valid
	if !tokenIsValid {
		return 0, errors.New("Invalid Token")
	}

	// 아래의 코드는 JWT 토큰에서의 필요한 정보 추출 방법
	claims, ok := parseToken.Claims.(jwt.MapClaims)
	println(claims)
	if !ok {
		return 0, errors.New("Invalid Token")
	}
	//email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId, nil
}
