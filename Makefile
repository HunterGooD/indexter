GOOS?=linux
GOARCH?=amd64
GOPATH?=~/go

.PHONY: build clean easyjson

build: clean
	mkdir -p ./dist/bin
	go build -o ./dist/bin/linux/indexter ./cmd/indexter/main.go
	GOOS=windows go build -o ./dist/bin/win/indexter.exe ./cmd/indexter/main.go

clean:
	rm -rf ./dist

easyjson:
	easyjson -all internal/models/model.go
