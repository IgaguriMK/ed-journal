package event

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

var factory map[string]func() Event

func init() {
	factory = make(map[string]func() Event)
}

func RegisterEvent(typeName string, newFunc func() Event) {
	_, found := factory[typeName]
	if found {
		panic("Already event type exist: " + typeName)
	}

	factory[typeName] = newFunc
}

type Event interface {
	GetEvent() string
	GetTimestamp() time.Time
}

const InvalidEvent = "!!INVALID!!"

// DEPRECATED
func ParseEvent(jsonStr string) (Event, error) {
	return Parse(jsonStr)
}

func Parse(jsonStr string) (Event, error) {
	jsonBytes := []byte(jsonStr)

	eventType := struct {
		Event string `json:"event"`
	}{
		Event: InvalidEvent,
	}

	err := json.Unmarshal(jsonBytes, &eventType)
	if err != nil {
		return nil, errors.Wrap(err, "Failed get event type")
	}

	event, err := parseWithType(jsonBytes, eventType.Event)
	if err != nil {
		return nil, errors.Wrap(err, "Failed parse event")
	}

	return event, nil
}

func parseWithType(bytes []byte, eventType string) (Event, error) {
	f, ok := factory[eventType]
	if ok {
		e := f()
		err := json.Unmarshal(bytes, e)
		return e, err
	}

	//log.Println("[WARNING] Old parse style of ", eventType)

	switch eventType {
	case "Rank":
		var e Rank
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "RebootRepair":
		var e RebootRepair
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ReceiveText":
		var e ReceiveText
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "RedeemVoucher":
		var e RedeemVoucher
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "RefuelAll":
		var e RefuelAll
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Repair":
		var e Repair
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "RepairAll":
		var e RepairAll
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "RestockVehicle":
		var e RestockVehicle
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Resurrect":
		var e Resurrect
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Scan":
		var e Scan
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Scanned":
		var e Scanned
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Screenshot":
		var e Screenshot
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "SelfDestruct":
		var e SelfDestruct
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "SellDrones":
		var e SellDrones
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "SellExplorationData":
		var e SellExplorationData
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "SendText":
		var e SendText
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "SetUserShipName":
		var e SetUserShipName
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ShieldState":
		var e ShieldState
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ShipyardBuy":
		var e ShipyardBuy
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ShipyardNew":
		var e ShipyardNew
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ShipyardSell":
		var e ShipyardSell
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ShipyardSwap":
		var e ShipyardSwap
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ShipyardTransfer":
		var e ShipyardTransfer
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "StartJump":
		var e StartJump
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "SupercruiseEntry":
		var e SupercruiseEntry
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "SupercruiseExit":
		var e SupercruiseExit
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Synthesis":
		var e Synthesis
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Touchdown":
		var e Touchdown
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Undocked":
		var e Undocked
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "USSDrop":
		var e USSDrop
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "WingAdd":
		var e WingAdd
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "WingInvite":
		var e WingInvite
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "WingJoin":
		var e WingJoin
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "WingLeave":
		var e WingLeave
		err := json.Unmarshal(bytes, &e)
		return e, err

		///////////////////////////////////////////////////////////
	}

	return nil, &UnknownEventType{Raw: string(bytes), Type: eventType}
}

type UnknownEventType struct {
	Type string
	Raw  string
}

func (e *UnknownEventType) Error() string {
	return fmt.Sprintf("Unknown event type [%s]: %s", e.Type, e.Raw)
}
