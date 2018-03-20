package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("ShipTargeted", func() event.Event {
		return new(ShipTargeted)
	})
}

type ShipTargeted struct {
	Faction            string    `json:"Faction"`
	HullHealth         float64   `json:"HullHealth"`
	LegalStatus        string    `json:"LegalStatus"`
	PilotName          string    `json:"PilotName"`
	PilotNameLocalised string    `json:"PilotName_Localised"`
	PilotRank          string    `json:"PilotRank"`
	ScanStage          int64     `json:"ScanStage"`
	ShieldHealth       float64   `json:"ShieldHealth"`
	Ship               string    `json:"Ship"`
	ShipLocalised      string    `json:"Ship_Localised"`
	TargetLocked       bool      `json:"TargetLocked"`
	Event              string    `json:"event"`
	Timestamp          time.Time `json:"timestamp"`
}

func (e ShipTargeted) GetEvent() string {
	return e.Event
}

func (e ShipTargeted) GetTimestamp() time.Time {
	return e.Timestamp
}
