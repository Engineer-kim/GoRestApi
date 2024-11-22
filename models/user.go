package models

import (
	"Go-RestApi/db"
	"Go-RestApi/utils"
	"errors"
	"log"
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

func (userReceiver User) ValidateCredential() error {
	query := `SELECT email, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, userReceiver.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		log.Println("err::", err.Error())
		return errors.New("Invalid Credential")
	}

	passwordIsValid := utils.CheckPasswordHash(userReceiver.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("Invalid Credential")
	}
	return nil
}
