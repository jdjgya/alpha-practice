.EXPORT_ALL_VARIABLES:
GIT_COMMIT ?= $(shell git rev-parse --short HEAD)
SERVICE_NAME ?= alphanats


.PHONY: prepare-nats
prepare-nats:
	@docker run -itd --name $(SERVICE_NAME) -p 4222:4222 nats:2.7.0 -js

.PHONY: clean-nats
clean-nats:
	@docker rm -f $(SERVICE_NAME)
