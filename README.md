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

$ operator-sdk init --plugins=ansible.sdk.operatorframework.io/v1 --domain example.com
```

### Create the API

```bash
$ operator-sdk create api --group apps --version v1alpha1 --kind AppService --generate-role --generate-playbook
```

### Define the API

This command specifies that the CRD will be called **AppService** and creates the role **roles/appservice**, which you can modify to specify the input parameter of your CRD:

```go
// PresentationSpec defines the desired state of Presentation
type PresentationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Presentation. Edit Presentation_types.go to remove/update
	Markdown string `json:"markdown,omitempty"`
}
```

After modifying the `*_types.go` file always run the following command to update the generated code for that resource type

```bash
$ make generate
```

### Generating CRD manifests

Once the API is defined with spec/status fields and CRD validation markers, the CRD manifests can be generated and updated with the following command:

```bash
$ make manifests
```

This makefile target will invoke controller-gen to generate the CRD manifests at `config/crd/bases/presentation.example.com_presentations.yaml`.

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
