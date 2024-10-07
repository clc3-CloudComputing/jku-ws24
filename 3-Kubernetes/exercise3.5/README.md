# Exercise 3.5: Creating a ConfigMap

In this exercise, you will create a ConfigMap, which allows us to control the output format of our application.
More precisely, the ConfigMap will control if the output is in plain text or HTML-formatted.
Therefore, you will mount the ConfigMap as an environment variable in the Pod.
By editing the ConfigMap, you will be able to control the output format without changing the image.

## Instructions

1. Create a file `configmap.yaml` with the following content:

    ```yaml
    apiVersion: v1
    kind: ConfigMap
    metadata:
      name: outputformat
    data:
      format: "html"
    ```

1. Apply the ConfigMap to your Kubernetes cluster using: 

    ```console
    kubectl apply -f configmap.yaml
    ```

1. Configure the container inside the Pod by adding an environment variable that refers to the ConfigMap:

    ```yaml
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: demo
    spec:
      replicas: 1
      selector:
        matchLabels:
          app: demo
      template:
        metadata:
          labels:
            app: demo
        spec:
          containers:
          - name: demo
            image: agrimmer/demo:formatted
            ports:
            - containerPort: 8888
            env:
              - name: "FORMAT"
                valueFrom:
                  configMapKeyRef:
                    name: outputformat
                    key: format
    ```

1. Apply the changes to your Kubernetes cluster using: 

    ```console
    kubectl apply -f deployment.yaml
    ```

1. To forward your local port **9999** to the container port **8888**, use:

    ```console
    kubectl port-forward deployment/demo 9999:8888
    ```

    :mag: Now, access the website [http://localhost:9999](http://localhost:9999). Do you see an HTML-formatted output? 

1. Change the output format back to *plain*. What is needed that the changes are applied?

1. Mount the ConfigMap as volume and check the files by getting a shell to the running container.
