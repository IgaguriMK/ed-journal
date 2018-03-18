package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("USSDrop", func() event.Event {
		return new(USSDrop)
	})
}

type USSDrop struct {
	USSThreat        int64     `json:"USSThreat"`
	USSType          string    `json:"USSType"`
	USSTypeLocalised string    `json:"USSType_Localised"`
	Event            string    `json:"event"`
	Timestamp        time.Time `json:"timestamp"`
}

func (e USSDrop) GetEvent() string {
	return e.Event
}

func (e USSDrop) GetTimestamp() time.Time {
	return e.Timestamp
}
