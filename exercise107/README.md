# Exercise 1.07: External access with Ingress

"Log output" application currently outputs a timestamp and a random string to the logs.

Add an endpoint to request the current status (timestamp and string) and an ingress so that you can access it with a browser.

You can just store the string and timestamp to the memory.

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