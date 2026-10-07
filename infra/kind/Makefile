.PHONY: start_clusters stop_clusters create_clusters delete_clusters

CLUSTER1 = prod
CLUSTER2 = dev
LABEL = io.k8s.sigs.kind.cluster

start_clusters:
	docker start $(docker ps -a -q -f label=$(LABEL))

stop_clusters:
	docker stop $(docker ps -q -f label=$(LABEL))

create_clusters:
	kind create cluster --name $(CLUSTER1) --config cluster.yaml
	kind create cluster --name $(CLUSTER2) --config cluster.yaml

delete_clusters:
	kind delete cluster --name $(CLUSTER1)
	kind delete cluster --name $(CLUSTER2)
