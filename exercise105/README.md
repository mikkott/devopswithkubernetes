# Project

Builds and deploys Golang http server.

Automated build steps in `build.sh` to build minimal sized golang binary in Docker image, import Docker image to k3d and configure k8s manifest template.

To deploy in k8s: `kubectl apply -f manifests/deployment.yaml`