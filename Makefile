.PHONY: run r build b test t

run r: build
	./tdi

build b:
	go build .

test t:
	go test ./... -v
