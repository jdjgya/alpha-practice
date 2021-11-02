.EXPORT_ALL_VARIABLES:
GIT_COMMIT ?= $(shell git rev-parse --short HEAD)
SERVICE_NAME ?= alpharedis


.PHONY: prepare-redis
prepare-redis:
	@docker run -itd --name $(SERVICE_NAME) --hostname redis -p 6379:6379 redis:6.2.6

.PHONY: clean-redis
clean-redis:
	@docker rm -f $(SERVICE_NAME)

.PHONY: prepare-redis-cluster
prepare-redis-cluster:
	@kubectl apply -f deploy/configmap.yaml
	@kubectl apply -f deploy/svc-headless.yaml
	@kubectl apply -f deploy/svc.clusterip.yaml
	@kubectl apply -f deploy/statefulset.yaml

.PHONY: clean-redis-cluster
clean-redis-cluster:
	@kubectl delete -f deploy/configmap.yaml
	@kubectl delete -f deploy/svc-headless.yaml
	@kubectl delete -f deploy/svc-clusterip.yaml
	@kubectl delete -f deploy/statefulset.yaml

.PHONY: init-redis-cluster
init-redis-cluster:
	$(eval redis_pod_ips := $(shell kubectl get pods -l app=redis-cacher -o jsonpath='{range.items[*]}{.status.podIP}:6397 {end}'))
	@kubectl exec -it redis-cacher-0 -- redis-cli --cluster create --cluster-replicas 1 $(redis_pod_ips)
