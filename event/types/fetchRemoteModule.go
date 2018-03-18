package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("FetchRemoteModule", func() event.Event {
		return new(FetchRemoteModule)
	})
}

type FetchRemoteModule struct {
	ServerID            int64     `json:"ServerId"`
	Ship                string    `json:"Ship"`
	ShipID              int64     `json:"ShipID"`
	StorageSlot         int64     `json:"StorageSlot"`
	StoredItem          string    `json:"StoredItem"`
	StoredItemLocalised string    `json:"StoredItem_Localised"`
	TransferCost        int64     `json:"TransferCost"`
	TransferTime        int64     `json:"TransferTime"`
	Event               string    `json:"event"`
	Timestamp           time.Time `json:"timestamp"`
}

func (e FetchRemoteModule) GetEvent() string {
	return e.Event
}

func (e FetchRemoteModule) GetTimestamp() time.Time {
	return e.Timestamp
}
