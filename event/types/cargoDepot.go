package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("CargoDepot", func() event.Event {
		return new(CargoDepot)
	})
}

type CargoDepot struct {
	CargoType           string    `json:"CargoType"`
	Count               int64     `json:"Count"`
	EndMarketID         int64     `json:"EndMarketID"`
	ItemsCollected      int64     `json:"ItemsCollected"`
	ItemsDelivered      int64     `json:"ItemsDelivered"`
	MissionID           int64     `json:"MissionID"`
	Progress            float64   `json:"Progress"`
	StartMarketID       int64     `json:"StartMarketID"`
	TotalItemsToDeliver int64     `json:"TotalItemsToDeliver"`
	UpdateType          string    `json:"UpdateType"`
	Event               string    `json:"event"`
	Timestamp           time.Time `json:"timestamp"`
}

func (e CargoDepot) GetEvent() string {
	return e.Event
}

func (e CargoDepot) GetTimestamp() time.Time {
	return e.Timestamp
}
