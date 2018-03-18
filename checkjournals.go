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
	_ "github.com/IgaguriMK/ed-journal/event/types"
	jfile "github.com/IgaguriMK/ed-journal/file"
)

const MaxFail = 1000000

var failCount = 0

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

	if err := os.RemoveAll("./_out"); err != nil {
		log.Fatal("Can't remove dir: ", err)
	}
	if err := os.MkdirAll("./_out/unknown", 0644); err != nil {
		log.Fatal("Can't create dir: ", err)
	}
	if err := os.MkdirAll("./_out/mismatch", 0644); err != nil {
		log.Fatal("Can't create dir: ", err)
	}

	totalCount := 0
	unknownCount := 0
	mismatchCount := 0

	for _, jf := range jfs {
		sc, err := jf.OpenScanner()
		if err != nil {
			log.Fatal("Can't open files: ", err)
		}

		for sc.Scan() {
			totalCount++

			str := sc.Text()
			e, err := event.Parse(str)

			if err != nil {
				switch err := errors.Cause(err).(type) {
				case *event.UnknownEventType:
					saveFailRecord("unknown/"+err.Type+".", ".json", err.Raw)
					unknownCount++
					failed()
				default:
					log.Fatal("Parse error:", err)
				}

				continue
			}

			ok := checkLackOfField(str, e)
			if !ok {
				mismatchCount++
				failed()
			}
		}

		sc.Close()
	}

	fmt.Printf("Unknown:    %d (%.1f %%)\n", unknownCount, 100.0*float64(unknownCount)/float64(totalCount))
	fmt.Printf("Mismatches:    %d (%.1f %%)\n", mismatchCount, 100.0*float64(mismatchCount)/float64(totalCount))
}

func failed() {
	failCount++

	if failCount >= MaxFail {
		log.Fatal("Too many fail.")
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

	getMap := make(map[string]bool)
	for _, g := range getLines {
		g = strings.Trim(g, "\t")
		getMap[g] = true
	}

	failed := false
	lackedLines := make([]string, 0)
	for _, w := range wantLines {
		w = strings.Trim(w, "\t")
		if !getMap[w] {
			failed = true
			lackedLines = append(lackedLines, w)
		}
	}

	if failed {
		saveFailRecord(
			"mismatch/"+e.GetEvent()+".", ".got.txt",
			fmt.Sprintf(
				"GET:\n%s\nWANT:%s\nNEED:\n%s\n",
				actualStr,
				eventStr,
				strings.Join(lackedLines, "\n"),
			),
		)
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
