# Exercise 3.4: Creating a Service

In this exercise, you will specify a Service for your Deployment.
Therefore, you will use the Service types *ClusterIP*, *NodePort*, and *LoadBalancer*.
The Service resource will be created by defining a service manifest (`service.yaml`) and afterwards 
applying it to your Kubernetes cluster.


## Instructions

### ClusterIP
1. Create a file `service.yaml` with the following content:

    ```yaml
    apiVersion: v1
    kind: Service
    metadata:
      name: demo
      labels:
        app: demo
    spec:
      ports:
      - port: 80
        protocol: TCP
        targetPort: 8888
      selector:
        app: demo
      type: ClusterIP
    ```

    * The `ports`-section specifies that the Service forwards all traffic received on port 80 to the Pod's port on 8888.

    * The `selector` tells the Service to which Pods the requests should be routed. Here, the requests are forwarded to any Pods matching the label `app: demo`.

1. Apply the *service* to your Kubernetes cluster using: 
    
    ```console
    kubectl apply -f service.yaml
    ```

1. Your service now has a cluster-internal IP, which is not exposed to the public:

    ```console
    kubectl describe service demo
    ```

    ```source
    Name:              demo
    Namespace:         default
    Labels:            app=demo
    Annotations:       Selector:  app=demo
    Type:              ClusterIP
    IP:                10.0.29.19
    Port:              <unset>  80/TCP
    TargetPort:        8888/TCP
    Endpoints:         10.44.2.17:8888
    Session Affinity:  None
    Events:            <none>
    ```

1. Setup a port-forward with:

    ```console
    kubectl port-forward service/demo 8080:80
    ```

1. Use the http://localhost:8080 to access your website.

    :mag: Enter http://localhost:8080 in a browser.

### NodePort (only for GKE and minikube)
**Note for `kind` users:** This service type would require a [dedicated `kind` configuration](https://kind.sigs.k8s.io/docs/user/configuration/#extra-port-mappings).

**Note for `AKS` users:** In AKS, the node does not get a public IP address.

1. Change the service type to *NodePort* with:

    ```console
    kubectl patch service demo -p '{"spec": {"type": "NodePort"}}'
    ```

1. Your service is now exposed on each node's IP at a static port:

    ```console
    kubectl describe service demo
    ```

    ```source
    Name:                     demo
    Namespace:                default
    Labels:                   app=demo
    Annotations:              Selector:  app=demo
    Type:                     NodePort
    IP:                       10.0.29.19
    Port:                     <unset>  80/TCP
    TargetPort:               8888/TCP
    NodePort:                 <unset>  31425/TCP
    Endpoints:                10.44.2.17:8888
    Session Affinity:         None
    External Traffic Policy:  Cluster
    Events:                   <none>
    ```

1. Get an IP of any of your nodes:

    ```console
    kubectl get nodes -owide
    ```

1. Use the IP address of any of your nodes and the assigned port to access your website.

    :mag: Enter the IP:port in a browser.

    **Note for Google Cloud:**
    You need to create a firewall rule to allow TCP traffic on the node port using
    `gcloud compute firewall-rules create test-node-port --allow tcp:31425 --target-tags gke-grimmer-dev-6cb41d0c-node`

### LoadBalancer
**Note for `kind` users:** This service type is not supported.

**Note for `minikube` users:** A dedicated tunnel is needed (i.e. execute `minikube tunnel`).

1. Change the service type to *LoadBalancer* with:

    ```console
    kubectl patch service demo -p '{"spec": {"type": "LoadBalancer"}}'
    ```

1. Your service now acts as *LoadBalancer* and received an externally accessible IP address. You can obtain this IP with:

    ```console
    kubectl describe service demo
    ```

    ```source
    Name:                     demo
    Namespace:                default
    Labels:                   app=demo
    Annotations:              kubectl.kubernetes.io/last-applied-configuration:
                                {"apiVersion":"v1","kind":"Service","metadata":{"annotations":{},"labels":{"app":"demo"},"name":"demo","namespace":"default"},"spec":{"por...
    Selector:                 app=demo
    Type:                     LoadBalancer
    IP:                       10.0.87.134
    LoadBalancer Ingress:     20.185.79.32
    Port:                     <unset>  9999/TCP
    TargetPort:               8888/TCP
    NodePort:                 <unset>  31343/TCP
    Endpoints:                10.244.1.4:8888
    Session Affinity:         None
    External Traffic Policy:  Cluster
    Events:
      Type    Reason                Age   From                Message
      ----    ------                ----  ----                -------
      Normal  EnsuringLoadBalancer  15s   service-controller  Ensuring load balancer
      Normal  EnsuredLoadBalancer   11s   service-controller  Ensured load balancer
    ```

1. Use the IP address of the *LoadBalancer Ingress* to access your website.

    :mag: Enter the IP in a browser.
