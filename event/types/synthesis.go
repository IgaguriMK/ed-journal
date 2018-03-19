package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Synthesis", func() event.Event {
		return new(Synthesis)
	})
}

type Synthesis struct {
	Materials []struct {
		Count          int64  `json:"Count"`
		Name           string `json:"Name"`
		Name_Localised string `json:"Name_Localised"`
	} `json:"Materials"`
	Name      string    `json:"Name"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Synthesis) GetEvent() string {
	return e.Event
}

func (e Synthesis) GetTimestamp() time.Time {
	return e.Timestamp
}
