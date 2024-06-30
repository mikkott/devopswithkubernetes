# Exercise 1.10: Even more services

Split the "Log output" application into two different containers within a single pod:

One generates a new timestamp every 5 seconds and saves it into a file.

The other reads that file and outputs it with a hash for the user to see.

Either application can generate the hash. The reader or the writer. 

## Requirements

- [Golang](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/engine/install/)
- [Kubectl](https://kubernetes.io/docs/reference/kubectl/)
- [k3d](https://github.com/rancher/k3d#get)

## Usage
1. `make all`