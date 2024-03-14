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

type FileInitParameters struct {

	// (String) Blobstore name
	// Blobstore name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// work directory
	// The path to the blobstore contents. This can be an absolute path to anywhere on the system nxrm has access to or it can be a path relative to the sonatype-work directory
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (Block List, Max: 1) Soft quota of the blobstore (see below for nested schema)
	// Soft quota of the blobstore
	SoftQuota []FileSoftQuotaInitParameters `json:"softQuota,omitempty" tf:"soft_quota,omitempty"`
}

type FileObservation struct {

	// (Number) Available space in Bytes
	// Available space in Bytes
	AvailableSpaceInBytes *float64 `json:"availableSpaceInBytes,omitempty" tf:"available_space_in_bytes,omitempty"`

	// (Number) Count of blobs
	// Count of blobs
	BlobCount *float64 `json:"blobCount,omitempty" tf:"blob_count,omitempty"`

	// (String) Used to identify resource at nexus
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Blobstore name
	// Blobstore name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// work directory
	// The path to the blobstore contents. This can be an absolute path to anywhere on the system nxrm has access to or it can be a path relative to the sonatype-work directory
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (Block List, Max: 1) Soft quota of the blobstore (see below for nested schema)
	// Soft quota of the blobstore
	SoftQuota []FileSoftQuotaObservation `json:"softQuota,omitempty" tf:"soft_quota,omitempty"`

	// (Number) The total size of the blobstore in Bytes
	// The total size of the blobstore in Bytes
	TotalSizeInBytes *float64 `json:"totalSizeInBytes,omitempty" tf:"total_size_in_bytes,omitempty"`
}

type FileParameters struct {

	// (String) Blobstore name
	// Blobstore name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// work directory
	// The path to the blobstore contents. This can be an absolute path to anywhere on the system nxrm has access to or it can be a path relative to the sonatype-work directory
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (Block List, Max: 1) Soft quota of the blobstore (see below for nested schema)
	// Soft quota of the blobstore
	// +kubebuilder:validation:Optional
	SoftQuota []FileSoftQuotaParameters `json:"softQuota,omitempty" tf:"soft_quota,omitempty"`
}

type FileSoftQuotaInitParameters struct {

	// (Number) The limit in Bytes. Minimum value is 1000000
	// The limit in Bytes. Minimum value is 1000000
	Limit *float64 `json:"limit,omitempty" tf:"limit,omitempty"`

	// (String) The type to use such as spaceRemainingQuota, or spaceUsedQuota
	// The type to use such as spaceRemainingQuota, or spaceUsedQuota
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FileSoftQuotaObservation struct {

	// (Number) The limit in Bytes. Minimum value is 1000000
	// The limit in Bytes. Minimum value is 1000000
	Limit *float64 `json:"limit,omitempty" tf:"limit,omitempty"`

	// (String) The type to use such as spaceRemainingQuota, or spaceUsedQuota
	// The type to use such as spaceRemainingQuota, or spaceUsedQuota
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FileSoftQuotaParameters struct {

	// (Number) The limit in Bytes. Minimum value is 1000000
	// The limit in Bytes. Minimum value is 1000000
	// +kubebuilder:validation:Optional
	Limit *float64 `json:"limit" tf:"limit,omitempty"`

	// (String) The type to use such as spaceRemainingQuota, or spaceUsedQuota
	// The type to use such as spaceRemainingQuota, or spaceUsedQuota
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// FileSpec defines the desired state of File
type FileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FileParameters `json:"forProvider"`
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
	InitProvider FileInitParameters `json:"initProvider,omitempty"`
}

// FileStatus defines the observed state of File.
type FileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// File is the Schema for the Files API. Use this resource to create a Nexus file blobstore.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nexus}
type File struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   FileSpec   `json:"spec"`
	Status FileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FileList contains a list of Files
type FileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []File `json:"items"`
}

// Repository type metadata.
var (
	File_Kind             = "File"
	File_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: File_Kind}.String()
	File_KindAPIVersion   = File_Kind + "." + CRDGroupVersion.String()
	File_GroupVersionKind = CRDGroupVersion.WithKind(File_Kind)
)

func init() {
	SchemeBuilder.Register(&File{}, &FileList{})
}
