.PHONY: start_cluster stop_cluster create_cluster delete_cluster

CLUSTER1 = cluster1
CLUSTER2 = cluster2
LABEL = io.k8s.sigs.kind.cluster

start:
	docker start $(docker ps -a -q -f label=$(LABEL))

stop:
	docker stop $(docker ps -q -f label=$(LABEL))

create:
	kind create cluster --name $(CLUSTER1) --config cluster.yaml
	kind create cluster --name $(CLUSTER2) --config cluster.yaml

delete:
	kind delete cluster --name $(CLUSTER1)
	kind delete cluster --name $(CLUSTER2)
