package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (receiverEvent Event) Save() {
	events = append(events, receiverEvent)
}

func GetAllEvents() []Event {
	return events
}
