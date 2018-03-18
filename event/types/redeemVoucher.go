package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("RedeemVoucher", func() event.Event {
		return new(RedeemVoucher)
	})
}

type RedeemVoucher struct {
	Amount   int64  `json:"Amount"`
	Faction  string `json:"Faction"`
	Factions []struct {
		Amount  int64  `json:"Amount"`
		Faction string `json:"Faction"`
	} `json:"Factions"`
	Type      string    `json:"Type"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e RedeemVoucher) GetEvent() string {
	return e.Event
}

func (e RedeemVoucher) GetTimestamp() time.Time {
	return e.Timestamp
}
