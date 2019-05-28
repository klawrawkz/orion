.PHONY: test

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt
BINARY_NAME=orion

all: test build
build: 
	$(GOFMT)
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v -cover ./... 
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
