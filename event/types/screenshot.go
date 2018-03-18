package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Screenshot", func() event.Event {
		return new(Screenshot)
	})
}

type Screenshot struct {
	Body      string    `json:"Body"`
	Filename  string    `json:"Filename"`
	Height    int64     `json:"Height"`
	System    string    `json:"System"`
	Width     int64     `json:"Width"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Screenshot) GetEvent() string {
	return e.Event
}

func (e Screenshot) GetTimestamp() time.Time {
	return e.Timestamp
}
