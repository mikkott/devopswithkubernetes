# Exercise 1.12: Project v0.6

 The goal is to add an hourly image to the project.

Get a random picture from Lorem Picsum like https://picsum.photos/1200 and display it in the project. Find a way to store the image so it stays the same for 60 minutes.

Make sure to cache the image into a volume so that the API isn't needed for new images every time we access the application or the container crashes.

## Requirements

- [Golang](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/engine/install/)
- [Kubectl](https://kubernetes.io/docs/reference/kubectl/)
- [k3d](https://github.com/rancher/k3d#get)

## Usage
1. `make all`