package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/IgaguriMK/ed-journal/event"
	_ "github.com/IgaguriMK/ed-journal/event/types"
	jfile "github.com/IgaguriMK/ed-journal/file"
)

var MaxFail = 1000
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

	if err := removeAllIn("./_out/unknown"); err != nil {
		log.Fatal("Can't remove dir unknown: ", err)
	}
	if err := removeAllIn("./_out/mismatch"); err != nil {
		log.Fatal("Can't remove dir mismatch: ", err)
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
					saveFailRecord("unknown/"+err.Type+".", ".json", err.Raw+"\n")
					unknownCount++
					failed()
				default:
					saveFailRecord("parseError.", ".json", str+"\n")
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

	fmt.Printf("Unknown:    %d (%.2f %%)\n", unknownCount, 100.0*float64(unknownCount)/float64(totalCount))
	fmt.Printf("Mismatches:    %d (%.2f %%)\n", mismatchCount, 100.0*float64(mismatchCount)/float64(totalCount))

	if unknownCount > 0 || mismatchCount > 0 {
		os.Exit(1)
	}
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
		g = strings.TrimSuffix(g, ",")
		getMap[g] = true
	}

	failed := false
	lackedLines := make([]string, 0)
	for _, w := range wantLines {
		w = strings.TrimSuffix(w, ",")
		if !getMap[w] {
			failed = true
			lackedLines = append(lackedLines, w)
		}
	}

	if failed {
		saveFailRecord(
			"mismatch/"+e.GetEvent()+".", ".txt",
			fmt.Sprintf(
				"GET:\n%s\nWANT:\n%s\nNEED:\n%s\n",
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

func removeAllIn(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		fname := filepath.Join(dir, f.Name())
		err := os.Remove(fname)
		if err != nil {
			return err
		}
	}

	return nil
}
