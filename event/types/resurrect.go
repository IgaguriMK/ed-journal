package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Resurrect", func() event.Event {
		return new(Resurrect)
	})
}

type Resurrect struct {
	Bankrupt  bool      `json:"Bankrupt"`
	Cost      int64     `json:"Cost"`
	Option    string    `json:"Option"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Resurrect) GetEvent() string {
	return e.Event
}

func (e Resurrect) GetTimestamp() time.Time {
	return e.Timestamp
}
