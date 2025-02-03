/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// VPCPeeringConnectionParameters defines the desired state of VPCPeeringConnection
type VPCPeeringConnectionParameters struct {
	// Region is which region the VPCPeeringConnection will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The Amazon Web Services account ID of the owner of the accepter VPC.
	//
	// Default: Your Amazon Web Services account ID
	PeerOwnerID *string `json:"peerOwnerID,omitempty"`
	// The Region code for the accepter VPC, if the accepter VPC is located in a
	// Region other than the Region in which you make the request.
	//
	// Default: The Region in which you make the request.
	PeerRegion *string `json:"peerRegion,omitempty"`
	// The tags to assign to the peering connection.
	TagSpecifications                    []*TagSpecification `json:"tagSpecifications,omitempty"`
	CustomVPCPeeringConnectionParameters `json:",inline"`
}

// VPCPeeringConnectionSpec defines the desired state of VPCPeeringConnection
type VPCPeeringConnectionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VPCPeeringConnectionParameters `json:"forProvider"`
}

// VPCPeeringConnectionObservation defines the observed state of VPCPeeringConnection
type VPCPeeringConnectionObservation struct {
	// Information about the accepter VPC. CIDR block information is only returned
	// when describing an active VPC peering connection.
	AccepterVPCInfo *VPCPeeringConnectionVPCInfo `json:"accepterVPCInfo,omitempty"`
	// The time that an unaccepted VPC peering connection will expire.
	ExpirationTime *metav1.Time `json:"expirationTime,omitempty"`
	// Information about the requester VPC. CIDR block information is only returned
	// when describing an active VPC peering connection.
	RequesterVPCInfo *VPCPeeringConnectionVPCInfo `json:"requesterVPCInfo,omitempty"`
	// The status of the VPC peering connection.
	Status *VPCPeeringConnectionStateReason `json:"status,omitempty"`
	// Any tags assigned to the resource.
	Tags []*Tag `json:"tags,omitempty"`
	// The ID of the VPC peering connection.
	VPCPeeringConnectionID *string `json:"vpcPeeringConnectionID,omitempty"`

	CustomVPCPeeringConnectionObservation `json:",inline"`
}

// VPCPeeringConnectionStatus defines the observed state of VPCPeeringConnection.
type VPCPeeringConnectionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VPCPeeringConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCPeeringConnection is the Schema for the VPCPeeringConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCPeeringConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCPeeringConnectionSpec   `json:"spec"`
	Status            VPCPeeringConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCPeeringConnectionList contains a list of VPCPeeringConnections
type VPCPeeringConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCPeeringConnection `json:"items"`
}

// Repository type metadata.
var (
	VPCPeeringConnectionKind             = "VPCPeeringConnection"
	VPCPeeringConnectionGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCPeeringConnectionKind}.String()
	VPCPeeringConnectionKindAPIVersion   = VPCPeeringConnectionKind + "." + GroupVersion.String()
	VPCPeeringConnectionGroupVersionKind = GroupVersion.WithKind(VPCPeeringConnectionKind)
)

func init() {
	SchemeBuilder.Register(&VPCPeeringConnection{}, &VPCPeeringConnectionList{})
}
