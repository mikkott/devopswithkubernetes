# Exercise 1.04: Project v0.2

Create a deployment.yaml for the project.

You won't have access to the port yet but that'll come soon.

## Requirements

[Golang](https://go.dev/doc/install)
[Docker](https://docs.docker.com/engine/install/)
[Kubectl](https://kubernetes.io/docs/reference/kubectl/)
[k3d](https://github.com/rancher/k3d#get)

## Usage
`./build.sh`
`kubectl apply -f manifests/deployment.yaml`