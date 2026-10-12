.PHONY: recreate secrets

recreate:
	docker stop $$(docker ps -q) || true
	cd ./infra/kind && make delete_clusters && make create_clusters
	cd ./infra/ansible && make install-cilium && make install-argocd
	make secrets

secrets:
	kubectx kind-dev
	kubectl create secret generic postgres --from-literal=username=postgres --from-literal=password='dev-password-123' --dry-run=client -o yaml | kubeseal --scope cluster-wide --format yaml > ./apps/subscriptions/k8s/base/db/sealed-secret.yaml
	kubectx kind-prod
	kubectl create secret generic postgres --from-literal=username=postgres --from-literal=password='~prod-passw0rd-123!!@' --dry-run=client -o yaml | kubeseal --scope cluster-wide --format yaml > ./apps/subscriptions/k8s/overlays/prod/db/sealed-secret.yaml
	kubectl create secret generic postgres --from-literal=username=postgres --from-literal=password='stage-passw0rd-12345' --dry-run=client -o yaml | kubeseal --scope cluster-wide --format yaml > ./apps/subscriptions/k8s/overlays/stage/db/sealed-secret.yaml