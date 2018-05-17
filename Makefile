.PHONY: build
build:
	go build checkjournals.go

.PHONY: clean
clean:
	- rm *.exe
	- rm checkjournals
	- rm -r _out
