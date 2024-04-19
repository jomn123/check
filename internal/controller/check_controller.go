/*
Copyright 2024.

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
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CheckReconciler reconciles a Check object
type CheckReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=ctf.check.com,resources=checks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ctf.check.com,resources=checks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ctf.check.com,resources=checks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Check object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile

func (r *CheckReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	podNames := []string{"x", "y", "w", "z"}

	for {
		select {
		case <-ctx.Done():
			// If the context is cancelled, stop the ticker and return
			return ctrl.Result{}, nil
		case <-ticker.C:
			// Every minute, check the status of the pods
			for _, podName := range podNames {
				pod := &corev1.Pod{}
				err := r.Get(ctx, client.ObjectKey{Namespace: "default", Name: podName}, pod)
				if err != nil { // Fix: Added formatting directive for the error message
					continue
				}

				creationTime := pod.GetCreationTimestamp()
				timeSinceCreation := time.Since(creationTime.Time)
				fmt.Printf("Pod exists: podName=%s, timeSinceCreation=%v\n", podName, timeSinceCreation) // Fix: Added formatting directives for podName and timeSinceCreation
			}
		}
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *CheckReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Pod{}).
		Complete(r)
}
