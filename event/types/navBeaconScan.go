package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("NavBeaconScan", func() event.Event {
		return new(NavBeaconScan)
	})
}

type NavBeaconScan struct {
	NumBodies int64     `json:"NumBodies"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e NavBeaconScan) GetEvent() string {
	return e.Event
}

func (e NavBeaconScan) GetTimestamp() time.Time {
	return e.Timestamp
}
