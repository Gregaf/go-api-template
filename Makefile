.PHONY: help clean go-test $(goTestDirs) $(binFiles) $(logFiles)
.SHELL := /bin/bash
.DEFAULT_GOAL := help

PROGRAM_NAME := $(shell basename $(shell pwd))

ARCH := $(shell uname -m)
OS := $(shell uname -s | tr A-Z a-z)
CHARACTER_LEN_SET := $(shell seq 80)

baseDir := $(shell pwd)
appPath := $(baseDir)/cmd/main.go

binDir := $(baseDir)/dist
logDir := $(baseDir)/make_logs

goTestDirs = $(addprefix item-, $(shell find $$PWD -name '*.go' -exec dirname "{}" \;))

formatTestLog = $(shell echo $(logDir)/$$(basename $1)_tests.log)
currentDate := $(shell date)

go-start: ## Start the basic Golang app without Docker
	go run $(appPath)

go-build: ## Build the basic Golang app without Docker

go-test: $(goTestDirs) ## Run all the basic Golang app tests without Docker
$(goTestDirs): item-%:
	@echo "Testing '$*'"
	@mkdir -p $(logDir)
	@echo "[$(currentDate)]" >> $(call formatTestLog,$*)
	@go test -v $* | tee -a $(logDir)/$$(basename $*)_tests.log

docker-start: ## Start the basic Golang app with Docker compose

docker-build: ## Build the current Dockerfile with Docker build

docker-stop: ## Stop the basic Golang app with Docker compose

binFiles = $(addprefix item-, $(wildcard $(binDir)/*))
logFiles = $(addprefix item-, $(wildcard $(logDir)/*))

$(logFiles): item-%:
	@echo "Removing '$*'"
	@rm -f $*

$(binFiles): item-%:
	@echo "Removing '$*'"
	@rm -f $*

clean: $(binFiles) $(logFiles) ## Remove all the build files
	go clean
	@if [ -d "$(binDir)" ]; then \
		echo "Removing '$(binDir)'"; \
		rm -r $(binDir); \
	fi
	@if [ -d "$(logDir)" ]; then \
		echo "Removing '$(logDir)'"; \
		rm -r $(logDir); \
	fi

help: ## Provides a help menu
	@echo "Application: $(PROGRAM_NAME)"
	@for i in $(CHARACTER_LEN_SET); do printf "%s" "="; done
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'