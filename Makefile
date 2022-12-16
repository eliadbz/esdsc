FILE="esdsc.go"
BIN="esdsc"
CWD=$(shell pwd)
PREFIX=/usr/local

tidy:
	go mod tidy

build: tidy
	go build -o ./debug/$(BIN) $(FILE)

clean:
	go clean -modcache

run: build
	./bin/$(BIN)

release: tidy
	go build -ldflags "-s -w" -o ./bin/$(BIN) $(FILE)

install: release
	install -d $(PREFIX)/bin/
	install ./bin/$(BIN) $(PREFIX)/bin/

uninstall:
	rm -f $(PREFIX)/bin/$(BIN)