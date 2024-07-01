# Exercise 2.02: Project v1.0

Let us get back to our Project. In the previous part we added a random pic and a form for creating todos to the app. The next step is to create a new container that takes care of saving the todo items.

This new service, let us call it todo-backend, should have a GET /todos endpoint for fetching the list of todos and a POST /todos endpoint for creating a new todo. The todos can be saved into memory, we'll add a database later.

Use ingress routing to enable access to the todo-backend.

After this exercise, the project should look like the following:

![diagram](https://devopswithkubernetes.com/static/bc4bed9387ebafa11912ae48b2339d14/b3c31/p2-2.webp)

The role of the service that we made in previous exercises (Todo-app in the figure) is to serve the HTML and possibly JavaScript to the browser. Also, the logic for serving random pictures and caching those remain in that service.

The new service then takes care of the todo items.

After this exercise, you should be able to create new todos using the form, and the created todos should be rendered in the browser.

## Requirements

- [Golang](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/engine/install/)
- [Kubectl](https://kubernetes.io/docs/reference/kubectl/)
- [k3d](https://github.com/rancher/k3d#get)

## Usage
1. `make all`