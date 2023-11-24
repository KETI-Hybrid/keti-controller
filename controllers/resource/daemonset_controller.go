/*
Copyright 2023.

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

package resource

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	resourcev1 "github.com/KETI-Hybrid/keti-controller/apis/resource/v1"
)

// DaemonsetReconciler reconciles a Daemonset object
type DaemonsetReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *DaemonsetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var err error
	result := &ClientManager{}

	result.KetiClient, err = crd.NewClient()
	if err != nil {
		klog.Errorln(err)
	}
	result.KubeClient, err = k8s.NewClient()
	if err != nil {
		klog.Errorln(err)
	}

	return result

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DaemonsetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&resourcev1.Daemonset{}).
		Complete(r)
}
