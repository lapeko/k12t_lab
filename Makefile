.PHONY: recreate secrets

recreate:
	docker stop $$(docker ps -q) || true
	cd ./infra/kind && make delete_clusters && make create_clusters
	cd ./infra/ansible && make install-cilium && make install-argocd
	make secrets

secrets:
	sudo fuser -k 30443/tcp || true
	sudo fuser -k 30444/tcp || true
	kubectx kind-dev
	kubectl apply -f ./argocd/tools/dev-root-tools.yaml
	kubectl -n argocd port-forward svc/argocd-server 30444:443 &
	@echo
	@echo ====================================================================================================
	@echo "\tDEV"
	@echo "address:  https://localhost:30444"
	@echo "password: $$(kubectl -n argocd get secret argocd-initial-admin-secret -o json | jq .data.password -r | base64 -d)"
	@echo ====================================================================================================
	@echo
	kubectx kind-prod
	kubectl apply -f ./argocd/tools/prod-root-tools.yaml
	sleep 5
	kubectl -n argocd port-forward service/argocd-server 30443:443 &
	@echo
	@echo ====================================================================================================
	@echo "\tPROD"
	@echo "address: https://localhost:30443"
	@echo "password: $$(kubectl -n argocd get secret argocd-initial-admin-secret -o json | jq .data.password -r | base64 -d)"
	@echo ====================================================================================================
	@echo
# 	sleep 
# 	kubectx kind-dev
# 	kubectl create secret generic postgres --from-literal=username=postgres --from-literal=password='dev-password-123' --dry-run=client -o yaml | kubeseal --scope cluster-wide --format yaml > ./apps/subscriptions/k8s/base/db/sealed-secret.yaml
# 	kubectx kind-prod
# 	kubectl create secret generic postgres --from-literal=username=postgres --from-literal=password='~prod-passw0rd-123!!@' --dry-run=client -o yaml | kubeseal --scope cluster-wide --format yaml > ./apps/subscriptions/k8s/overlays/prod/db/sealed-secret.yaml
# 	kubectl create secret generic postgres --from-literal=username=postgres --from-literal=password='stage-passw0rd-12345' --dry-run=client -o yaml | kubeseal --scope cluster-wide --format yaml > ./apps/subscriptions/k8s/overlays/stage/db/sealed-secret.yaml