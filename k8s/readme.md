`https://kubernetes.io/docs/concepts/`

## Components
- master
  - scheduling
  - pod management
  - kube api
  - etcd
  - kube controller manager
  - cloud controller manager

- node
  - kubelet: agent
  - kube proxy
  - container runtime

- addons
  - DNS
  - Dashboard

## Objects

- Pod
- Service
- Volume
- Namespace

###

- ReplicaSet
  - `replicas: 3`
- Deployment
  - Updates for Pods and ReplicaSets
- SatefulSet
  - from 1.9
  - Stable, unique network identifiers.
  - Stable, persistent storage.
  - Ordered, graceful deployment and scaling.
  - Ordered, automated rolling updates.
- DaemonSet
  - Ensures all or some Nodes run a copy of a Pod
  - Those Pods are garbage collected while nodes are removed from the cluster
- Job


## Service

- Defines a logical set of Pods
- Access the Pods

## Pod
- Probes:
  - Liveness: `Indicates whether the Container is running`
  - Readiness: `Indicates whether the Container is ready to service requests`

  - Ex:
    ```
    livenessProbe:
      httpGet:
        path: /healthz
        port: 8080
        httpHeaders:
        - name: X-Custom-Header
          value: Awesome
      initialDelaySeconds: 3
      periodSeconds: 3
    ```