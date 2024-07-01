.PHONY: all
all: build tag push configure deploy

.PHONY: build
build: build-logoutputreader build-pingpong

.PHONY: tag
tag:tag-logoutputreader tag-pingpong

.PHONY: push
push:push-logoutputreader push-pingpong

.PHONY: build-logoutputreader
build-logoutputreader:
	docker build -t logoutputreader:latest -f logoutputreader/Dockerfile ./logoutputreader/

.PHONY: tag-logoutputreader
tag-logoutputreader:
	$(eval DIGEST := $(shell docker images --no-trunc --quiet logoutputreader:latest | cut -d':' -f2))
	echo "$(DIGEST)" > logoutputreader/digest.txt
	docker image tag logoutputreader:latest logoutputreader:$(DIGEST)

.PHONY: push-logoutputreader
push-logoutputreader:
	$(eval DIGEST := $(shell cat logoutputreader/digest.txt))
	k3d image import logoutputreader:$(DIGEST)

.PHONY: build-pingpong
build-pingpong:
	docker build -t pingpong:latest -f pingpong/Dockerfile ./pingpong/

.PHONY: tag-pingpong
tag-pingpong:
	$(eval DIGEST := $(shell docker images --no-trunc --quiet pingpong:latest | cut -d':' -f2))
	echo "$(DIGEST)" > pingpong/digest.txt
	docker image tag pingpong:latest pingpong:$(DIGEST)

.PHONY: push-pingpong
push-pingpong:
	$(eval DIGEST := $(shell cat pingpong/digest.txt))
	k3d image import pingpong:$(DIGEST)

.PHONY: configure
configure:
	$(eval DIGEST_LOGOUTPUTREADER := $(shell docker images --no-trunc --quiet logoutputreader:latest | cut -d':' -f2))
	$(eval DIGEST_PINGPONG := $(shell docker images --no-trunc --quiet pingpong:latest | cut -d':' -f2))
	sed -e "s/logoutputreader:latest/logoutputreader:$(DIGEST_LOGOUTPUTREADER)/g" -e "s/pingpong:latest/pingpong:$(DIGEST_PINGPONG)/g" manifests/deployment_template.yaml > manifests/deployment.yaml

.PHONY: deploy
deploy:
	kubectl apply -f manifests/pvc/persistentvolume.yaml
	kubectl apply -f manifests/persistentvolumeclaim.yaml
	kubectl apply -f manifests/deployment.yaml
	kubectl apply -f manifests/service.yaml
	kubectl apply -f manifests/ingress.yaml