// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package v2

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var queuelog = logf.Log.WithName("queue-resource")

const missingResourcesError = "resources must be specified"

func (r *Queue) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		WithValidator(&Queue{}).
		Complete()
}

func (_ *Queue) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	queue, ok := obj.(*Queue)
	if !ok {
		return nil, fmt.Errorf("expected a Queue but got a %T", obj)
	}
	queuelog.Info("validate create", "name", queue.Name)

	if queue.Spec.Resources == nil {
		return []string{missingResourcesError}, fmt.Errorf(missingResourcesError)
	}
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (_ *Queue) ValidateUpdate(_ context.Context, _ runtime.Object, newObj runtime.Object) (admission.Warnings, error) {
	queue, ok := newObj.(*Queue)
	if !ok {
		return nil, fmt.Errorf("expected a Queue but got a %T", newObj)
	}
	queuelog.Info("validate update", "name", queue.Name)

	if queue.Spec.Resources == nil {
		return []string{missingResourcesError}, fmt.Errorf(missingResourcesError)
	}
	return nil, nil
}

func (_ *Queue) ValidateDelete(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	queue, ok := obj.(*Queue)
	if !ok {
		return nil, fmt.Errorf("expected a Queue but got a %T", obj)
	}
	queuelog.Info("validate delete", "name", queue.Name)
	return nil, nil
}
