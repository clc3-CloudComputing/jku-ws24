# Exercise 3.2: Running a Demo App

In this exercise, you will deploy a container in a Kubernetes cluster. Thus, a Pod will be created that runs the container. To access the running container, a port forward to a port on your local machine will be conducted. 

## Requirements

* `kubectl` is connected to the Kubernetes cluster created in the previous [exercise 3.1](../exercise%203.1)

## Instructions

1. In Kubernetes, the equivalent to `docker container run` is `kubectl run`. Use this command to run your container:

    ```console
    kubectl run demo --image=agrimmer/demo:time
    ```

1. To verify that the container has started and the app is running, use:
    
    ```console
    kubectl get pods
    ```

1. To forward your local port **9999** to the container port **8888**, use:
    
    ```console
    kubectl port-forward pod/demo 9999:8888
    ```

    :mag: Now, access the website [http://localhost:9999](http://localhost:9999). What is shown there? 

1. To delete the pod, use:

    ```console
    kubectl delete pod demo
    ```
