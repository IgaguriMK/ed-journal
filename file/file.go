package file

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type JournalFile struct {
	name      string
	dirPath   string
	startTime time.Time
}

var journalExpr = regexp.MustCompile(`^Journal\.[0-9]{12}\.[0-9]{2}\.log$`)

func JournalFiles(findPath string) ([]JournalFile, error) {
	journals := make([]JournalFile, 0)

	files, err := ioutil.ReadDir(findPath)
	if err != nil {
		return journals, fmt.Errorf("Can't open journal dir.\n    ", err)
	}

	for _, f := range files {
		if journalExpr.MatchString(f.Name()) {
			jt, err := journalNameToTime(f.Name())
			if err != nil {
				continue
			}
			journals = append(
				journals,
				JournalFile{
					name:      f.Name(),
					dirPath:   findPath,
					startTime: jt,
				})
		}
	}

	return journals, nil
}

func JournalDir() (string, error) {
	jd := os.Getenv("USERPROFILE")
	if jd == "" {
		return "", errors.New("%USERPROFILE% is Empty")
	}
	return jd + `\Saved Games\Frontier Developments\Elite Dangerous`, nil
}

const journalTime = "060102150405"

func journalNameToTime(name string) (time.Time, error) {
	name = strings.TrimPrefix(name, "Journal.")
	name = strings.TrimSuffix(name, ".01.log")

	return time.ParseInLocation(journalTime, name, time.Local)
}

func (j *JournalFile) Name() string {
	return j.name
}

func (j *JournalFile) FullPath() string {
	return j.dirPath + `\` + j.name
}

func (j *JournalFile) StartAt() time.Time {
	return j.startTime
}

func (j *JournalFile) Open() (*os.File, error) {
	return os.Open(j.FullPath())
}

type Scanner struct {
	sc   *bufio.Scanner
	file *os.File
}

func (j *JournalFile) OpenScanner() (*Scanner, error) {
	file, err := j.Open()
	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(file)

	return &Scanner{
		sc:   sc,
		file: file,
	}, nil
}

func (js *Scanner) Scan() bool {
	return js.sc.Scan()
}

func (js *Scanner) Text() string {
	return js.sc.Text()
}

func (js *Scanner) Bytes() []byte {
	return js.sc.Bytes()
}

func (js *Scanner) Err() error {
	return js.sc.Err()
}

func (js *Scanner) Close() {
	js.file.Close()
}
