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
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(receiverEvent.Name, receiverEvent.Description, receiverEvent.Location, receiverEvent.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	receiverEvent.ID = int(id)
	//events = append(events, receiverEvent)
	return err
}

func GetAllEvents() []Event {
	return events
}
