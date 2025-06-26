GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/rakyll/hey@latest

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api/grpc \
		   --go-http_out=paths=source_relative:./api/grpc \
 	       --go-grpc_out=paths=source_relative:./api/grpc \
 	       --experimental_allow_proto3_optional \
	       --openapi_out=fq_schema_naming=true,default_response=false,enum_type=string:./embeds \
		   --validate_out=paths=source_relative,lang=go:./api/grpc \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go generate ./...
	go mod tidy

.PHONY: all
# generate all
all:
	make api;
	make config;
	make generate;

.PHONY: api-image
# build docker image for the API
api-image:
	docker build -f platform/docker/Dockerfile -t thmanyah-api:latest .

.PHONY: up
# start all services using docker compose
up:
	docker compose -p thmanyah -f platform/docker/docker-compose.yaml up -d

.PHONY: down
# stop and remove all services and volumes
down:
	docker compose -p thmanyah -f platform/docker/docker-compose.yaml down -v

.PHONY: start
# build image and start all services
start: api-image up

.PHONY: stop
# stop all services without removing volumes
stop:
	docker compose -p thmanyah -f platform/docker/docker-compose.yaml down

.PHONY: restart
# restart all services with fresh image
restart: down api-image up

.PHONY: load-test
# run load testing against the discover search endpoint
search-load-test:
	hey -n 300000 -q 1000 -c 100 -m POST -H "Content-Type: application/json" -d '{"query": "AI"}' http://localhost:8000/api/v1/discover/search

featured-load-test:
	hey -n 300000 -q 1000 -c 100 http://localhost:8000/api/v1/discover/featured

.PHONY: debug-logs
# show real-time logs from all services
debug-logs:
	docker compose -p thmanyah -f platform/docker/docker-compose.yaml logs -f

.PHONY: debug-api-logs
# show real-time logs from API service only
debug-api-logs:
	docker compose -p thmanyah -f platform/docker/docker-compose.yaml logs -f api

.PHONY: debug-db-logs
# show real-time logs from database service
debug-db-logs:
	docker compose -p thmanyah -f platform/docker/docker-compose.yaml logs -f postgres

.PHONY: debug-shell
# open interactive shell in the API container
debug-shell:
	docker compose -p thmanyah -f platform/docker/docker-compose.yaml exec api sh

.PHONY: debug-db-shell
# open PostgreSQL shell for database debugging
debug-db-shell:
	docker compose -p thmanyah -f platform/docker/docker-compose.yaml exec postgres psql -U postgres -d thmanyah

.PHONY: debug-ps
# show status of all containers
debug-ps:
	docker compose -p thmanyah -f platform/docker/docker-compose.yaml ps

.PHONY: debug-minio
# open MinIO console at http://localhost:9001
debug-minio:
	@echo "MinIO Console available at: http://localhost:9001"
	@echo "Credentials placed at platform/docker/docker-compose.yaml"

.PHONY: debug-reset
# reset development environment (rebuild and restart)
debug-reset: down
	docker system prune -f
	make start

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
