# Exercise 4.1: Performing a Rolling Update

In this exercise, you will learn how to roll out a new version of our application using
a **rolling update** strategy. Rolling updates allow updates of Deployments to take place with zero downtime by incrementally updating Pods instances with new ones.

To get a zero-downtime deployment (i.e. without having an availability gap in our deployment), we will add a readiness probe as well as a lifecycle hook:
- The **readiness probe** checks whether our application is ready to handle the traffic.
- The lifecycle hook **preStop** allows to add a waiting time before the final termination signal is sent to our container. This allows Kubernetes to remove the Pod from the endpoints and,
by this, the LoadBalancer is reconfigured before the application process is terminated. With this grace period, we can ensure no requests are routed to a terminated pod.

## Instructions

1. In order to prepare your deployment for a zero-downtime rolling update, you need to specify the following fields:
    - `.spec.replicas`: Specifies the used replicas. Here, at least two replicas are needed to perform a rolling update, e.g. use 5 replicas.
    - `.spec.strategy.type`: Specifies the strategy used to replace old Pods by new ones. `.spec.strategy.type` can be "Recreate" or "RollingUpdate".
    - `.spec.strategy.rollingUpdate.maxUnavailable`: Specifies the maximum number of Pods that can be unavailable during the update process.
    - `.spec.strategy.rollingUpdate.maxSurge`: Specifies the maximum number of Pods that can be created over the desired number of Pods.
    - `.spec.container.readinessProbe`: Specifies a readiness probe for checking whether our application is ready to handle the traffic.
    - `.spec.container.lifecycle.preStop`: Specifies a lifecycle hook, which must be completed before the container is terminated. 


    ```yaml
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: demo
    spec:
      replicas: 5
      selector:
        matchLabels:
          app: demo
      strategy:
        type: RollingUpdate
        rollingUpdate:
          maxSurge: 1
          maxUnavailable: 1
      template:
        metadata:
          labels:
            app: demo
        spec:
          containers:
          - name: demo
            image: agrimmer/demo:plain
            ports:
            - containerPort: 8888
            readinessProbe:
              httpGet:
                path: /
                port: 8888
              initialDelaySeconds: 5
            lifecycle:
              preStop:
                exec:
                  command: ["sleep", "15"]
    ```

1. Watch the updating of the Pods in a terminal:

    ```console
    watch kubectl get pods
    ```

1. Open a second terminal and poll your application in a loop:

    ```console
    while true; do sleep 1; curl http://HOSTNAME; echo; done
    ```

1. Update the image of your deployment:

    ```console
    kubectl set image deployments/demo demo=agrimmer/demo:html
    ```

:rocket: Watch the update of your application without experiencing any downtime.
