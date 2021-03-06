/*
Copyright The Kubeform Authors.

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ComputeInstanceGroupManager struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeInstanceGroupManagerSpec   `json:"spec,omitempty"`
	Status            ComputeInstanceGroupManagerStatus `json:"status,omitempty"`
}

type ComputeInstanceGroupManagerSpecNamedPort struct {
	Name string `json:"name" tf:"name"`
	Port int64  `json:"port" tf:"port"`
}

type ComputeInstanceGroupManagerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	BaseInstanceName string `json:"baseInstanceName" tf:"base_instance_name"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Fingerprint string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`
	// +optional
	InstanceGroup    string `json:"instanceGroup,omitempty" tf:"instance_group,omitempty"`
	InstanceTemplate string `json:"instanceTemplate" tf:"instance_template"`
	Name             string `json:"name" tf:"name"`
	// +optional
	NamedPort []ComputeInstanceGroupManagerSpecNamedPort `json:"namedPort,omitempty" tf:"named_port,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	// +optional
	TargetPools []string `json:"targetPools,omitempty" tf:"target_pools,omitempty"`
	// +optional
	TargetSize int64 `json:"targetSize,omitempty" tf:"target_size,omitempty"`
	// +optional
	UpdateStrategy string `json:"updateStrategy,omitempty" tf:"update_strategy,omitempty"`
	// +optional
	WaitForInstances bool `json:"waitForInstances,omitempty" tf:"wait_for_instances,omitempty"`
	// +optional
	Zone string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ComputeInstanceGroupManagerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeInstanceGroupManagerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeInstanceGroupManagerList is a list of ComputeInstanceGroupManagers
type ComputeInstanceGroupManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeInstanceGroupManager CRD objects
	Items []ComputeInstanceGroupManager `json:"items,omitempty"`
}
