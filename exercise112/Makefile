.PHONY: all
all: build tag push configure deploy

.PHONY: build
build:
	docker build -t project:latest .

.PHONY: tag
tag:
	$(eval DIGEST := $(shell docker images --no-trunc --quiet project:latest | cut -d':' -f2))
	echo "$(DIGEST)" > digest.txt
	docker image tag project:latest project:$(DIGEST)

.PHONY: push
push:
	$(eval DIGEST := $(shell cat digest.txt))
	k3d image import project:$(DIGEST)

.PHONY: configure
configure:
	$(eval DIGEST_PROJECT := $(shell docker images --no-trunc --quiet project:latest | cut -d':' -f2))
	sed -e "s/project:latest/project:$(DIGEST_PROJECT)/g" manifests/deployment_template.yaml > manifests/deployment.yaml

.PHONY: deploy
deploy:
	kubectl apply -f manifests/pvc/persistentvolume.yaml
	kubectl apply -f manifests/persistentvolumeclaim.yaml
	kubectl apply -f manifests/deployment.yaml
	kubectl apply -f manifests/service.yaml
	kubectl apply -f manifests/ingress.yaml
