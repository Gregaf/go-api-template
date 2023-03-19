.PHONY: help clean go-test $(goTestDirs) $(binFiles) $(logFiles)
.SHELL := /bin/bash
.DEFAULT_GOAL := help

PROGRAM_NAME := $(shell basename $(shell pwd))
USER := $(shell whoami)

CHARACTER_LEN_SET := $(shell seq 80)
DIVIDER := $(shell for i in $(CHARACTER_LEN_SET); do printf "%s" "="; done)

# Output Colors
NEUTRAL := \033[0m
WHITE := \033[1;37m
RED := \033[0;31m
GREEN := \033[0;32m
CYAN := \033[0;36m

baseDir := $(shell pwd)
appPath := $(baseDir)/cmd/main.go

binDir := $(baseDir)/dist
logDir := $(baseDir)/make_logs
airDir := $(baseDir)/tmp

goTestDirs = $(addprefix item-, $(shell find $$PWD -name '*_test.go' -exec dirname "{}" \; | sort -u))

formatLogFile = $(shell echo $(logDir)/$$(basename $1)_$2.log)
currentDate := $(shell date)

timeCmd = /usr/bin/time --format "$1: %E" $2

go-start: ## Start the basic Golang app without Docker
	go run $(appPath)

goPlatforms := linux/amd64 darwin/amd64 windows/amd64
GOOS =	$(word 1,$(subst /, ,$1))
GOARCH = $(word 2,$(subst /, ,$1))

$(goPlatforms):
	@echo "$(WHITE)Building '$@'$(NEUTRAL)"
	@echo $(DIVIDER)
	@GOOS=$(call GOOS,$@) GOARCH=$(call GOARCH,$@) \
	$(call timeCmd,Build Time,go build -a -o $(binDir)/$(PROGRAM_NAME)-$(subst /,-,$@) $(appPath))
	@echo "$(GREEN)Build successfully completed$(NEUTRAL)"
	@echo ""

go-build: _create_log_dir $(goPlatforms)## Build the basic Golang app without Docker

_create_log_dir:
	@mkdir -p $(logDir)

go-test: _create_log_dir $(goTestDirs) ## Run all the basic Golang app tests without Docker
$(goTestDirs): item-%:
	@echo "$(WHITE)Testing '$*'$(NEUTRAL)"
	@echo $(DIVIDER)
	@echo "[$(currentDate)]" >> $(call formatLogFile,$*,tests)
	@go test -v $* | tee -a $(logDir)/$$(basename $*)_tests.log

docker-build: ## Build the current Dockerfile with Docker build
	docker build -t $(PROGRAM_NAME):latest .

docker-start: ## Start the basic Golang app with Docker compose
	docker compose up $(DOCK_OPTS)

docker-create: DOCK_OPTS := -d
docker-create: docker-start ## Start the basic Golang app with Docker compose

docker-destroy: ## Stop the basic Golang app with Docker compose
	docker compose down

# There has to be a better way to handle tmp directory ownership...
clean: ## Remove all the build files
	go clean
	rm -rf $(binDir)
	rm -rf $(logDir)
	if [ -d $(airDir) ]; then \
	sudo chown -R $(USER):$(USER) $(airDir); \
	fi
	rm -rf $(airDir)

help: ## Provides a help menu
	@echo ""
	@echo "$(WHITE)Application: $(PROGRAM_NAME)$(NEUTRAL)"
	@echo $(DIVIDER)
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'