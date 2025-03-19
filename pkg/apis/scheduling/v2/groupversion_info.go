// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

// Package v2 contains API Schema definitions for the scheduling v2 API group
// +kubebuilder:object:generate=true
// +groupName=scheduling.run.ai
package v2

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "scheduling.run.ai", Version: "v2"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

// SchemeGroupVersion is the group version used to register these objects.
var SchemeGroupVersion = GroupVersion

// Resource takes an unqualified resource and returns a Group-qualified GroupResource.
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
