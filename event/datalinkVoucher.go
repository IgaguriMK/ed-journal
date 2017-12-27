package event

import "time"

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
