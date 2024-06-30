# Exercise 1.06: Project v0.4

Use a NodePort Service to enable access to the project.

## Requirements

- [Golang](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/engine/install/)
- [Kubectl](https://kubernetes.io/docs/reference/kubectl/)
- [k3d](https://github.com/rancher/k3d#get)

## Usage
1. `./build.sh`
2. `kubectl apply -f manifests/deployment.yaml`
3. `kubectl apply -f manifests/service.yaml`