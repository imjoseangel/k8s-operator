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

```yaml

```
