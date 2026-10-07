.PHONY: install-cilium install-argocd

install-cilium:
	ansible-playbook -i inventory.ini ./playbooks/install-cilium.yaml
	
install-argocd:
	ansible-playbook -i inventory.ini ./playbooks/install-argocd.yaml