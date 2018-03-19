package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("EngineerApply", func() event.Event {
		return new(EngineerApply)
	})
}

type EngineerApply struct {
	Blueprint string    `json:"Blueprint"`
	Engineer  string    `json:"Engineer"`
	Level     int64     `json:"Level"`
	Event     string    `json:"event"`
	Override  string    `json:"Override"`
	Timestamp time.Time `json:"timestamp"`
}

func (e EngineerApply) GetEvent() string {
	return e.Event
}

func (e EngineerApply) GetTimestamp() time.Time {
	return e.Timestamp
}
