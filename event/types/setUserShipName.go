package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("SetUserShipName", func() event.Event {
		return new(SetUserShipName)
	})
}

type SetUserShipName struct {
	Ship         string    `json:"Ship"`
	ShipID       int64     `json:"ShipID"`
	UserShipID   string    `json:"UserShipId"`
	UserShipName string    `json:"UserShipName"`
	Event        string    `json:"event"`
	Timestamp    time.Time `json:"timestamp"`
}

func (e SetUserShipName) GetEvent() string {
	return e.Event
}

func (e SetUserShipName) GetTimestamp() time.Time {
	return e.Timestamp
}
