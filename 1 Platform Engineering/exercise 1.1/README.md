# Exercise 1.1: Working with Docker and Dockerfile

In this exercise, you will write a Dockerfile and you will build an image from this Dockerfile. For more information about a Dockerfile, click [here](https://docs.docker.com/engine/reference/builder/).

## Requirements

* [Docker](https://www.docker.com/) installed: 
    <details><summary>Click here for install instructions.</summary>
    <p>

    * Docker for Windows: https://docs.docker.com/docker-for-windows/install/

    * Docker for Mac: https://docs.docker.com/docker-for-mac/install

    </p>
    </details>

* [Docker Hub](https://hub.docker.com/) Account created:
    <details><summary>Click here for sign-up instructions.</summary>
    <p>

    * To sign up: https://hub.docker.com/signup

    </p>
    </details>

## Instructions

1. Create a Dockerfile with the following instructions:
    * Base image: `golang:1.13-alpine`
    * Set author label: `author='<YOUR-EMAIL'>`
    * Set working directory: `/opt`
    * Copy local file `main.go` to the image folder `/opt/`
    * List items in the working directory (`ls -lsa`) and show content of the `main.go` file (cat)
    * Build the app (`CGO_ENABLED=0 go build -o /usr/myapp`) and show the /usr directory  (`ls -lsa`)

1. Build a Docker image based on the Dockerfile:
    * Image tag: `[YOUR-DOCKERHUB-ACCOUNT]/my-first-image:0.0.1`

    ```console
    docker image build -f Dockerfile -t [YOUR-DOCKERHUB-ACCOUNT]/my-first-image:0.0.1 ./ 
    ```
    *Note:* If your Docker version is greater than 0.18, you are using Docker's new BuildKit engine that does not display the output from a `RUN` step. Therefore, add the option: `--progress=plain`

1. Execute the command a second time: 

    ```console
    docker image build -f Dockerfile -t [YOUR-DOCKERHUB-ACCOUNT]/my-first-image:0.0.1 ./
    ```

    :mag: What is your observation? 
  
1. List all images that are stored in your local registry:

    ```console
    docker images
    ```
