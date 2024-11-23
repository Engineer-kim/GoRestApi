package models

import (
	"Go-RestApi/db"
	"log"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events = []Event{}

func (receiverEvent *Event) Save() error {
	query :=
		`INSERT INTO events(name, description, location, dateTime, user_id) 
		 VALUES (?,?,?,?,?)`
	stmt, err := db.DB.Prepare(query) //Prepare 메서드를 사용하여 쿼리를 준비,  SQL 인젝션 공격을 방지하고 성능을 향상시키기 위해 쿼리를 미리 컴파일
	if err != nil {
		return err
	}
	defer stmt.Close()
	// stmt.Exec 메서드를 사용하여 준비된 쿼리를 실행
	result, err := stmt.Exec(receiverEvent.Name, receiverEvent.Description, receiverEvent.Location, receiverEvent.DateTime, receiverEvent.UserID)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}
	//마지막으로 삽입된 레코드의 ID를 가져옴. 이 ID는 보통 자동 증가하는 기본 키(Auto Increment)
	id, err := result.LastInsertId()
	receiverEvent.ID = id
	//events = append(events, receiverEvent)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	//실행은 Exec , Fetch(데이터 읽어옴) Query 사용
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // 리소스가 항상 시행(메모리 누수 방지)

	var events []Event
	for rows.Next() { // 하나씩 읽는데 false가돠면 중단
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) UpdateEvent() error {
	query := `UPDATE events SET 
                  name = ?, 
                  description = ?, 
                  location = ?,
                  dateTime = ?
              WHERE id = ?`
	//Prepare 문이란 쿼리 실행 계획이고( 준비된 쿼리에 대한 최적의 실행 계획을 미리 생성해 놓기에 성능 향상에 도움이 됨)
	//SQL 인젝션 방지가능 ==> ?를 사용하여 값을 바인딩 하기때문(악의 적인 쿼리 전체를 하나의 문자열로 인식)
	//SELECT * FROM users WHERE username = 'admin OR 1=1';  ==> 'admin OR 1=1' 전체를 값이 아닌 스트링으로 인식 하기때문
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		println(err.Error())
		return err
	}
	defer stmt.Close()

	//쿼리의 ? 부분에 값 바인딩 하는 부분
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	return err
}

func (event Event) Register(userID int64) error {
	query := "INSERT INTO registrations(event_id , user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID, userID)
	return err

}

func (event Event) CancelRegistration(userID int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID, userID)
	return err
}
