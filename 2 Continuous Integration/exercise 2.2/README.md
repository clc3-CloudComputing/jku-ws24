# Exercise 2.2: Build and push a Container Image to DockerHub

In this exercise, you will integrate DockerHub into the CI workflow to push an image to your DockerHub account. Therefore, a Dockerfile needs to be added to the repository, and the steps to build, tag, and push an image need to be extended to the GitHub Action workflow. 

## Requirements

* DockerHub Account account
    <details><summary>Click here for sign-up instructions.</summary>
    <p>

    To sign up: https://hub.docker.com/signup

    </p>
    </details>

* Exercise 2.1 completed

## Instructions

1. Setup credentials for DockerHub in GitHub:

    1. In your repository, go to `Settings` > `Secrets and variables` > `Actions` > **New repository secret**

    1. Add the secret `DOCKERHUB_USERNAME` with your DockerHub registry user name

    1. Add the secret `DOCKERHUB_TOKEN` with your DockerHub registry password. 

1. Add and commit the provided `Dockerfile` to your repository.

1. Extend the `CI.yml` workflow with the following step. Note: This is *not required but recommended* using it to be able to build multi-platform images:

    ```yaml
    # Boot Docker builder using by default the docker-container
    - name: Set up Docker Buildx
      uses: docker/setup-buildx-action@v3
    ```


1. Extend the `CI.yml` workflow with a Docker `login` step:

    ```yaml
    # Login to DockerHub account
    - name: Login to DockerHub
      uses: docker/login-action@v3 
      with:
        username: ${{ secrets.DOCKERHUB_USERNAME }}
        password: ${{ secrets.DOCKERHUB_TOKEN }}
    ```
    
1. Extend the `CI.yml` workflow with a Docker `build-push` step. Please change the [user] to your DockerHub account name: 

    ```yaml
    # Build a Docker image based on the provided Dockerfile
    - name: Build and push
      id: docker_build
      uses: docker/build-push-action@v5
      with:
        context: .
        push: true
        tags: [user]/app:latest
    ```

1. Finally, trigger a build by a code change. 

1. Watch GitHub Actions executing your tests, building the artifact, and pushing it to DockerHub.

:mag: Go to your DockerHub account and find your image. Can you find it?


### Use Git SHA instead of `latest`

It is not good practice to use the version `latest` for an image tag. Instead, make it as concrete as possible, e.g., by using the Git commit SHA instead: 

1. Extend the `CI.yml` workflow with a step that derives the short Git commit SHA and stores it in a variable: 

    ```yaml
    # Declare variables to store branch name and short Git commit SHA
    - name: Declare variables
      id: vars
      shell: bash
      run: |
        echo "sha_short=$(git rev-parse --short HEAD)" >> $GITHUB_OUTPUT
    ```

1. Change the `Build and push` step to reference the `sha_short` variable for the tag version: 

    ```yaml
        tags: [user]/app:${{ steps.vars.outputs.sha_short }}
    ```

1. Finally, trigger a build by a code change. 
