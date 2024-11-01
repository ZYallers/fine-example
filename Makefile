# Cross compilation may require setting GO111MODULE
export GO111MODULE=on

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GO_BUILD_FLAGS ?= -a
GO_BUILD_LDFLAGS ?= -w

GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GO_BUILD_LDFLAGS += -X 'main.BuildDate=$(shell date '+%Y-%m-%d %H:%M:%S')'
GO_BUILD_LDFLAGS += -X 'main.GitBranch=$(GIT_BRANCH)'
GO_BUILD_LDFLAGS += -X 'main.GitCommit=$(GIT_COMMIT)'

GO := go
MAKE := make
MAIN_FILE := main.go
PROJECT_ROOT := $(shell pwd)
PROJECT_NAME := $(shell basename "$(PWD)")
BIN_DIR := $(PROJECT_ROOT)/bin
LOG_DIR := $(PROJECT_ROOT)/logs

BINARY_NAME := $(PROJECT_NAME).$(GOOS).$(GOARCH)
ifeq ($(GOOS), windows)
	BINARY_NAME := $(addsuffix .exe, $(BINARY_NAME))
endif

# Build all by default, even if it's not first
.DEFAULT_GOAL := all

.PHONY: all
all: help

## go.fmt: Gofmt (reformat) package sources.
.PHONY: go.fmt
go.fmt:
	@echo "> Formatting project with $(PROJECT_NAME)..."
	$(GO) fmt $(go list ./... | grep -v /vendor/)

## go.deps: Add missing modules and delete unused modules to generate a directory of dependencies.
.PHONY: go.deps
go.deps: go.mod.tidy go.mod.vendor

## go.test: Test packages.
.PHONY: go.test
go.test:
	@echo "> Testing project with $(PROJECT_NAME)..."
	$(GO) test $(PROJECT_ROOT)/tests/...

## go.build: Compile the project and generate an executable file.
.PHONY: go.build
go.build:
	@echo "> Building binary $(BINARY_NAME)..."
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) \
	$(GO) build $(GO_BUILD_FLAGS) -ldflags "$(GO_BUILD_LDFLAGS)" \
	-o $(BIN_DIR)/$(BINARY_NAME) $(PROJECT_ROOT)/$(MAIN_FILE)

## go.run: Run packages.
.PHONY: go.run
go.run:
	@echo "> Running server with $(MAIN_FILE)..."
	$(GO) run -race $(PROJECT_ROOT)/$(MAIN_FILE)

## server.start: Start server.
.PHONY: server.start
server.start:
	@echo "> Starting server with $(BINARY_NAME)..."
	@pid=$$(pgrep "$(BINARY_NAME)" 2>/dev/null); \
	if [ -n "$$pid" ]; then \
		echo "Server has been started, PID: $$pid, please stop it first"; \
    else \
		mkdir -p $(LOG_DIR)/$(PROJECT_NAME); \
		nohup $(BIN_DIR)/$(BINARY_NAME) >> $(LOG_DIR)/$(PROJECT_NAME)/$(PROJECT_NAME).log 2>&1 & \
		pid=$$(echo $$!); \
		if [ -n "$$pid" ]; then \
			echo "Server started, PID: $$pid"; \
		else \
			echo "Failed to start server"; \
		fi; \
	fi

## server.stop: Stop server.
.PHONY: server.stop
server.stop:
	@echo "> Stopping server with $(BINARY_NAME)..."
	@for i in {1..31}; do \
		pid=$$(pgrep "$(BINARY_NAME)" 2>/dev/null); \
		if [ -n "$$pid" ]; then \
		    if [ $$i -eq 31 ]; then \
				echo "Forcibly kill process with PID: $$pid"; \
				kill -9 $$pid; \
				break; \
			else \
			  	echo "Killing process with PID: $$pid"; \
				kill $$pid; \
			fi; \
		else \
			echo "Server process not found"; \
			break; \
		fi; \
		sleep 1; \
	done

## server.reload: Reload server.
.PHONY: server.reload
server.reload: server.stop server.start

## server.status: Show server status.
.PHONY: server.status
server.status:
	@echo "> Checking server status with $(BINARY_NAME)..."
	@pid=$$(pgrep "$(BINARY_NAME)" 2>/dev/null); \
	if [ -n "$$pid" ]; then \
		echo "Server is running, PID: $$pid"; \
		ps -p $$pid -o user,pid,ppid,%cpu,%mem,vsz,rss,tty,stat,start,time,command; \
	else \
		echo "Server is not running"; \
	fi

## clean: Clean package.
.PHONY: clean
clean:
	@echo "> Cleaning project with $(PROJECT_NAME)..."
	rm -f $(BIN_DIR)/$(BINARY_NAME)
	rm -rf $(PROJECT_ROOT)/vendor

## swag.fmt: Format Swag interface comments.
.PHONY: swag.fmt
swag.fmt: tools.verify.swag
	@echo "> Formatting swag comment..."
	@CGO_ENABLED="0" swag fmt --dir "./internal/controller"

## swag.gen: Generate Swagger.json interface document.
.PHONY: swag.gen
swag.gen: tools.verify.swag
	@echo "> Generating api document json file..."
	@CGO_ENABLED="0" \
	swag init \
	  --dir "./internal/controller" \
	  --generalInfo "../../main.go" \
	  --output "./doc" \
	  --outputTypes "json" \
	  --parseDependency

.PHONY: go.mod.tidy
go.mod.tidy:
	@echo "> Add missing and remove unused modules..."
	$(GO) mod tidy

.PHONY: go.mod.vendor
go.mod.vendor:
	@echo "> Make vendored copy of dependencies..."
	$(GO) mod vendor

.PHONY: install.swag
install.swag:
	@$(GO) install github.com/swaggo/swag/cmd/swag@v1.16.2

.PHONY: tools.install.%
tools.install.%:
	@echo "===========> Installing $*"
	@$(MAKE) install.$*

.PHONY: tools.verify.%
tools.verify.%:
	@if which $* &>/dev/null; then echo "> Verify ${*} has been installed!"; else $(MAKE) tools.install.$*; fi


define USAGE_OPTIONS
Options:
  GOOS                GO compilation system, defaults to the current system.
  GOARCH              GO compilation architecture, defaults to the current architecture.
  GO_BUILD_FLAGS      GO compilation additional parameters, default is -a.
  GO_BUILD_LDFLAGS    GO compilation link parameter, default is -w.
 endef
export USAGE_OPTIONS

## help: Show help message.
.PHONY: help
help: Makefile
	@echo "Choose a command run in project "$(PROJECT_NAME)":"
	@echo "Usage: make <TARGETS> <OPTIONS> ...\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"

