# Go settings
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOLINT=golint
GOGET=$(GOCMD) get

# Build settings
BINARY_PATH=./bin/
BINARY_NAME=ygrep
BINARY_LINUX=$(BINARY_NAME)_linux
BINARY_WINDOWS=$(BINARY_NAME).exe
BINARY_MACOS=$(BINARY_NAME)_macos

# Test Settings
TEST_FILES := $(shell $(GOCMD) list ./...)

all: deps tests build


build:
		$(GOBUILD) -o $(BINARY_PATH)$(BINARY_NAME) -v cmd/ygrep/main.go

tests:
		mkdir -p report
		$(GOTEST) -v -short -covermode=count -coverprofile report/cover.out $(TEST_FILES)
		$(GOCMD) tool cover -html=report/cover.out -o report/cover.html
		$(GOLINT) -set_exit_status $(TEST_FILES)
		CC=clang $(GOTEST) -v -msan -short $(TEST_FILES)
		staticcheck $(TEST_FILES)

clean:
		$(GOCLEAN)
		rm -rf $(BINARY_PATH)
		rm -rf ./report/

deps:
		GO111MODULE=on $(GOCMD) mod vendor
