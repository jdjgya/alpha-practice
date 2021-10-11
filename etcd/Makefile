.EXPORT_ALL_VARIABLES:
GIT_COMMIT ?= $(shell git rev-parse --short HEAD)
SERVICE_NAME ?= alphaetcd


.PHONY: prepare-etcd
prepare-etcd:
	@docker run -itd --name $(SERVICE_NAME) -p 2379:2379 -e ETCD_ADVERTISE_CLIENT_URLS="http://0.0.0.0:2379" -e ETCD_LISTEN_CLIENT_URLS="http://0.0.0.0:2379" quay.io/coreos/etcd:v3.4.16

.PHONY: clean-etcd
clean-etcd:
	@docker rm -f $(SERVICE_NAME)
