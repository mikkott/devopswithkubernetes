.PHONY: all
all: build tag push configure deploy

.PHONY: build
build: build-writer build-reader

.PHONY: tag
tag: tag-writer tag-reader

.PHONY: push
push: push-writer push-reader

.PHONY: build-writer
build-writer:
	docker build -t logoutputwriter:latest -f writer/Dockerfile ./writer/

.PHONY: tag-writer
tag-writer:
	$(eval DIGEST := $(shell docker images --no-trunc --quiet logoutputwriter:latest | cut -d':' -f2))
	echo "$(DIGEST)" > writer/digest.txt
	docker image tag logoutputwriter:latest logoutputwriter:$(DIGEST)

.PHONY: push-writer
push-writer:
	$(eval DIGEST := $(shell cat writer/digest.txt))
	k3d image import logoutputwriter:$(DIGEST)

.PHONY: build-reader
build-reader:
	docker build -t logoutputreader:latest -f reader/Dockerfile ./reader/

.PHONY: tag-reader
tag-reader:
	$(eval DIGEST := $(shell docker images --no-trunc --quiet logoutputreader:latest | cut -d':' -f2))
	echo "$(DIGEST)" > reader/digest.txt
	docker image tag logoutputreader:latest logoutputreader:$(DIGEST)

.PHONY: push-reader
push-reader:
	$(eval DIGEST := $(shell cat reader/digest.txt))
	k3d image import logoutputreader:$(DIGEST)

.PHONY: configure
configure:
	$(eval DIGEST_WRITER := $(shell docker images --no-trunc --quiet logoutputwriter:latest | cut -d':' -f2))
	$(eval DIGEST_READER := $(shell docker images --no-trunc --quiet logoutputreader:latest | cut -d':' -f2))
	sed -e "s/logoutputwriter:latest/logoutputwriter:$(DIGEST_WRITER)/g" -e "s/logoutputreader:latest/logoutputreader:$(DIGEST_READER)/g" manifests/deployment_template.yaml > manifests/deployment.yaml

.PHONY: deploy
deploy:
	kubectl apply -f manifests/deployment.yaml
	kubectl apply -f manifests/service.yaml
	kubectl apply -f manifests/ingress.yaml