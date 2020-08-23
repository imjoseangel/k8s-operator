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
mkdir ~/k8s-operator && cd ~/k8s-operator/

operator-sdk init --project-version="2" --repo github.com/imjoseangel/k8s-operator --owner "imjoseangel" --domain example.com
```

### Create the API

```bash
operator-sdk create api --kind Presentation --group presentation --version v2
```

### Configure the test environment

[Setup the envtest binaries and environment](https://sdk.operatorframework.io/docs/building-operators/golang/references/envtest-setup) for your project. Update your test Makefile target to the following:

```makefile
# Run tests
ENVTEST_ASSETS_DIR=$(shell pwd)/testbin
test: generate fmt vet manifests
	mkdir -p ${ENVTEST_ASSETS_DIR}
	test -f ${ENVTEST_ASSETS_DIR}/setup-envtest.sh || curl -sSLo ${ENVTEST_ASSETS_DIR}/setup-envtest.sh https://raw.githubusercontent.com/kubernetes-sigs/controller-runtime/master/hack/setup-envtest.sh
	source ${ENVTEST_ASSETS_DIR}/setup-envtest.sh; fetch_envtest_tools $(ENVTEST
```
