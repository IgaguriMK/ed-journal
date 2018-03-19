package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("PayLegacyFines", func() event.Event {
		return new(PayLegacyFines)
	})
}

type PayLegacyFines struct {
	Amount           int64     `json:"Amount"`
	BrokerPercentage float64   `json:"BrokerPercentage"`
	Event            string    `json:"event"`
	Timestamp        time.Time `json:"timestamp"`
}

func (e PayLegacyFines) GetEvent() string {
	return e.Event
}

func (e PayLegacyFines) GetTimestamp() time.Time {
	return e.Timestamp
}
