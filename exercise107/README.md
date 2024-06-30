# Exercise 1.07: External access with Ingress

"Log output" application currently outputs a timestamp and a random string to the logs.

Add an endpoint to request the current status (timestamp and string) and an ingress so that you can access it with a browser.

You can just store the string and timestamp to the memory.

## Requirements

[Golang](https://go.dev/doc/install)
[Docker](https://docs.docker.com/engine/install/)
[Kubectl](https://kubernetes.io/docs/reference/kubectl/)
[k3d](https://github.com/rancher/k3d#get)

## Usage
`./build.sh`
`kubectl apply -f manifests/deployment.yaml`
`kubectl apply -f manifests/service.yaml`
`kubectl apply -f manifests/ingress.yaml`