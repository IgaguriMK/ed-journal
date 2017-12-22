BASEPKG=github.com/IgaguriMK/ed-journal

.PHONY: all
all: build test

.PHONY: build
build:
	go build $(BASEPKG)/event
	go build $(BASEPKG)/file

.PHONY: test
test:
	go test $(BASEPKG)/event
	go test $(BASEPKG)/file
