package types

import (
	"github.com/IgaguriMK/ed-journal/event"
	"time"
)

func init() {
	event.RegisterEvent("MissionCompleted", func() event.Event {
		return new(MissionCompleted)
	})
}

type MissionCompleted struct {
	Commodity       string `json:"Commodity"`
	CommodityReward []struct {
		Count int64  `json:"Count"`
		Name  string `json:"Name"`
	} `json:"CommodityReward"`
	CommodityLocalised string `json:"Commodity_Localised"`
	Count              int64  `json:"Count"`
	DestinationStation string `json:"DestinationStation"`
	DestinationSystem  string `json:"DestinationSystem"`
	//Donation           int64     `json:"Donation"`
	Faction        string `json:"Faction"`
	FactionEffects []struct {
		Effects []struct {
			Effect          string `json:"Effect"`
			EffectLocalised string `json:"Effect_Localised"`
			Trend           string `json:"Trend"`
		} `json:"Effects"`
		Faction   string `json:"Faction"`
		Influence []struct {
			SystemAddress int64  `json:"SystemAddress"`
			Trend         string `json:"Trend"`
		} `json:"Influence"`
		Reputation string `json:"Reputation"`
	} `json:"FactionEffects"`
	KillCount     int64     `json:"KillCount"`
	MissionID     int64     `json:"MissionID"`
	Name          string    `json:"Name"`
	Reward        int64     `json:"Reward"`
	TargetFaction string    `json:"TargetFaction"`
	Event         string    `json:"event"`
	Timestamp     time.Time `json:"timestamp"`
}

func (e MissionCompleted) GetEvent() string {
	return e.Event
}

func (e MissionCompleted) GetTimestamp() time.Time {
	return e.Timestamp
}
