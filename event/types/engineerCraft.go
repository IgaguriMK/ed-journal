package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("EngineerCraft", func() event.Event {
		return new(EngineerCraft)
	})
}

type EngineerCraft struct {
	Blueprint   string `json:"Blueprint"`
	Engineer    string `json:"Engineer"`
	Ingredients []struct {
		Count int64  `json:"Count"`
		Name  string `json:"Name"`
	} `json:"Ingredients"`
	Level     int64     `json:"Level"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e EngineerCraft) GetEvent() string {
	return e.Event
}

func (e EngineerCraft) GetTimestamp() time.Time {
	return e.Timestamp
}