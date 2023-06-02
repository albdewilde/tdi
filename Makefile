.PHONY: run build

run: build
	./tdi

build:
	go build .
