GOFILES=$(wildcard *.go)
BINARY_NAME=challenge

all: dep build

dep:
	@echo "Installing dependencies..."
	go get -u github.com/urfave/cli
	go get -u github.com/omise/omise-go
	go get -u github.com/leekchan/accounting

build:
	@echo "Building the project"
	mkdir -p bin
	go build -o bin/$(BINARY_NAME) $(GOFILES)

help:
	bin/$(BINARY_NAME) -h

example:
	bin/$(BINARY_NAME) -f data/fng.1000.csv.rot128 -w 100 -r 50

test:
	go test -v $(shell go list ./...)

clean:
	@echo "Cleaning..."
	rm -rf bin
	rm -rf pkg
	rm -rf src
	rm -rf vendor
	go clean



