# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Name of the executable
BINARY_NAME=tailflare

all: cd test build

build:
	cd src && $(GOBUILD) -o ../build/$(BINARY_NAME) -v

test:
	cd src $(GOTEST) -v ./...

clean: cd
	cd src && $(GOCLEAN)
	rm -f $(BINARY_NAME)
