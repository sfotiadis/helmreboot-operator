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

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	helmv2 "github.com/fluxcd/helm-controller/api/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
)

var _ = Describe("HelmReboot Controller", func() {
	Context("When reconciling a HelmRelease with timeout error", func() {
		It("should add a fluxcd.io/reconcileAt annotation", func() {
			ctx := context.Background()

			hr := &helmv2.HelmRelease{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-hr",
					Namespace: "default",
				},
				Spec: helmv2.HelmReleaseSpec{
					Chart: &helmv2.HelmChartTemplate{
						Spec: helmv2.HelmChartTemplateSpec{
							Chart:   "nginx",
							Version: "1.0.0",
							SourceRef: helmv2.CrossNamespaceObjectReference{
								Kind: "HelmRepository",
								Name: "test-repo",
							},
						},
					},
				},
			}

			Expect(k8sClient.Create(ctx, hr)).To(Succeed())

			// Update the status separately since status is usually managed by controllers
			hr.Status = helmv2.HelmReleaseStatus{
				Conditions: []metav1.Condition{
					{
						Type:               "Ready",
						Status:             metav1.ConditionFalse,
						Reason:             "ReconciliationFailed",
						Message:            "context deadline exceeded",
						LastTransitionTime: metav1.Now(),
					},
				},
			}
			Expect(k8sClient.Status().Update(ctx, hr)).To(Succeed())

			reconciler := &HelmRebootReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := reconciler.Reconcile(ctx, ctrl.Request{
				NamespacedName: types.NamespacedName{
					Name:      hr.Name,
					Namespace: hr.Namespace,
				},
			})
			Expect(err).ToNot(HaveOccurred())

			updated := &helmv2.HelmRelease{}
			Expect(k8sClient.Get(ctx, types.NamespacedName{
				Name:      hr.Name,
				Namespace: hr.Namespace,
			}, updated)).To(Succeed())

			Expect(updated.Annotations).To(HaveKey("fluxcd.io/reconcileAt"))
		})
	})
})
