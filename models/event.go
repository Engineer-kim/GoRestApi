package models

import (
	"Go-RestApi/db"
	"time"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (receiverEvent Event) Save() error {
	query :=
		`INSERT INTO events(name, description, location, user_id) 
		 VALUES (?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query) //Prepare 메서드를 사용하여 쿼리를 준비,  SQL 인젝션 공격을 방지하고 성능을 향상시키기 위해 쿼리를 미리 컴파일
	if err != nil {
		return err
	}
	defer stmt.Close()
	// stmt.Exec 메서드를 사용하여 준비된 쿼리를 실행
	result, err := stmt.Exec(receiverEvent.Name, receiverEvent.Description, receiverEvent.Location, receiverEvent.UserID)
	if err != nil {
		return err
	}
	//마지막으로 삽입된 레코드의 ID를 가져옴. 이 ID는 보통 자동 증가하는 기본 키(Auto Increment)
	id, err := result.LastInsertId()
	receiverEvent.ID = int(id)
	//events = append(events, receiverEvent)
	return err
}

func GetAllEvents() []Event {
	return events
}
