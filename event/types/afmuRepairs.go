package types

import (
	"time"

	"github.com/IgaguriMK/ed-journal/event"
)

func init() {
	event.RegisterEvent("AfmuRepairs", func() Event {
		return new(AfmuRepairs)
	})
}

type AfmuRepairs struct {
	FullyRepaired   bool      `json:"FullyRepaired"`
	Health          float64   `json:"Health"`
	Module          string    `json:"Module"`
	ModuleLocalised string    `json:"Module_Localised"`
	Event           string    `json:"event"`
	Timestamp       time.Time `json:"timestamp"`
}

func (e AfmuRepairs) GetEvent() string {
	return e.Event
}

func (e AfmuRepairs) GetTimestamp() time.Time {
	return e.Timestamp
}
