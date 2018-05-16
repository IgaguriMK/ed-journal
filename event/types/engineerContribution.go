package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("EngineerContribution", func() event.Event {
		return new(EngineerContribution)
	})
}

type EngineerContribution struct {
	Commodity          string    `json:"Commodity"`
	CommodityLocalised string    `json:"Commodity_Localised"`
	Engineer           string    `json:"Engineer"`
	EngineerID         int64     `json:"EngineerID"`
	Quantity           int64     `json:"Quantity"`
	Material           string    `json:"Material"`
	MaterialLocalised  string    `json:"Material_Localised"`
	TotalQuantity      int64     `json:"TotalQuantity"`
	Type               string    `json:"Type"`
	Event              string    `json:"event"`
	Timestamp          time.Time `json:"timestamp"`
}

func (e EngineerContribution) GetEvent() string {
	return e.Event
}

func (e EngineerContribution) GetTimestamp() time.Time {
	return e.Timestamp
}
