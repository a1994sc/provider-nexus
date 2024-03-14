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

type MavenHostedCleanupInitParameters struct {

	// (Set of String) List of policy names
	// List of policy names
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type MavenHostedCleanupObservation struct {

	// (Set of String) List of policy names
	// List of policy names
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type MavenHostedCleanupParameters struct {

	// (Set of String) List of policy names
	// List of policy names
	// +kubebuilder:validation:Optional
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type MavenHostedComponentInitParameters struct {

	// (Boolean) Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	ProprietaryComponents *bool `json:"proprietaryComponents,omitempty" tf:"proprietary_components,omitempty"`
}

type MavenHostedComponentObservation struct {

	// (Boolean) Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	ProprietaryComponents *bool `json:"proprietaryComponents,omitempty" tf:"proprietary_components,omitempty"`
}

type MavenHostedComponentParameters struct {

	// (Boolean) Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	// +kubebuilder:validation:Optional
	ProprietaryComponents *bool `json:"proprietaryComponents" tf:"proprietary_components,omitempty"`
}

type MavenHostedInitParameters struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	Cleanup []MavenHostedCleanupInitParameters `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Max: 1) Component configuration for the hosted repository (see below for nested schema)
	// Component configuration for the hosted repository
	Component []MavenHostedComponentInitParameters `json:"component,omitempty" tf:"component,omitempty"`

	// (Block List, Min: 1, Max: 1) Maven contains additional data of maven repository (see below for nested schema)
	// Maven contains additional data of maven repository
	Maven []MavenInitParameters `json:"maven,omitempty" tf:"maven,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	Storage []MavenHostedStorageInitParameters `json:"storage,omitempty" tf:"storage,omitempty"`
}

type MavenHostedObservation struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	Cleanup []MavenHostedCleanupObservation `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Max: 1) Component configuration for the hosted repository (see below for nested schema)
	// Component configuration for the hosted repository
	Component []MavenHostedComponentObservation `json:"component,omitempty" tf:"component,omitempty"`

	// (String) Used to identify resource at nexus
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List, Min: 1, Max: 1) Maven contains additional data of maven repository (see below for nested schema)
	// Maven contains additional data of maven repository
	Maven []MavenObservation `json:"maven,omitempty" tf:"maven,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	Storage []MavenHostedStorageObservation `json:"storage,omitempty" tf:"storage,omitempty"`
}

type MavenHostedParameters struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	// +kubebuilder:validation:Optional
	Cleanup []MavenHostedCleanupParameters `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Max: 1) Component configuration for the hosted repository (see below for nested schema)
	// Component configuration for the hosted repository
	// +kubebuilder:validation:Optional
	Component []MavenHostedComponentParameters `json:"component,omitempty" tf:"component,omitempty"`

	// (Block List, Min: 1, Max: 1) Maven contains additional data of maven repository (see below for nested schema)
	// Maven contains additional data of maven repository
	// +kubebuilder:validation:Optional
	Maven []MavenParameters `json:"maven,omitempty" tf:"maven,omitempty"`

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
	Storage []MavenHostedStorageParameters `json:"storage,omitempty" tf:"storage,omitempty"`
}

type MavenHostedStorageInitParameters struct {

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

type MavenHostedStorageObservation struct {

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

type MavenHostedStorageParameters struct {

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

type MavenInitParameters struct {

	// Disposition header as 'Attachment' to disable some content from being inline in a browse. Possible Value: INLINE or ATTACHMENT
	// Add Content-Disposition header as 'Attachment' to disable some content from being inline in a browse. Possible Value: `INLINE` or `ATTACHMENT`
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// (String) Validate that all paths are maven artifact or metadata paths. Possible Value: STRICT or PERMISSIVE
	// Validate that all paths are maven artifact or metadata paths. Possible Value: `STRICT` or `PERMISSIVE`
	LayoutPolicy *string `json:"layoutPolicy,omitempty" tf:"layout_policy,omitempty"`

	// (String) What type of artifacts does this repository store? Possible Value: RELEASE, SNAPSHOT or MIXED
	// What type of artifacts does this repository store? Possible Value: `RELEASE`, `SNAPSHOT` or `MIXED`
	VersionPolicy *string `json:"versionPolicy,omitempty" tf:"version_policy,omitempty"`
}

type MavenObservation struct {

	// Disposition header as 'Attachment' to disable some content from being inline in a browse. Possible Value: INLINE or ATTACHMENT
	// Add Content-Disposition header as 'Attachment' to disable some content from being inline in a browse. Possible Value: `INLINE` or `ATTACHMENT`
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// (String) Validate that all paths are maven artifact or metadata paths. Possible Value: STRICT or PERMISSIVE
	// Validate that all paths are maven artifact or metadata paths. Possible Value: `STRICT` or `PERMISSIVE`
	LayoutPolicy *string `json:"layoutPolicy,omitempty" tf:"layout_policy,omitempty"`

	// (String) What type of artifacts does this repository store? Possible Value: RELEASE, SNAPSHOT or MIXED
	// What type of artifacts does this repository store? Possible Value: `RELEASE`, `SNAPSHOT` or `MIXED`
	VersionPolicy *string `json:"versionPolicy,omitempty" tf:"version_policy,omitempty"`
}

type MavenParameters struct {

	// Disposition header as 'Attachment' to disable some content from being inline in a browse. Possible Value: INLINE or ATTACHMENT
	// Add Content-Disposition header as 'Attachment' to disable some content from being inline in a browse. Possible Value: `INLINE` or `ATTACHMENT`
	// +kubebuilder:validation:Optional
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// (String) Validate that all paths are maven artifact or metadata paths. Possible Value: STRICT or PERMISSIVE
	// Validate that all paths are maven artifact or metadata paths. Possible Value: `STRICT` or `PERMISSIVE`
	// +kubebuilder:validation:Optional
	LayoutPolicy *string `json:"layoutPolicy" tf:"layout_policy,omitempty"`

	// (String) What type of artifacts does this repository store? Possible Value: RELEASE, SNAPSHOT or MIXED
	// What type of artifacts does this repository store? Possible Value: `RELEASE`, `SNAPSHOT` or `MIXED`
	// +kubebuilder:validation:Optional
	VersionPolicy *string `json:"versionPolicy" tf:"version_policy,omitempty"`
}

// MavenHostedSpec defines the desired state of MavenHosted
type MavenHostedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MavenHostedParameters `json:"forProvider"`
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
	InitProvider MavenHostedInitParameters `json:"initProvider,omitempty"`
}

// MavenHostedStatus defines the observed state of MavenHosted.
type MavenHostedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MavenHostedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MavenHosted is the Schema for the MavenHosteds API. Use this resource to create a hosted maven repository.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nexus}
type MavenHosted struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.maven) || (has(self.initProvider) && has(self.initProvider.maven))",message="spec.forProvider.maven is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storage) || (has(self.initProvider) && has(self.initProvider.storage))",message="spec.forProvider.storage is a required parameter"
	Spec   MavenHostedSpec   `json:"spec"`
	Status MavenHostedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MavenHostedList contains a list of MavenHosteds
type MavenHostedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MavenHosted `json:"items"`
}

// Repository type metadata.
var (
	MavenHosted_Kind             = "MavenHosted"
	MavenHosted_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MavenHosted_Kind}.String()
	MavenHosted_KindAPIVersion   = MavenHosted_Kind + "." + CRDGroupVersion.String()
	MavenHosted_GroupVersionKind = CRDGroupVersion.WithKind(MavenHosted_Kind)
)

func init() {
	SchemeBuilder.Register(&MavenHosted{}, &MavenHostedList{})
}
