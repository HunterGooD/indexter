GOOS?=linux
GOARCH?=amd64
GOPATH?=~/go

.PHONY: build clean easyjson

build: clean
	mkdir -p ./dist/bin
	go build -o ./dist/bin/indexter ./cmd/indexter/main.go

clean:
	rm -rf ./dist

easyjson:
	easyjson -all internal/models/model.go
