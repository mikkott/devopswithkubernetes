# Exercise 1.09: More services

Develop a second application that simply responds with "pong 0" to a GET request and increases a counter (the 0) so that you can see how many requests have been sent. The counter should be in memory so it may reset at some point. Create a new deployment for it and have it share ingress with "Log output" application. Route requests directed '/pingpong' to it.

In future exercises, this second application will be referred to as "ping-pong application". It will be used with "Log output" application.

The ping-pong application will need to listen requests on '/pingpong', so you may have to make changes to its code. 

## Requirements

- [Golang](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/engine/install/)
- [Kubectl](https://kubernetes.io/docs/reference/kubectl/)
1. - [k3d](https://github.com/rancher/k3d#get)1. 

## Usage
1. `make all`