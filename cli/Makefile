.PHONY: test

build:
	go fmt
	go build -o orion -v

run:
	go build -o orion -v
	./orion -h
	rm orion

test: 
	go test ./... -v -cover

clean:
	rm ./orion