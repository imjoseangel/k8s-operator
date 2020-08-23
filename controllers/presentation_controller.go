/*
Copyright 2020 imjoseangel.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/prometheus/common/log"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	presentationv1 "github.com/imjoseangel/k8s-operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
)

// PresentationReconciler reconciles a Presentation object
type PresentationReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=presentation.example.com,resources=presentations,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=presentation.example.com,resources=presentations/status,verbs=get;update;patch

func (r *PresentationReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("presentation", req.NamespacedName)

	// your logic here
	// Fetch the Presentation instance
	presentation := &presentationv1.Presentation{}
	err := r.Get(ctx, req.NamespacedName, presentation)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			log.Info("Presentation resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get Presentation")
		return ctrl.Result{}, err
	}

	// Check if the deployment already exists, if not create a new one
	found := &appsv1.Deployment{}
	err = r.Get(ctx, types.NamespacedName{Name: presentation.Name, Namespace: presentation.Namespace}, found)

	return ctrl.Result{}, nil
}

func (r *PresentationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&presentationv1.Presentation{}).
		Complete(r)
}
