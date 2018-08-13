BUILD=go build
VERSION := $(shell git describe --abbrev=4 --dirty --always --tags)
Minversion := $(shell date)
BUILD_IOEX_CLI = -ldflags "-X main.Version=$(VERSION)"

all:
	$(BUILD) $(BUILD_IOEX_CLI) ioex-cli.go

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(BUILD) $(BUILD_IOEX_CLI) ioex-cli.go
