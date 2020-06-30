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

type SqlFailoverGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlFailoverGroupSpec   `json:"spec,omitempty"`
	Status            SqlFailoverGroupStatus `json:"status,omitempty"`
}

type SqlFailoverGroupSpecPartnerServers struct {
	ID string `json:"ID" tf:"id"`
	// +optional
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	Role string `json:"role,omitempty" tf:"role,omitempty"`
}

type SqlFailoverGroupSpecReadWriteEndpointFailoverPolicy struct {
	// +optional
	GraceMinutes int64  `json:"graceMinutes,omitempty" tf:"grace_minutes,omitempty"`
	Mode         string `json:"mode" tf:"mode"`
}

type SqlFailoverGroupSpecReadonlyEndpointFailoverPolicy struct {
	Mode string `json:"mode" tf:"mode"`
}

type SqlFailoverGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Databases []string `json:"databases,omitempty" tf:"databases,omitempty"`
	// +optional
	Location       string                               `json:"location,omitempty" tf:"location,omitempty"`
	Name           string                               `json:"name" tf:"name"`
	PartnerServers []SqlFailoverGroupSpecPartnerServers `json:"partnerServers" tf:"partner_servers"`
	// +kubebuilder:validation:MaxItems=1
	ReadWriteEndpointFailoverPolicy []SqlFailoverGroupSpecReadWriteEndpointFailoverPolicy `json:"readWriteEndpointFailoverPolicy" tf:"read_write_endpoint_failover_policy"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ReadonlyEndpointFailoverPolicy []SqlFailoverGroupSpecReadonlyEndpointFailoverPolicy `json:"readonlyEndpointFailoverPolicy,omitempty" tf:"readonly_endpoint_failover_policy,omitempty"`
	ResourceGroupName              string                                               `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Role       string `json:"role,omitempty" tf:"role,omitempty"`
	ServerName string `json:"serverName" tf:"server_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SqlFailoverGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SqlFailoverGroupSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqlFailoverGroupList is a list of SqlFailoverGroups
type SqlFailoverGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqlFailoverGroup CRD objects
	Items []SqlFailoverGroup `json:"items,omitempty"`
}
