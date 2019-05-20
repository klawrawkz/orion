.PHONY: test

build:
	go build -o orion -v

run:
	go build -o orion -v
	./orion -h
	rm orion

test: 
	go test ./test -v