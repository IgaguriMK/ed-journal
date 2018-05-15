package stream

import (
	"log"
	"time"

	"github.com/pkg/errors"

	"github.com/IgaguriMK/ed-journal/event"
	"github.com/IgaguriMK/ed-journal/event/types"
	"github.com/IgaguriMK/ed-journal/file"
)

func JournalStream(filters ...EventFilter) (chan event.Event, error) {
	jd, err := file.JournalDir()
	if err != nil {
		return nil, errors.Wrap(err, "Can't get Journal dir.")
	}

	jfs, err := file.JournalFiles(jd)
	if err != nil {
		return nil, errors.Wrap(err, "Can't get Journal files.")
	}

	ch := make(chan event.Event)

	rch := ch
	for _, f := range filters {
		rch = f.Apply(rch)
	}

	go func() {
		for _, jf := range jfs {
			ch <- &FileStart{
				start: jf.StartAt(),
			}

			sc, err := jf.OpenScanner()
			if err != nil {
				log.Fatal("Can't open file:", err)
			}

			for sc.Scan() {
				e, err := event.Parse(sc.Text())
				if err != nil {
					switch err := errors.Cause(err).(type) {
					case *event.UnknownEventType:
						log.Println("Unknown event detected:", err)
						continue
					default:
						log.Println("JSON decode error:", err)
					}
				}
				ch <- e
			}

			sc.Close()
		}

		close(ch)
	}()

	return rch, nil
}

type FileStart struct {
	start time.Time
}

func (_ FileStart) GetEvent() string {
	return "_FileStart_"
}

func (e FileStart) GetTimestamp() time.Time {
	return e.start
}

////

type EventFilter interface {
	Apply(ch chan event.Event) chan event.Event
}

type CmdrFilter struct {
	name string
}

func NewCmdrFilter(name string) *CmdrFilter {
	return &CmdrFilter{
		name: name,
	}
}

func (f CmdrFilter) Apply(ch chan event.Event) chan event.Event {
	nch := make(chan event.Event)

	go func() {
		defer close(nch)
		send := false

		for {
			// 本体ループ
			for {
				e, ok := <-ch
				if !ok {
					return
				}

				if fs, fsok := e.(*FileStart); fsok {
					nch <- fs
					break
				} else {
					if send {
						nch <- e
					}
				}
			}

			// ヘッダ待ちループ
			queue := make([]event.Event, 0)
			for {
				e, ok := <-ch
				if !ok {
					return
				}

				if lg, lgok := e.(*types.LoadGame); lgok {
					if lg.Commander == f.name {
						sendAll(nch, queue)
						nch <- lg
						send = true
						break
					} else {
						send = false
						break
					}
				} else {
					queue = append(queue, e)
				}
			}
		}

	}()

	return nch
}

func sendAll(ch chan event.Event, events []event.Event) {
	for _, e := range events {
		ch <- e
	}
}
