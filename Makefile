PROJECTNAME := $(shell basename "$(PWD)")
OS := $(shell uname -s | awk '{print tolower($$0)}')
GOARCH := amd64

## tidy: special go mod tidy without golang database checksum(GOSUMDB) 
.PHONY: tidy
tidy:
	export GOSUMDB=off ; go mod tidy

## test: run go test
.PHONY: test
test:
	go clean -testcache & go test -v -race ./...

## run-local-consumer: execute ct-consumer-service background on local.
.PHONY: run-local-consumer
run-local-consumer:
	go run --race app/cmd/consumer/main.go serve --env=local

## run-local-http: execute ct-http-service httpserver on local.
.PHONY: run-local-http
run-local-http:
	go run --race app/cmd/http/main.go serve --env=local


## up-container: run docker container with logs.
.PHONY: up-container
up-container:
	docker compose up

## up-container-d: run docker container with skip logs.
.PHONY: up-container-d
up-container-d:
	docker compose up -d 

## stop-container: stop docker container.
.PHONY: stop-container
stop-container:
	docker compose stop 

## down-container: remove docker container.
.PHONY: down-container
down-container:
	docker compose down

## set_private_repo_global: set a "" to be a private repo in go global environment 
# set_private_repo_global:
# 	go env -w GOPRIVATE=""

## update_standard_lib: update standard library ([lib-repo-name]) with GOPRIVATE option
# update_standard_lib:
# 	GOPRIVATE=


## help: helper
.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Project: ["$(PROJECTNAME)"]"
	@echo " Please choose a command"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## gosec: run for scan code vulnerability by securego/gosec
.PHONY: gosec 
gosec: 
	gosec ./... 

## govulncheck: run for scan vulnerability package from Go vulnerability database
.PHONY: govulncheck
govulncheck: 
	govulncheck ./...


## security: run make gosec and make govulncheck
security: gosec govulncheck

## sqlite: Initializing sqlite3 to create database file.
sqlite:
	./tools/bash-script/sqlite-init.sh

## mocks: generate mock test for all files.
mocks:
	./tools/bash-script/mocks.sh


