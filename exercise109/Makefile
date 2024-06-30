.PHONY: all
all: build tag push configure deploy

.PHONY: build
build:
	docker build -t pingpong:latest .

.PHONY: tag
tag:
	$(eval DIGEST := $(shell docker images --no-trunc --quiet pingpong:latest | cut -d':' -f2))
	echo "$(DIGEST)" > digest.txt
	docker image tag pingpong:latest pingpong:$(DIGEST)

.PHONY: push
push:
	$(eval DIGEST := $(shell cat digest.txt))
	k3d image import pingpong:$(DIGEST)

.PHONY: configure
configure:
	$(eval DIGEST := $(shell docker images --no-trunc --quiet pingpong:latest | cut -d':' -f2))
	sed "s/latest/$(DIGEST)/g" manifests/deployment_template.yaml > manifests/deployment.yaml

.PHONY: deploy
deploy:
	kubectl apply -f manifests/deployment.yaml
	kubectl apply -f manifests/service.yaml
	kubectl apply -f manifests/ingress.yaml