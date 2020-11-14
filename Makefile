.PHONY: build

build:
	go build -o ${GOPATH}/bin/steno ./steno-cli/main.go
