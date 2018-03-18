package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("EscapeInterdiction", func() event.Event {
		return new(EscapeInterdiction)
	})
}

type EscapeInterdiction struct {
	Interdictor          string    `json:"Interdictor"`
	InterdictorLocalised string    `json:"Interdictor_Localised"`
	IsPlayer             bool      `json:"IsPlayer"`
	Event                string    `json:"event"`
	Timestamp            time.Time `json:"timestamp"`
}

func (e EscapeInterdiction) GetEvent() string {
	return e.Event
}

func (e EscapeInterdiction) GetTimestamp() time.Time {
	return e.Timestamp
}
