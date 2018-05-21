package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
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

var MaxFail int

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
	var dir string
	flag.StringVar(&dir, "d", "", "Check set dir.")
	var autoFind bool
	flag.BoolVar(&autoFind, "autofind", false, "Auto find default journal dir.")
	flag.IntVar(&MaxFail, "maxfail", 100000, "Max fail count.")

	flag.Parse()

	if err := removeAllIn("./_out/unknown"); err != nil {
		log.Println("Can't remove dir unknown: ", err)
	}
	if err := removeAllIn("./_out/mismatch"); err != nil {
		log.Println("Can't remove dir mismatch: ", err)
	}
	if err := os.MkdirAll("./_out/unknown", 0744); err != nil {
		log.Fatal("Can't create dir: ", err)
	}
	if err := os.MkdirAll("./_out/mismatch", 0744); err != nil {
		log.Fatal("Can't create dir: ", err)
	}

	jds := make([]string, 0)

	if dir == "" || autoFind {
		jd, err := jfile.JournalDir()
		if err != nil {
			log.Fatal("Can't find journal dir: ", err)
		}
		jds = append(jds, jd)
	}

	if dir != "" {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal("Can't read dir: ", err)
		}

		for _, f := range files {
			if f.IsDir() {
				jds = append(jds, filepath.Join(dir, f.Name()))
			}
		}
	}

	c := new(Counter)

	for _, jd := range jds {
		c.Check(jd)
	}

	c.Summary()

	if c.Failed() {
		os.Exit(1)
	}
}

type Counter struct {
	totalCount    int
	unknownCount  int
	mismatchCount int
}

func (c *Counter) Check(jd string) {
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
			c.total()

			str := sc.Text()
			e, err := event.Parse(str)

			if err != nil {
				switch err := errors.Cause(err).(type) {
				case *event.UnknownEventType:
					saveFailRecord("unknown/"+err.Type+".", ".json", err.Raw+"\n")
					c.unknown()
				default:
					saveFailRecord("parseError.", ".json", str+"\n")
					log.Fatal("Parse error:", err)
				}

				continue
			}

			ok := checkLackOfField(str, e)
			if !ok {
				c.mismatch()
			}
		}

		sc.Close()
	}
}

func (c *Counter) total() {
	c.totalCount++
}

func (c *Counter) mismatch() {
	c.mismatchCount++

	if c.mismatchCount+c.unknownCount >= MaxFail {
		log.Fatal("Too many fail.")
	}
}

func (c *Counter) unknown() {
	c.unknownCount++

	if c.mismatchCount+c.unknownCount >= MaxFail {
		log.Fatal("Too many fail.")
	}
}

func (c *Counter) Summary() {
	fmt.Printf("Unknown:    %d (%.2f %%)\n", c.unknownCount, 100.0*float64(c.unknownCount)/float64(c.totalCount))
	fmt.Printf("Mismatches:    %d (%.2f %%)\n", c.mismatchCount, 100.0*float64(c.mismatchCount)/float64(c.totalCount))
	fmt.Printf("Total: %d\n", c.total)
}

func (c *Counter) Failed() bool {
	return c.unknownCount > 0 || c.mismatchCount > 0
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
