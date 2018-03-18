BASEPKG=github.com/IgaguriMK/ed-journal

.PHONY: check
check: build testrun

.PHONY: build
build:
	go build checkjournals.go

.PHONY: testrun
testrun:
	./checkjournals
