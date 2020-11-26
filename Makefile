.PHONY: build

build:
	go build -o ${GOPATH}/bin/steno ./cli/main.go
