.PHONY: start_cluster stop_cluster create_cluster delete_cluster

CLUSTER1 = cluster1
CLUSTER2 = cluster2
LABEL = io.k8s.sigs.kind.cluster

start_cluster:
	docker start $(docker ps -a -q -f label=$(LABEL))

stop_cluster:
	docker stop $(docker ps -q -f label=$(LABEL))

create_cluster:
	kind create cluster --name $(CLUSTER1) --config cluster.yaml
	kind create cluster --name $(CLUSTER2) --config cluster.yaml

delete_cluster:
	kind delete cluster --name $(CLUSTER1)
	kind delete cluster --name $(CLUSTER2)
