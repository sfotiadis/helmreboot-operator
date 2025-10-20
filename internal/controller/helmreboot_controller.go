/*
Copyright 2025.

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

package controller

import (
	"context"
	"strings"
	"time"

	helmv2 "github.com/fluxcd/helm-controller/api/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// HelmRebootReconciler reconciles a HelmReboot object
type HelmRebootReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=helm.io.github.sfotiadis,resources=helmreboots,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=helm.io.github.sfotiadis,resources=helmreboots/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=helm.io.github.sfotiadis,resources=helmreboots/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the HelmReboot object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.22.1/pkg/reconcile
func (r *HelmRebootReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var hr helmv2.HelmRelease
	if err := r.Get(ctx, req.NamespacedName, &hr); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log := ctrl.LoggerFrom(ctx)
	log.Info("Processing HelmRelease", "name", hr.Name, "conditions", len(hr.Status.Conditions))

	for _, cond := range hr.Status.Conditions {
		log.Info("Checking condition", "type", cond.Type, "status", cond.Status, "message", cond.Message)
		if cond.Type == "Ready" && cond.Status == metav1.ConditionFalse &&
			strings.Contains(cond.Message, "context deadline exceeded") {

			log.Info("Found timeout condition, adding annotation")
			// Annotation setzen, um Flux Reconcile zu triggern
			patch := client.MergeFrom(hr.DeepCopy())
			if hr.Annotations == nil {
				hr.Annotations = map[string]string{}
			}
			hr.Annotations["fluxcd.io/reconcileAt"] = time.Now().Format(time.RFC3339)
			hr.Annotations["helmreboot.myorg.io/lastRestart"] = time.Now().Format(time.RFC3339)

			if err := r.Patch(ctx, &hr, patch); err != nil {
				return ctrl.Result{}, err
			}

			log.Info("Restarted HelmRelease due to context deadline exceeded", "name", hr.Name)
		}
	}

	return ctrl.Result{RequeueAfter: 10 * time.Minute}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *HelmRebootReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&helmv2.HelmRelease{}).
		Named("helmreboot").
		Complete(r)
}
