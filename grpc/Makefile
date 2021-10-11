.EXPORT_ALL_VARIABLES:
GIT_COMMIT ?= $(shell git rev-parse --short HEAD)
SERVICE_NAME ?= alphagrpc


.PHONY: build-proto
build-proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative example/example.proto
