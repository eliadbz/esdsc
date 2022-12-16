FILE="esdsc.go"
BIN="esdsc"
CWD=$(shell pwd)

tidy:
	go mod tidy

build: tidy
	go build -o ./bin/$(BIN) $(FILE)

clean:
	go clean -modcache

run: build
	./bin/$(BIN)