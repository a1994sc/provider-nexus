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

type GitlfsHostedCleanupInitParameters struct {

	// (Set of String) List of policy names
	// List of policy names
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type GitlfsHostedCleanupObservation struct {

	// (Set of String) List of policy names
	// List of policy names
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type GitlfsHostedCleanupParameters struct {

	// (Set of String) List of policy names
	// List of policy names
	// +kubebuilder:validation:Optional
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type GitlfsHostedComponentInitParameters struct {

	// (Boolean) Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	ProprietaryComponents *bool `json:"proprietaryComponents,omitempty" tf:"proprietary_components,omitempty"`
}

type GitlfsHostedComponentObservation struct {

	// (Boolean) Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	ProprietaryComponents *bool `json:"proprietaryComponents,omitempty" tf:"proprietary_components,omitempty"`
}

type GitlfsHostedComponentParameters struct {

	// (Boolean) Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// +kubebuilder:validation:Optional
	ProprietaryComponents *bool `json:"proprietaryComponents" tf:"proprietary_components,omitempty"`
}

type GitlfsHostedInitParameters struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	Cleanup []GitlfsHostedCleanupInitParameters `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Max: 1) Component configuration for the hosted repository (see below for nested schema)
	// Component configuration for the hosted repository
	Component []GitlfsHostedComponentInitParameters `json:"component,omitempty" tf:"component,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	Storage []GitlfsHostedStorageInitParameters `json:"storage,omitempty" tf:"storage,omitempty"`
}

type GitlfsHostedObservation struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	Cleanup []GitlfsHostedCleanupObservation `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Max: 1) Component configuration for the hosted repository (see below for nested schema)
	// Component configuration for the hosted repository
	Component []GitlfsHostedComponentObservation `json:"component,omitempty" tf:"component,omitempty"`

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
	Storage []GitlfsHostedStorageObservation `json:"storage,omitempty" tf:"storage,omitempty"`
}

type GitlfsHostedParameters struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	// +kubebuilder:validation:Optional
	Cleanup []GitlfsHostedCleanupParameters `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Max: 1) Component configuration for the hosted repository (see below for nested schema)
	// Component configuration for the hosted repository
	// +kubebuilder:validation:Optional
	Component []GitlfsHostedComponentParameters `json:"component,omitempty" tf:"component,omitempty"`

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
	Storage []GitlfsHostedStorageParameters `json:"storage,omitempty" tf:"storage,omitempty"`
}

type GitlfsHostedStorageInitParameters struct {

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

type GitlfsHostedStorageObservation struct {

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

type GitlfsHostedStorageParameters struct {

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

// GitlfsHostedSpec defines the desired state of GitlfsHosted
type GitlfsHostedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GitlfsHostedParameters `json:"forProvider"`
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
	InitProvider GitlfsHostedInitParameters `json:"initProvider,omitempty"`
}

// GitlfsHostedStatus defines the observed state of GitlfsHosted.
type GitlfsHostedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GitlfsHostedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GitlfsHosted is the Schema for the GitlfsHosteds API. Use this resource to create a hosted gitlfs repository.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nexus}
type GitlfsHosted struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storage) || (has(self.initProvider) && has(self.initProvider.storage))",message="spec.forProvider.storage is a required parameter"
	Spec   GitlfsHostedSpec   `json:"spec"`
	Status GitlfsHostedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GitlfsHostedList contains a list of GitlfsHosteds
type GitlfsHostedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitlfsHosted `json:"items"`
}

// Repository type metadata.
var (
	GitlfsHosted_Kind             = "GitlfsHosted"
	GitlfsHosted_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GitlfsHosted_Kind}.String()
	GitlfsHosted_KindAPIVersion   = GitlfsHosted_Kind + "." + CRDGroupVersion.String()
	GitlfsHosted_GroupVersionKind = CRDGroupVersion.WithKind(GitlfsHosted_Kind)
)

func init() {
	SchemeBuilder.Register(&GitlfsHosted{}, &GitlfsHostedList{})
}
