# Go parameters
GOCMD=go
GOMAIN=./app
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run $(GOMAIN)
GOCLEAN=$(GOCMD) clean $(GOMAIN)
GOTEST=$(GOCMD) test $(GOMAIN)
GOGET=$(GOCMD) get $(GOMAIN)

# Name of the executable
BINARY_NAME=tailflare

all: cd test build

run:
	$(GORUN)

build:
	cd app && $(GOBUILD) -o ../../build/$(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean: cd
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
