package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("EngineerProgress", func() event.Event {
		return new(EngineerProgress)
	})
}

type EngineerProgress struct {
	Engineer   string    `json:"Engineer"`
	EngineerID int64     `json:"EngineerID"`
	Progress   string    `json:"Progress"`
	Rank       int64     `json:"Rank"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e EngineerProgress) GetEvent() string {
	return e.Event
}

func (e EngineerProgress) GetTimestamp() time.Time {
	return e.Timestamp
}
