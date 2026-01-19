.PHONY: all build __build __build-dev

VERSION:=v0.0.1
# COMMIT_SHA:=$(shell git rev-parse main)
COMMIT_SHA:="blyat"
BINARY_NAME:=azrael
LDFLAGS:=-X main.Version=$(VERSION) -X main.CommitSHA=$(COMMIT_SHA)

all:
	@echo "Usage: make build-os-arch e.g. make build-linux-amd64"
	@echo "To build for all operating systems and architecture (for some reason) use build-all"

build-linux-amd64:
	@make __build OS=linux ARCH=amd64

build-linux-i386:
	@make __build OS=linux ARCH=386

build-linux-arm64:
	@make __build OS=linux ARCH=arm64

build-linux-arm32:
	@make __build OS=linux ARCH=arm

build-windows-amd64:
	@make __build OS=windows ARCH=amd64

build-macos-amd64:
	@make __build OS=darwin ARCH=amd64

build-macos-arm64:
	@make __build OS=darwin ARCH=arm64

build-all:
	@make build-linux-amd64
	@make build-linux-i386
	@make build-linux-arm64
	@make build-linux-arm32
	@make build-windows-amd64
	@make build-macos-amd64
	@make build-macos-arm64

__build:
ifndef OS
	$(error OS is not defined.)
endif
ifndef ARCH
	$(error ARCH is not defined.)
endif
	@[[ "$(OS)" == "windows" ]] && EXTENSION=".exe" || EXTENSION=""
	@echo "Building $(BINARY_NAME) for $(OS)/$(ARCH)!"
	@GOOS=$(OS) GOARCH=$(ARCH) go build -ldflags="-w -s $(LDFLAGS)" -o ./$(BINARY_NAME)-$(VERSION)-$(OS)-$(ARCH)$(EXTENSION) ./cmd/cli/...

build-dev:
	@echo "Building $(BINARY_NAME) for host OS/Arch!"
	@[[ "$(OS)" == "windows" ]] && EXTENSION=".exe" || EXTENSION=""
	@go build -ldflags="$(LDFLAGS)" -o ./$(BINARY_NAME)$(EXTENSION) ./cmd/cli/...

clean:
	go clean
	rm -f ./$(BINARY_NAME)*
