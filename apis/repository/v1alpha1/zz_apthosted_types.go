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

type AptHostedInitParameters struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	Cleanup []CleanupInitParameters `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Max: 1) Component configuration for the hosted repository (see below for nested schema)
	// Component configuration for the hosted repository
	Component []ComponentInitParameters `json:"component,omitempty" tf:"component,omitempty"`

	// (String) Distribution to fetch
	// Distribution to fetch
	Distribution *string `json:"distribution,omitempty" tf:"distribution,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// (Block List, Min: 1, Max: 1) Signing contains signing data of hosted repositores of format Apt (see below for nested schema)
	// Signing contains signing data of hosted repositores of format Apt
	Signing []SigningInitParameters `json:"signing,omitempty" tf:"signing,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	Storage []StorageInitParameters `json:"storage,omitempty" tf:"storage,omitempty"`
}

type AptHostedObservation struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	Cleanup []CleanupObservation `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Max: 1) Component configuration for the hosted repository (see below for nested schema)
	// Component configuration for the hosted repository
	Component []ComponentObservation `json:"component,omitempty" tf:"component,omitempty"`

	// (String) Distribution to fetch
	// Distribution to fetch
	Distribution *string `json:"distribution,omitempty" tf:"distribution,omitempty"`

	// (String) Used to identify resource at nexus
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// (Block List, Min: 1, Max: 1) Signing contains signing data of hosted repositores of format Apt (see below for nested schema)
	// Signing contains signing data of hosted repositores of format Apt
	Signing []SigningParameters `json:"signing,omitempty" tf:"signing,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	Storage []StorageObservation `json:"storage,omitempty" tf:"storage,omitempty"`
}

type AptHostedParameters struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	// +kubebuilder:validation:Optional
	Cleanup []CleanupParameters `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Max: 1) Component configuration for the hosted repository (see below for nested schema)
	// Component configuration for the hosted repository
	// +kubebuilder:validation:Optional
	Component []ComponentParameters `json:"component,omitempty" tf:"component,omitempty"`

	// (String) Distribution to fetch
	// Distribution to fetch
	// +kubebuilder:validation:Optional
	Distribution *string `json:"distribution,omitempty" tf:"distribution,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	// +kubebuilder:validation:Optional
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// (Block List, Min: 1, Max: 1) Signing contains signing data of hosted repositores of format Apt (see below for nested schema)
	// Signing contains signing data of hosted repositores of format Apt
	// +kubebuilder:validation:Optional
	Signing []SigningParameters `json:"signing,omitempty" tf:"signing,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	// +kubebuilder:validation:Optional
	Storage []StorageParameters `json:"storage,omitempty" tf:"storage,omitempty"`
}

type CleanupInitParameters struct {

	// (Set of String) List of policy names
	// List of policy names
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type CleanupObservation struct {

	// (Set of String) List of policy names
	// List of policy names
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type CleanupParameters struct {

	// (Set of String) List of policy names
	// List of policy names
	// +kubebuilder:validation:Optional
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type ComponentInitParameters struct {

	// (Boolean) Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	ProprietaryComponents *bool `json:"proprietaryComponents,omitempty" tf:"proprietary_components,omitempty"`
}

type ComponentObservation struct {

	// (Boolean) Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	ProprietaryComponents *bool `json:"proprietaryComponents,omitempty" tf:"proprietary_components,omitempty"`
}

type ComponentParameters struct {

	// (Boolean) Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// +kubebuilder:validation:Optional
	ProprietaryComponents *bool `json:"proprietaryComponents" tf:"proprietary_components,omitempty"`
}

type SigningInitParameters struct {
}

type SigningObservation struct {
}

type SigningParameters struct {

	// -export-secret-key --armor)
	// If passphrase is unset, the keypair cannot be read from the nexus api.
	// When reading the resource, the keypair will be read from the previous state,
	// so external changes won't be detected in this case.
	// PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor)
	// If passphrase is unset, the keypair cannot be read from the nexus api.
	// When reading the resource, the keypair will be read from the previous state,
	// so external changes won't be detected in this case.
	// +kubebuilder:validation:Required
	KeypairSecretRef v1.SecretKeySelector `json:"keypairSecretRef" tf:"-"`

	// (String, Sensitive) Passphrase to access PGP signing key.
	// This value cannot be read from the nexus api.
	// When reading the resource, the value will be read from the previous state,
	// so external changes won't be detected.
	// Passphrase to access PGP signing key.
	// This value cannot be read from the nexus api.
	// When reading the resource, the value will be read from the previous state,
	// so external changes won't be detected.
	// +kubebuilder:validation:Optional
	PassphraseSecretRef *v1.SecretKeySelector `json:"passphraseSecretRef,omitempty" tf:"-"`
}

type StorageInitParameters struct {

	// (String) Blob store used to store repository contents
	// Blob store used to store repository contents
	BlobStoreName *string `json:"blobStoreName,omitempty" tf:"blob_store_name,omitempty"`

	// (Boolean) Whether to validate uploaded content's MIME type appropriate for the repository format
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation *bool `json:"strictContentTypeValidation,omitempty" tf:"strict_content_type_validation,omitempty"`

	// (String) Controls if deployments of and updates to assets are allowed
	// Controls if deployments of and updates to assets are allowed
	WritePolicy *string `json:"writePolicy,omitempty" tf:"write_policy,omitempty"`
}

type StorageObservation struct {

	// (String) Blob store used to store repository contents
	// Blob store used to store repository contents
	BlobStoreName *string `json:"blobStoreName,omitempty" tf:"blob_store_name,omitempty"`

	// (Boolean) Whether to validate uploaded content's MIME type appropriate for the repository format
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation *bool `json:"strictContentTypeValidation,omitempty" tf:"strict_content_type_validation,omitempty"`

	// (String) Controls if deployments of and updates to assets are allowed
	// Controls if deployments of and updates to assets are allowed
	WritePolicy *string `json:"writePolicy,omitempty" tf:"write_policy,omitempty"`
}

type StorageParameters struct {

	// (String) Blob store used to store repository contents
	// Blob store used to store repository contents
	// +kubebuilder:validation:Optional
	BlobStoreName *string `json:"blobStoreName" tf:"blob_store_name,omitempty"`

	// (Boolean) Whether to validate uploaded content's MIME type appropriate for the repository format
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	// +kubebuilder:validation:Optional
	StrictContentTypeValidation *bool `json:"strictContentTypeValidation" tf:"strict_content_type_validation,omitempty"`

	// (String) Controls if deployments of and updates to assets are allowed
	// Controls if deployments of and updates to assets are allowed
	// +kubebuilder:validation:Optional
	WritePolicy *string `json:"writePolicy,omitempty" tf:"write_policy,omitempty"`
}

// AptHostedSpec defines the desired state of AptHosted
type AptHostedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AptHostedParameters `json:"forProvider"`
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
	InitProvider AptHostedInitParameters `json:"initProvider,omitempty"`
}

// AptHostedStatus defines the observed state of AptHosted.
type AptHostedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AptHostedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AptHosted is the Schema for the AptHosteds API. Use this resource to create a hosted apt repository.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nexus}
type AptHosted struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.distribution) || (has(self.initProvider) && has(self.initProvider.distribution))",message="spec.forProvider.distribution is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.signing) || (has(self.initProvider) && has(self.initProvider.signing))",message="spec.forProvider.signing is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storage) || (has(self.initProvider) && has(self.initProvider.storage))",message="spec.forProvider.storage is a required parameter"
	Spec   AptHostedSpec   `json:"spec"`
	Status AptHostedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AptHostedList contains a list of AptHosteds
type AptHostedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AptHosted `json:"items"`
}

// Repository type metadata.
var (
	AptHosted_Kind             = "AptHosted"
	AptHosted_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AptHosted_Kind}.String()
	AptHosted_KindAPIVersion   = AptHosted_Kind + "." + CRDGroupVersion.String()
	AptHosted_GroupVersionKind = CRDGroupVersion.WithKind(AptHosted_Kind)
)

func init() {
	SchemeBuilder.Register(&AptHosted{}, &AptHostedList{})
}