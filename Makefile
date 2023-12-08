TARGET     := kubelogin
OS         := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH       := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))
GOARM      := $(if $(GOARM),$(GOARM),)
BIN         = bin/$(OS)_$(ARCH)$(if $(GOARM),v$(GOARM),)/$(TARGET)
ifeq ($(OS),windows)
  BIN = bin/$(OS)_$(ARCH)$(if $(GOARM),v$(GOARM),)/$(TARGET).exe
endif

GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
GIT_HASH   := $(shell git rev-parse --verify HEAD)
GIT_TAG    := $(shell git describe --tags --exact-match --abbrev=0 2>/dev/null || echo "")
PACKAGE    := github.com/Azure/kubelogin
BUILD_TIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
PLATFORM   := $(OS)/$(ARCH)$(if $(GOARM),v$(GOARM),)

ifdef GIT_TAG
	VERSION := $(GIT_TAG)/$(GIT_HASH)
else
	VERSION := $(GIT_BRANCH)/$(GIT_HASH)
endif

LDFLAGS    := -X $(PACKAGE)/internal/cmd.version=$(VERSION) \
    -X $(PACKAGE)/internal/cmd.goVersion=$(shell go version | cut -d " " -f 3) \
	-X $(PACKAGE)/internal/cmd.buildTime=$(BUILD_TIME) \
	-X '$(PACKAGE)/internal/cmd.platform=$(PLATFORM)'

all: $(TARGET)

lint:
	hack/verify-golangci-lint.sh

test: lint
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

version:
	@echo VERSION: $(VERSION)

$(TARGET): clean
	CGO_ENABLED=0 go build -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/kubelogin

clean:
	-rm -f $(BIN)
