package stream

import (
	"encoding/json"
	"io"
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
			f, err := jf.Open()
			if err != nil {
				log.Fatal("Can't open file:", err)
			}

			dec := json.NewDecoder(f)

			for {
				var e event.Event
				err = dec.Decode(&e)
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal("JSON decode error:", err)
				}
				ch <- e
			}

			f.Close()
		}

		close(ch)
	}()

	return ch, nil
}
