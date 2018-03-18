package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/IgaguriMK/ed-journal/event"
	jfile "github.com/IgaguriMK/ed-journal/file"
)

func main() {
	logfile, err := os.Create("error.log")
	if err != nil {
		os.Exit(1)
	}
	defer logfile.Close()
	logw := io.MultiWriter(logfile, os.Stderr)
	log.SetOutput(logw)
	log.SetFlags(log.Lshortfile)

	////

	jd, err := jfile.JournalDir()
	if err != nil {
		log.Fatal("Can't find journal dir: ", err)
	}

	jfs, err := jfile.JournalFiles(jd)
	if err != nil {
		log.Fatal("Can't search journal files: ", err)
	}

	for _, jf := range jfs {
		sc, err := jf.OpenScanner()
		if err != nil {
			log.Fatal("Can't open files: ", err)
		}

		for sc.Scan() {
			str := sc.Text()

			e, err := event.Parse(str)

			if err != nil {
				switch err := errors.Cause(err).(type) {
				case *event.UnknownEventType:
					saveFailRecord("0.unknown."+err.Type+".", ".json", err.Raw)
				default:
					log.Fatal("Parse error:", err)
				}

				continue
			}

			checkLackOfField(str, e)
		}

		sc.Close()
	}
}

func checkLackOfField(eventStr string, e event.Event) bool {
	bs := []byte(eventStr)
	var v interface{}
	err := json.Unmarshal(bs, &v)
	if err != nil {
		log.Fatal(err)
	}
	bs, err = json.MarshalIndent(v, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	eventStr = string(bs)

	bs, err = json.MarshalIndent(e, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	actualStr := string(bs)

	wantLines := strings.Split(eventStr, "\n")
	getLines := strings.Split(actualStr, "\n")

	if len(wantLines) != len(getLines) {
		saveFailRecord("1.mismatch."+e.GetEvent()+".", ".txt", eventStr+"\n"+actualStr)
		return false
	}

	return true
}

func saveFailRecord(prefix, suffix, content string) {
	hash := sha256.Sum256([]byte(content))
	hashStr := hex.EncodeToString(hash[:])

	name := prefix + hashStr[:10] + suffix
	file, err := os.Create("_out/" + name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Fprint(file, content)
}

type RawEvevtSaver struct {
	file    *os.File
	isFirst bool
}
