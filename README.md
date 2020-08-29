# k8s-operator

Kubernetes Operator - Golang

## How to run on Mac

### Install with brew
```bash
brew install operator-sdk
```

### Confirm version

```bash
operator-sdk version

operator-sdk version: "v1.0.0", commit: "d7d5e0cd6cf5468bb66e0849f08fda5bf557f4fa", kubernetes version: "v1.18.2", go version: "go1.14.7 darwin/amd64", GOOS: "darwin", GOARCH: "amd64"
```

### Create and Initialize the Project

```bash
$ mkdir ~/k8s-operator && cd ~/k8s-operator/

$ operator-sdk init --plugins=ansible --domain example.com
```

### Create the API

```bash
$ operator-sdk create api --group cache --version v1alpha1 --kind Memcached --generate-role
```

### Modify the Manager

Update the file `roles/memcached/tasks/main.yml:`

```yaml
---
# tasks file for Memcached
- name: start memcached
  community.kubernetes.k8s:
    definition:
      kind: Deployment
      apiVersion: apps/v1
      metadata:
        name: '{{ ansible_operator_meta.name }}-memcached'
        namespace: '{{ ansible_operator_meta.namespace }}'
      spec:
        replicas: "{{ size }}"
        selector:
          matchLabels:
            app: memcached
        template:
          metadata:
            labels:
              app: memcached
          spec:
            containers:
            - name: memcached
              command:
              - memcached
              - -m=64
              - -o
              - modern
              - -v
              image: "docker.io/memcached:1.4.36-alpine"
              ports:
                - containerPort: 11211

```

Add the defaults to `roles/memcached/defaults/main.yml`

```yaml
---
# defaults file for Memcached
size: 1
```

And update the file `config/samples/cache_v1alpha1_memcached.yaml`

```yaml
---
apiVersion: cache.example.com/v1alpha1
kind: Memcached
metadata:
  name: memcached-sample
spec:
  size: 3
```

Define your image and container registry:

```bash
export IMG=docker.io/imjoseangel/memcached-operator:v1
```

And run:

```bash
make docker-build docker-push IMG=$IMG
```

>**Note**: Be sure that there are no extra spaces in your PATH environment or the command will fail.

### Run the Operator

### Apply the Memcached Kind (CRD):

```bash
make install
```

### Deploy the operator:

```bash
export IMG=docker.io/imjoseangel/memcached-operator:v1
make deploy
```

Verify that the memcached-operator is up and running:

```bash
kubectl get deployment -n memcached-operator-system
```

### Create a memcached resource

```bash
kubectl apply -f config/samples/cache_v1alpha1_memcached.yaml -n memcached-operator-system
```

Verify that Memcached pods are created

```bash
kubectl get pods -n memcached-operator-system

NAME                                                     READY   STATUS    RESTARTS   AGE
memcached-operator-controller-manager-5d95cd576f-npmcl   2/2     Running   0          30s
memcached-sample-memcached-b885dcc75-69k8d               1/1     Running   0          21s
memcached-sample-memcached-b885dcc75-8s66v               1/1     Running   0          21s
memcached-sample-memcached-b885dcc75-s9blm               1/1     Running   0          21s


kubectl get all -n memcached-operator-system


NAME                                                         READY   STATUS    RESTARTS   AGE
pod/memcached-operator-controller-manager-5d95cd576f-npmcl   2/2     Running   0          66s
pod/memcached-sample-memcached-b885dcc75-69k8d               1/1     Running   0          57s
pod/memcached-sample-memcached-b885dcc75-8s66v               1/1     Running   0          57s
pod/memcached-sample-memcached-b885dcc75-s9blm               1/1     Running   0          57s

NAME                                                            TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)    AGE
service/memcached-operator-controller-manager-metrics-service   ClusterIP   10.0.5.131   <none>        8443/TCP   68s

NAME                                                    READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/memcached-operator-controller-manager   1/1     1            1           67s
deployment.apps/memcached-sample-memcached              3/3     3            3           58s

NAME                                                               DESIRED   CURRENT   READY   AGE
replicaset.apps/memcached-operator-controller-manager-5d95cd576f   1         1         1       67s
replicaset.apps/memcached-sample-memcached-b885dcc75               3         3         3       58s
```

### Cleanup

To leave the operator, but remove the memcached sample pods, delete the CR.

```bash
kubectl delete -f config/samples/cache_v1alpha1_memcached.yaml -n memcached-operator-system
```

To clean up everything:

```bash
make undeploy
```

## Troubleshooting

Run the following command to check the operator logs.

```bash
kubectl logs deployment.apps/memcached-operator-controller-manager -n memcached-operator-system -c manager
```
