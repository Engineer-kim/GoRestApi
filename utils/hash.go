package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	//입력 받은 암호를 바이트로 바꿈
	//bcrypt.MinCost ==> 해싱에 사용되는 비용(높을수록 보안성 올라가고, 처리 속도 느려짐)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil //암호가 유효하면 True 리턴 아니라면 False 리턴
}
