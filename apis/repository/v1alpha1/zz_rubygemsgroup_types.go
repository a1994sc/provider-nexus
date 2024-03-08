// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RubygemsGroupGroupInitParameters struct {

	// (Set of String) Member repositories names
	// Member repositories names
	MemberNames []*string `json:"memberNames,omitempty" tf:"member_names,omitempty"`
}

type RubygemsGroupGroupObservation struct {

	// (Set of String) Member repositories names
	// Member repositories names
	MemberNames []*string `json:"memberNames,omitempty" tf:"member_names,omitempty"`
}

type RubygemsGroupGroupParameters struct {

	// (Set of String) Member repositories names
	// Member repositories names
	// +kubebuilder:validation:Optional
	MemberNames []*string `json:"memberNames" tf:"member_names,omitempty"`
}

type RubygemsGroupInitParameters struct {

	// (Block List, Min: 1, Max: 1) Configuration for repository group (see below for nested schema)
	// Configuration for repository group
	Group []RubygemsGroupGroupInitParameters `json:"group,omitempty" tf:"group,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	Storage []RubygemsGroupStorageInitParameters `json:"storage,omitempty" tf:"storage,omitempty"`
}

type RubygemsGroupObservation struct {

	// (Block List, Min: 1, Max: 1) Configuration for repository group (see below for nested schema)
	// Configuration for repository group
	Group []RubygemsGroupGroupObservation `json:"group,omitempty" tf:"group,omitempty"`

	// (String) Used to identify resource at nexus
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	Storage []RubygemsGroupStorageObservation `json:"storage,omitempty" tf:"storage,omitempty"`
}

type RubygemsGroupParameters struct {

	// (Block List, Min: 1, Max: 1) Configuration for repository group (see below for nested schema)
	// Configuration for repository group
	// +kubebuilder:validation:Optional
	Group []RubygemsGroupGroupParameters `json:"group,omitempty" tf:"group,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	// +kubebuilder:validation:Optional
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	// +kubebuilder:validation:Optional
	Storage []RubygemsGroupStorageParameters `json:"storage,omitempty" tf:"storage,omitempty"`
}

type RubygemsGroupStorageInitParameters struct {

	// (String) Blob store used to store repository contents
	// Blob store used to store repository contents
	BlobStoreName *string `json:"blobStoreName,omitempty" tf:"blob_store_name,omitempty"`

	// (Boolean) Whether to validate uploaded content's MIME type appropriate for the repository format
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation *bool `json:"strictContentTypeValidation,omitempty" tf:"strict_content_type_validation,omitempty"`
}

type RubygemsGroupStorageObservation struct {

	// (String) Blob store used to store repository contents
	// Blob store used to store repository contents
	BlobStoreName *string `json:"blobStoreName,omitempty" tf:"blob_store_name,omitempty"`

	// (Boolean) Whether to validate uploaded content's MIME type appropriate for the repository format
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation *bool `json:"strictContentTypeValidation,omitempty" tf:"strict_content_type_validation,omitempty"`
}

type RubygemsGroupStorageParameters struct {

	// (String) Blob store used to store repository contents
	// Blob store used to store repository contents
	// +kubebuilder:validation:Optional
	BlobStoreName *string `json:"blobStoreName" tf:"blob_store_name,omitempty"`

	// (Boolean) Whether to validate uploaded content's MIME type appropriate for the repository format
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	// +kubebuilder:validation:Optional
	StrictContentTypeValidation *bool `json:"strictContentTypeValidation,omitempty" tf:"strict_content_type_validation,omitempty"`
}

// RubygemsGroupSpec defines the desired state of RubygemsGroup
type RubygemsGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RubygemsGroupParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RubygemsGroupInitParameters `json:"initProvider,omitempty"`
}

// RubygemsGroupStatus defines the observed state of RubygemsGroup.
type RubygemsGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RubygemsGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RubygemsGroup is the Schema for the RubygemsGroups API. Use this resource to create a group rubygems repository.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nexus}
type RubygemsGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.group) || (has(self.initProvider) && has(self.initProvider.group))",message="spec.forProvider.group is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storage) || (has(self.initProvider) && has(self.initProvider.storage))",message="spec.forProvider.storage is a required parameter"
	Spec   RubygemsGroupSpec   `json:"spec"`
	Status RubygemsGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RubygemsGroupList contains a list of RubygemsGroups
type RubygemsGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RubygemsGroup `json:"items"`
}

// Repository type metadata.
var (
	RubygemsGroup_Kind             = "RubygemsGroup"
	RubygemsGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RubygemsGroup_Kind}.String()
	RubygemsGroup_KindAPIVersion   = RubygemsGroup_Kind + "." + CRDGroupVersion.String()
	RubygemsGroup_GroupVersionKind = CRDGroupVersion.WithKind(RubygemsGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&RubygemsGroup{}, &RubygemsGroupList{})
}
