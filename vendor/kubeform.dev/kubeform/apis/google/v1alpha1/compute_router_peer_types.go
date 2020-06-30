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

type ComputeRouterPeer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeRouterPeerSpec   `json:"spec,omitempty"`
	Status            ComputeRouterPeerStatus `json:"status,omitempty"`
}

type ComputeRouterPeerSpecAdvertisedIPRanges struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Range string `json:"range,omitempty" tf:"range,omitempty"`
}

type ComputeRouterPeerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdvertiseMode string `json:"advertiseMode,omitempty" tf:"advertise_mode,omitempty"`
	// +optional
	AdvertisedGroups []string `json:"advertisedGroups,omitempty" tf:"advertised_groups,omitempty"`
	// +optional
	AdvertisedIPRanges []ComputeRouterPeerSpecAdvertisedIPRanges `json:"advertisedIPRanges,omitempty" tf:"advertised_ip_ranges,omitempty"`
	// +optional
	AdvertisedRoutePriority int64  `json:"advertisedRoutePriority,omitempty" tf:"advertised_route_priority,omitempty"`
	Interface               string `json:"interface" tf:"interface"`
	// +optional
	IpAddress string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`
	Name      string `json:"name" tf:"name"`
	PeerAsn   int64  `json:"peerAsn" tf:"peer_asn"`
	// +optional
	PeerIPAddress string `json:"peerIPAddress,omitempty" tf:"peer_ip_address,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	Router string `json:"router" tf:"router"`
}

type ComputeRouterPeerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeRouterPeerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeRouterPeerList is a list of ComputeRouterPeers
type ComputeRouterPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeRouterPeer CRD objects
	Items []ComputeRouterPeer `json:"items,omitempty"`
}
