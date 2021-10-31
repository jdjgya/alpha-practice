.EXPORT_ALL_VARIABLES:
GIT_COMMIT ?= $(shell git rev-parse --short HEAD)
SERVICE_NAME ?= alpharedis


.PHONY: prepare-redis
prepare-redis:
	@docker run -itd --name $(SERVICE_NAME) --hostname redis -p 6379:6379 redis:6.2.6

.PHONY: clean-redis
clean-redis:
	@docker rm -f $(SERVICE_NAME)
