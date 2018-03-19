package stream

import (
	"log"

	"github.com/pkg/errors"

	"github.com/IgaguriMK/ed-journal/event"
	"github.com/IgaguriMK/ed-journal/file"
)

func JournalStream() (chan event.Event, error) {
	jd, err := file.JournalDir()
	if err != nil {
		return nil, errors.Wrap(err, "Can't get Journal dir.")
	}

	jfs, err := file.JournalFiles(jd)
	if err != nil {
		return nil, errors.Wrap(err, "Can't get Journal files.")
	}

	ch := make(chan event.Event)

	go func() {
		for _, jf := range jfs {
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
						log.Fatal("JSON decode error:", err)
					}
				}
				ch <- e
			}

			sc.Close()
		}

		close(ch)
	}()

	return ch, nil
}
