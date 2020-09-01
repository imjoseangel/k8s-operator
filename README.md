# Ansible Kubernetes Operator

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
mkdir ~/restapi-operator  && cd ~/restapi-operator/
operator-sdk init --plugins=ansible --domain example.com
```

### Create the API

```bash
operator-sdk create api --group api --version v1alpha1 --kind Restapi --generate-role
```

### Modify the Manager

Update the file `roles/restapi/tasks/main.yml:`

```yaml
---
- name: start Restapi
  community.kubernetes.k8s:
    definition:
      kind: Deployment
      apiVersion: apps/v1
      metadata:
        name: '{{ ansible_operator_meta.name }}-restapi'
        namespace: '{{ ansible_operator_meta.namespace }}'
      spec:
        replicas: "{{ size }}"
        selector:
          matchLabels:
            app: restapi
        template:
          metadata:
            labels:
              app: restapi
          spec:
            containers:
            - name: restapi
              image: "docker.io/imjoseangel/restapi:latest"
              imagePullPolicy: Always
              ports:
                - containerPort: 5000
              resources:
                requests: # minimum resources required
                  cpu: 250m
                  memory: 64Mi
                limits: # maximum resources allocated
                  cpu: 500m
                  memory: 256Mi
              readinessProbe: # is the container ready to receive traffic?
                httpGet:
                  port: 5000
                  path: /
              livenessProbe: # is the container healthy?
                httpGet:
                  port: 5000
                  path: /
```

Add the defaults to `roles/restapi/defaults/main.yml`

```yaml
---
# defaults file for restapi
size: 1
```

And update the file `config/samples/api_v1alpha1_restapi.yaml`

```yaml
---
apiVersion: restapi.example.com/v1alpha1
kind: Restapi
metadata:
  name: restapi-sample
spec:
  size: 3
```

Define your image and container registry:

```bash
export IMG=docker.io/imjoseangel/restapi-operator:v1
```

And run:

```bash
make docker-build docker-push IMG=$IMG
```

>**Note**: Be sure that there are no extra spaces in your PATH environment or the command will fail.

### Run the Operator

### Create the namespace

```bash
kubectl create namespace restapi-operator-system
```

### Apply the restapi Kind (CRD)

```bash
make install
```

### Deploy the operator

```bash
export IMG=docker.io/imjoseangel/restapi-operator:v1
make deploy
```

Verify that the restapi-operator is up and running:

```bash
kubectl get deployment -n restapi-operator-system
```

### Create a restapi resource

```bash
kubectl apply -f config/samples/api_v1alpha1_restapi.yaml -n restapi-operator-system
```

Verify that restapi pods are created

```bash
kubectl get pods -n restapi-operator-system

NAME                                                     READY   STATUS    RESTARTS   AGE
restapi-operator-controller-manager-5d95cd576f-npmcl     2/2     Running   0          30s
restapi-sample-restapi-b885dcc75-69k8d                   1/1     Running   0          21s
restapi-sample-restapi-b885dcc75-8s66v                   1/1     Running   0          21s
restapi-sample-restapi-b885dcc75-s9blm                   1/1     Running   0          21s


kubectl get all -n restapi-operator-system


NAME                                                         READY   STATUS    RESTARTS   AGE
pod/restapi-operator-controller-manager-5d95cd576f-npmcl     2/2     Running   0          66s
pod/restapi-sample-restapi-b885dcc75-69k8d                   1/1     Running   0          57s
pod/restapi-sample-restapi-b885dcc75-8s66v                   1/1     Running   0          57s
pod/restapi-sample-restapi-b885dcc75-s9blm                   1/1     Running   0          57s

NAME                                                            TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)    AGE
service/restapi-operator-controller-manager-metrics-service     ClusterIP   10.0.5.131   <none>        8443/TCP   68s

NAME                                                    READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/restapi-operator-controller-manager     1/1     1            1           67s
deployment.apps/restapi-sample-restapi                  3/3     3            3           58s

NAME                                                               DESIRED   CURRENT   READY   AGE
replicaset.apps/restapi-operator-controller-manager-5d95cd576f     1         1         1       67s
replicaset.apps/restapi-sample-restapi-b885dcc75                   3         3         3       58s
```

### Cleanup

To leave the operator, but remove the restapi sample pods, delete the CR.

```bash
kubectl delete -f config/samples/api_v1alpha1_restapi.yaml -n restapi-operator-system
```

To clean up everything:

```bash
make undeploy
```

## Troubleshooting

Run the following command to check the operator logs.

```bash
kubectl logs deployment.apps/restapi-operator-controller-manager -n restapi-operator-system -c manager
```

## More Info

[Tutorial](https://learn.openshift.com/ansibleop/ansible-operator-overview/?extIdCarryOver=true&intcmp=701f20000012k6TAAQ&sc_cid=701f2000001Css5AAC)
