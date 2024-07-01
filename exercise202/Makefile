.PHONY: all
all: build-app tag-app push-app build-backend tag-backend push-backend configure deploy

.PHONY: build-app
build-app:
	docker build -t todo-app:latest -f todo-app/Dockerfile ./todo-app/

.PHONY: tag-app
tag-app:
	$(eval DIGEST := $(shell docker images --no-trunc --quiet todo-app:latest | cut -d':' -f2))
	echo "$(DIGEST)" > todo-app/digest.txt
	docker image tag todo-app:latest todo-app:$(DIGEST)

.PHONY: push-app
push-app:
	$(eval DIGEST := $(shell cat todo-app/digest.txt))
	k3d image import todo-app:$(DIGEST)

.PHONY: build-backend
build-backend:
	docker build -t todo-backend:latest -f todo-backend/Dockerfile ./todo-backend/

.PHONY: tag-backend
tag-backend:
	$(eval DIGEST := $(shell docker images --no-trunc --quiet todo-backend:latest | cut -d':' -f2))
	echo "$(DIGEST)" > todo-backend/digest.txt
	docker image tag todo-backend:latest todo-backend:$(DIGEST)

.PHONY: push-backend
push-backend:
	$(eval DIGEST := $(shell cat todo-backend/digest.txt))
	k3d image import todo-backend:$(DIGEST)

.PHONY: configure
configure:
	$(eval DIGEST_APP := $(shell docker images --no-trunc --quiet todo-app:latest | cut -d':' -f2))
	$(eval DIGEST_BACKEND := $(shell docker images --no-trunc --quiet todo-backend:latest | cut -d':' -f2))
	sed -e "s/todo-app:latest/todo-app:$(DIGEST_APP)/g" -e "s/todo-backend:latest/todo-backend:$(DIGEST_BACKEND)/g" manifests/deployment_template.yaml > manifests/deployment.yaml

.PHONY: deploy
deploy:
	kubectl apply -f manifests/pvc/persistentvolume.yaml
	kubectl apply -f manifests/persistentvolumeclaim.yaml
	kubectl apply -f manifests/deployment.yaml
	kubectl apply -f manifests/service.yaml
	kubectl apply -f manifests/ingress.yaml
