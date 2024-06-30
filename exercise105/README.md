# Exercise 1.05: Project v0.3

Have the project respond something to a GET request sent to the project. A simple html page is good or you can deploy something more complex like a single-page-application.

See here how you can define environment variables for containers.

Use kubectl port-forward to confirm that the project is accessible and works in the cluster by using a browser to access the project.

## Requirements

[Golang](https://go.dev/doc/install)
[Docker](https://docs.docker.com/engine/install/)
[Kubectl](https://kubernetes.io/docs/reference/kubectl/)
[k3d](https://github.com/rancher/k3d#get)

## Usage
`./build.sh`
`kubectl apply -f manifests/deployment.yaml`