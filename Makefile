GIT_SUMMARY := $(shell git describe --tags --dirty --always)
GIT_REVISION := $(shell git rev-parse --short HEAD 2> /dev/null || echo 'unknown')
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2> /dev/null || echo 'unknown')

TARGET := vali
TARGET_PATH := cmd/vali/main.go

REPO_PATH := github.com/vlamug/vali
LDFLAGS := -X $(REPO_PATH)/version.Version=$(GIT_SUMMARY)
LDFLAGS += -X $(REPO_PATH)/version.Revision=$(GIT_REVISION)
LDFLAGS += -X $(REPO_PATH)/version.Branch=$(GIT_BRANCH)
BUILDFLAGS := -ldflags "$(LDFLAGS)"

GO := go
GOOS ?= linux
GOARCH ?= amd64

info:
	@echo ">> $(GIT_SUMMARY) $(GIT_REVISION) $(GIT_BRANCH)"

build:
	@echo ">> building binary..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(BUILDFLAGS) -o $(TARGET) $(TARGET_PATH)

test:
	@echo ">> running tests..."
	$(GO) test ./...
