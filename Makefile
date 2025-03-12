GO111MODULE ?= on
CGO_ENABLED ?= 0
PLATFORM ?= arm64
VERSION ?= $(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')
COMMIT_HASH ?= $(shell git rev-parse --short HEAD)
BUILD_TIMESTAMP ?= $(shell date '+%Y-%m-%dT%H:%M:%S')

build_prod:
	GO111MODULE=$(GO111MODULE) CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=$(PLATFORM) \
		go build -mod=vendor -ldflags "-s -w \
		-X 'main.Version=$(VERSION)' \
		-X 'main.CommitHash=$(COMMIT_HASH)' \
		-X 'main.BuildTimestamp=$(BUILD_TIMESTAMP)'" \
		-o build/$(PLATFORM)/stacks-tracker cmd/stacks-tracker/main.go
		upx -4 build/$(PLATFORM)/stacks-tracker

build_dev:
	GO111MODULE=$(GO111MODULE) CGO_ENABLED=$(CGO_ENABLED) GOOS=darwin GOARCH=arm64 \
		go build -mod=vendor -ldflags "-s -w \
		-X 'main.Version=$(VERSION)' \
		-X 'main.CommitHash=$(COMMIT_HASH)' \
		-X 'main.BuildTimestamp=$(BUILD_TIMESTAMP)'" \
		-o build/dev/stacks-tracker cmd/stacks-tracker/main.go

dev:
	go run cmd/stacks-tracker/main.go

go_test:
	go generate -v ./...
	go test -v -coverprofile=coverage.out -race -covermode=atomic ./...

go_cover:
	go tool cover -html=coverage.out

.PHONY: dev build_prod clean go_test go_cover
