# Exercise 1.08: Project v0.5

Switch to using Ingress instead of NodePort to access the project. You can delete the ingress of the "Log output" application so they don't interfere with this exercise. We'll look more into paths and routing in the next exercise and at that point you can configure project to run with the "Log output" application side by side.

## Requirements

- [Golang](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/engine/install/)
- [Kubectl](https://kubernetes.io/docs/reference/kubectl/)
- [k3d](https://github.com/rancher/k3d#get)

## Usage
1. `./build.sh`
2. `kubectl apply -f manifests/deployment.yaml`
3. `kubectl apply -f manifests/service.yaml`
4. `kubectl apply -f manifests/ingress.yaml`