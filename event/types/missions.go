package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("Missions", func() event.Event {
		return new(Missions)
	})
}

type Missions struct {
	Active []struct {
		Expires          int64  `json:"Expires"`
		MissionID        int64  `json:"MissionID"`
		Name             string `json:"Name"`
		PassengerMission bool   `json:"PassengerMission"`
	} `json:"Active"`
	Complete []struct {
		Expires          int64  `json:"Expires"`
		MissionID        int64  `json:"MissionID"`
		Name             string `json:"Name"`
		PassengerMission bool   `json:"PassengerMission"`
	} `json:"Complete"`
	Failed []struct {
		Expires          int64  `json:"Expires"`
		MissionID        int64  `json:"MissionID"`
		Name             string `json:"Name"`
		PassengerMission bool   `json:"PassengerMission"`
	} `json:"Failed"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Missions) GetEvent() string {
	return e.Event
}

func (e Missions) GetTimestamp() time.Time {
	return e.Timestamp
}
