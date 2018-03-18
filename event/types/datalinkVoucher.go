package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("DatalinkVoucher", func() event.Event {
		return new(DatalinkVoucher)
	})
}

type DatalinkVoucher struct {
	PayeeFaction  string    `json:"PayeeFaction"`
	Reward        int64     `json:"Reward"`
	VictimFaction string    `json:"VictimFaction"`
	Event         string    `json:"event"`
	Timestamp     time.Time `json:"timestamp"`
}

func (e DatalinkVoucher) GetEvent() string {
	return e.Event
}

func (e DatalinkVoucher) GetTimestamp() time.Time {
	return e.Timestamp
}
