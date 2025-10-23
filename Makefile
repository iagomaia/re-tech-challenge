.PHONY: all build clean
.PHONY: all vendor clean

include .env
export

db-local:
	docker compose up -d mongo

ensure-dependencies:
	go mod tidy

test:
	go clean -testcache
	go test ./...

test-cover:
	mkdir -p cov-report
	go clean -testcache
	go test -coverprofile=./cov-report/coverage.out ./internal/services/usecases/*
	go tool cover -html=./cov-report/coverage.out
	
vendor: 
	go mod vendor

run-api: ensure-dependencies vendor
	clear
	go run cmd/api/main.go

build-docs:
	mkdir -p static
	cp ./docs/swagger.yaml ./static/swagger.yaml

build:
ifeq ($(OS),Windows_NT)
	if not exist build mkdir build
else
	mkdir -p build
endif
	go build -mod=readonly -v -o ./build ./...
	cp .env ./build/.env
	mkdir -p ./build/static
	cp ./docs/swagger.yaml ./build/static/swagger.yaml

docker-build: build
	VERSION_IMAGE=""
ifdef VERSION_IMAGE
	docker image build --platform=linux/arm64 . -t packaging-api:${VERSION_IMAGE}
else
	docker image build --platform=linux/arm64 . -t packaging-api
endif

docker-run:
	docker compose up -d --build api