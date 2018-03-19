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

	return nil, &UnknownEventType{Raw: string(bytes), Type: eventType}
}

type UnknownEventType struct {
	Type string
	Raw  string
}

func (e *UnknownEventType) Error() string {
	return fmt.Sprintf("Unknown event type [%s]: %s", e.Type, e.Raw)
}
