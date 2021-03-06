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

type BackupPlan struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupPlanSpec   `json:"spec,omitempty"`
	Status            BackupPlanStatus `json:"status,omitempty"`
}

type BackupPlanSpecRuleLifecycle struct {
	// +optional
	ColdStorageAfter int64 `json:"coldStorageAfter,omitempty" tf:"cold_storage_after,omitempty"`
	// +optional
	DeleteAfter int64 `json:"deleteAfter,omitempty" tf:"delete_after,omitempty"`
}

type BackupPlanSpecRule struct {
	// +optional
	CompletionWindow int64 `json:"completionWindow,omitempty" tf:"completion_window,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Lifecycle []BackupPlanSpecRuleLifecycle `json:"lifecycle,omitempty" tf:"lifecycle,omitempty"`
	// +optional
	RecoveryPointTags map[string]string `json:"recoveryPointTags,omitempty" tf:"recovery_point_tags,omitempty"`
	RuleName          string            `json:"ruleName" tf:"rule_name"`
	// +optional
	Schedule string `json:"schedule,omitempty" tf:"schedule,omitempty"`
	// +optional
	StartWindow     int64  `json:"startWindow,omitempty" tf:"start_window,omitempty"`
	TargetVaultName string `json:"targetVaultName" tf:"target_vault_name"`
}

type BackupPlanSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn  string               `json:"arn,omitempty" tf:"arn,omitempty"`
	Name string               `json:"name" tf:"name"`
	Rule []BackupPlanSpecRule `json:"rule" tf:"rule"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
}

type BackupPlanStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *BackupPlanSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BackupPlanList is a list of BackupPlans
type BackupPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BackupPlan CRD objects
	Items []BackupPlan `json:"items,omitempty"`
}
