PROJECT_NAME := ticker-upload-to-hub
BASE_NAME := $(PROJECT_NAME)
SERVER_NAME := ./bin/$(BASE_NAME)
VERSION := $(shell git describe --always --long --dirty)
COMMITID := $(shell git rev-parse --short HEAD)
GOVERSION := $(shell go version)
TIME := $(shell date +"%Y%m%d%H%M%S")
PROGRAM := $(SERVER_NAME).$(TIME)

.PHONY: all clean

all:
	go build -o ${SERVER_NAME} -gcflags "-N -l" \
		-ldflags="-X main.BinVersion=$(VERSION) -X main.CommitID=$(COMMITID) -X main.BuildTime=$(TIME) -X \"main.GoVersion=$(GOVERSION)\"" .
	cp $(SERVER_NAME) $(PROGRAM) 
clean:
	rm -rf bin/* log/*
