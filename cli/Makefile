.PHONY: test

build:
	go build -o cobalt -v

run:
	go build -o cobalt -v
	./cobalt -h
	rm cobalt

test: 
	go test ./test -v