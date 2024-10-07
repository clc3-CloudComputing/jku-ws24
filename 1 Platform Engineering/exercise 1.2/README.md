# Exercise 1.2: Pushing/Pulling a Docker image to/from a container registry

In this exercise, you will push an image to the container registry [DockerHub](https://hub.docker.com/). Afterwards, you are going to pull an image from your classmate to your local machine. 

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

1. Authenticate to the container registry:

    ```console
    docker login
    ```

    ```console
    Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
    Username: YOUR-DOCKERHUB-ACCOUNT
    Password: YOUR-DOCKERHUB-PASSWORD
    Login Succeeded
    ```

1. Push image from exercise 1.1 to your DockerHub account:

    ```console
    docker image push [YOUR-DOCKERHUB-ACCOUNT]/my-first-image:0.0.1
    ```

1. Verify the push on your account: https://hub.docker.com/

1. Pull the image from a classmate's DockerHub account:

    ```console
    docker image pull [YOUR-CLASSMATE'S-ACCOUNT]/my-first-image:0.0.1
    ```

1. Verify the pull on your local machine:

    ```console
    docker images
    ```
