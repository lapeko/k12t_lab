.PHONY: all recreate-clusters show-argo-connections recreate-sealed-secrets run-root-apps

all:
	make recreate-clusters
	sleep 75
	make show-argo-connections
	sleep 75
	recreate-sealed-secrets

recreate-clusters:
	docker stop $$(docker ps -q) || true
	cd ./infra/kind && make delete_clusters && make create_clusters
	cd ./infra/ansible && make install-cilium && make install-argocd
	kubectx kind-dev
	kubectl apply -f ./argocd/tools/dev-root-tools.yaml
	kubectl -n argocd patch svc argocd-server --type='json' -p='[{"op": "replace", "path": "/spec/type", "value": "NodePort"}, {"op": "replace", "path": "/spec/ports/1/nodePort", "value": 30443}]'
	kubectx kind-prod
	kubectl apply -f ./argocd/tools/prod-root-tools.yaml
	kubectl -n argocd patch svc argocd-server --type='json' -p='[{"op": "replace", "path": "/spec/type", "value": "NodePort"}, {"op": "replace", "path": "/spec/ports/1/nodePort", "value": 30443}]'

show-argo-connections:
	kubectx kind-dev
	@echo
	@echo ====================================================================================================
	@echo "\tDEV"
	@echo "address: https://$$(docker inspect dev-control-plane | jq .[0].NetworkSettings.Networks.kind.IPAddress -r):30443"
	@echo "password: $$(kubectl -n argocd get secret argocd-initial-admin-secret -o json | jq .data.password -r | base64 -d)"
	@echo ====================================================================================================
	@echo
	kubectx kind-prod
	@echo
	@echo ====================================================================================================
	@echo "\tPROD"
	@echo "address: https://$$(docker inspect prod-control-plane | jq .[0].NetworkSettings.Networks.kind.IPAddress -r):30443"
	@echo "password: $$(kubectl -n argocd get secret argocd-initial-admin-secret -o json | jq .data.password -r | base64 -d)"
	@echo ====================================================================================================
	@echo

recreate-sealed-secrets:
	kubectx kind-dev
	kubectl create secret generic postgres --from-literal=username=postgres --from-literal=password='dev-password-123' --dry-run=client -o yaml | kubeseal --scope cluster-wide --format yaml > ./apps/subscriptions/k8s/base/db/sealed-secret.yaml
	kubectx kind-prod
	kubectl create secret generic postgres --from-literal=username=postgres --from-literal=password='~prod-passw0rd-123!!@' --dry-run=client -o yaml | kubeseal --scope cluster-wide --format yaml > ./apps/subscriptions/k8s/overlays/prod/db/sealed-secret.yaml
	kubectl create secret generic postgres --from-literal=username=postgres --from-literal=password='stage-passw0rd-12345' --dry-run=client -o yaml | kubeseal --scope cluster-wide --format yaml > ./apps/subscriptions/k8s/overlays/stage/db/sealed-secret.yaml
	git pull
	git add .
	git commit -m "refresh sealed secrets [skip ci]"
	git push

run-root-apps:
	kubectx kind-dev
	kubectl apply -f ./argocd/apps/dev-root-app.yaml
	kubectx kind-prod
	kubectl apply -f ./argocd/apps/prod-root-app.yaml