package controller

import (
	"github.com/imjoseangel/k8s-operator/pkg/controller/presentation"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, presentation.Add)
}
