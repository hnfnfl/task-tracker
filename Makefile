.PHONY: all build test clean

GOBUILD=go build
GOCLEAN=go clean
BINARY_NAME=task-tracker

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) main.go

test: build
	./$(BINARY_NAME) add -desc "test task"

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)