
# The name of the executable (default is current directory name)
TARGET := `basename ${PWD}`
.DEFAULT_GOAL: build

# These will be provided to the target
VERSION := 1.0.0
BUILD := `git rev-parse HEAD`
PROJECT := essence
# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-s -X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all build clean install uninstall run docker/build docker/push

all: build

deps:
	go get -f -u ./...

build: prepare
	CGO_ENABLED=0 go build $(LDFLAGS) -o $(TARGET)

clean:
	@rm -f $(TARGET)

install: build
	@go install $(LDFLAGS)

uninstall: clean
	@rm -f $$(which ${TARGET})

run: install
	@$(TARGET)

docker/build: build
	@docker build . -t $(PROJECT)/$(TARGET)

docker/push: docker/build
	@docker push $(PROJECT)/$(TARGET)
