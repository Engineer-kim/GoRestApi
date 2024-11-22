package models

import (
	"Go-RestApi/db"
	"Go-RestApi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (userReceiver User) Save() error {
	query := `INSERT INTO s (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic("Fail to insert userData")
		return err
	}

	defer stmt.Close()

	cryptoPassword, err := utils.HashPassword(userReceiver.Password)

	if err != nil {
		println("Fail to hash password")
		return err
	}

	result, err := stmt.Exec(userReceiver.Email, cryptoPassword)

	if err != nil {
		return err
	}
	//새로운 행 삽입 이후
	//자동 증가(auto-increment) 컬럼에 부여된 마지막으로 생성된 ID 값을 반환하는 메서드
	userId, err := result.LastInsertId()

	userReceiver.ID = userId
	return err
}
