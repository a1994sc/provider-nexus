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

type DockerGroupGroupInitParameters struct {

	// (List of String) Member repositories names
	// Member repositories names
	MemberNames []*string `json:"memberNames,omitempty" tf:"member_names,omitempty"`

	// only: This field is for the Group Deployment feature available in NXRM Pro.
	// Pro-only: This field is for the Group Deployment feature available in NXRM Pro.
	WritableMember *string `json:"writableMember,omitempty" tf:"writable_member,omitempty"`
}

type DockerGroupGroupObservation struct {

	// (List of String) Member repositories names
	// Member repositories names
	MemberNames []*string `json:"memberNames,omitempty" tf:"member_names,omitempty"`

	// only: This field is for the Group Deployment feature available in NXRM Pro.
	// Pro-only: This field is for the Group Deployment feature available in NXRM Pro.
	WritableMember *string `json:"writableMember,omitempty" tf:"writable_member,omitempty"`
}

type DockerGroupGroupParameters struct {

	// (List of String) Member repositories names
	// Member repositories names
	// +kubebuilder:validation:Optional
	MemberNames []*string `json:"memberNames" tf:"member_names,omitempty"`

	// only: This field is for the Group Deployment feature available in NXRM Pro.
	// Pro-only: This field is for the Group Deployment feature available in NXRM Pro.
	// +kubebuilder:validation:Optional
	WritableMember *string `json:"writableMember,omitempty" tf:"writable_member,omitempty"`
}

type DockerGroupInitParameters struct {

	// (Block List, Min: 1, Max: 1) docker contains the configuration of the docker repository (see below for nested schema)
	// docker contains the configuration of the docker repository
	Docker []DockerInitParameters `json:"docker,omitempty" tf:"docker,omitempty"`

	// (Block List, Min: 1, Max: 1) Configuration for repository group (see below for nested schema)
	// Configuration for repository group
	Group []DockerGroupGroupInitParameters `json:"group,omitempty" tf:"group,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	Storage []DockerGroupStorageInitParameters `json:"storage,omitempty" tf:"storage,omitempty"`
}

type DockerGroupObservation struct {

	// (Block List, Min: 1, Max: 1) docker contains the configuration of the docker repository (see below for nested schema)
	// docker contains the configuration of the docker repository
	Docker []DockerObservation `json:"docker,omitempty" tf:"docker,omitempty"`

	// (Block List, Min: 1, Max: 1) Configuration for repository group (see below for nested schema)
	// Configuration for repository group
	Group []DockerGroupGroupObservation `json:"group,omitempty" tf:"group,omitempty"`

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
	Storage []DockerGroupStorageObservation `json:"storage,omitempty" tf:"storage,omitempty"`
}

type DockerGroupParameters struct {

	// (Block List, Min: 1, Max: 1) docker contains the configuration of the docker repository (see below for nested schema)
	// docker contains the configuration of the docker repository
	// +kubebuilder:validation:Optional
	Docker []DockerParameters `json:"docker,omitempty" tf:"docker,omitempty"`

	// (Block List, Min: 1, Max: 1) Configuration for repository group (see below for nested schema)
	// Configuration for repository group
	// +kubebuilder:validation:Optional
	Group []DockerGroupGroupParameters `json:"group,omitempty" tf:"group,omitempty"`

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
	Storage []DockerGroupStorageParameters `json:"storage,omitempty" tf:"storage,omitempty"`
}

type DockerGroupStorageInitParameters struct {

	// (String) Blob store used to store repository contents
	// Blob store used to store repository contents
	BlobStoreName *string `json:"blobStoreName,omitempty" tf:"blob_store_name,omitempty"`

	// (Boolean) Whether to validate uploaded content's MIME type appropriate for the repository format
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation *bool `json:"strictContentTypeValidation,omitempty" tf:"strict_content_type_validation,omitempty"`
}

type DockerGroupStorageObservation struct {

	// (String) Blob store used to store repository contents
	// Blob store used to store repository contents
	BlobStoreName *string `json:"blobStoreName,omitempty" tf:"blob_store_name,omitempty"`

	// (Boolean) Whether to validate uploaded content's MIME type appropriate for the repository format
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation *bool `json:"strictContentTypeValidation,omitempty" tf:"strict_content_type_validation,omitempty"`
}

type DockerGroupStorageParameters struct {

	// (String) Blob store used to store repository contents
	// Blob store used to store repository contents
	// +kubebuilder:validation:Optional
	BlobStoreName *string `json:"blobStoreName" tf:"blob_store_name,omitempty"`

	// (Boolean) Whether to validate uploaded content's MIME type appropriate for the repository format
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	// +kubebuilder:validation:Optional
	StrictContentTypeValidation *bool `json:"strictContentTypeValidation,omitempty" tf:"strict_content_type_validation,omitempty"`
}

type DockerInitParameters struct {

	// (Boolean) Whether to force authentication (Docker Bearer Token Realm required if false)
	// Whether to force authentication (Docker Bearer Token Realm required if false)
	ForceBasicAuth *bool `json:"forceBasicAuth,omitempty" tf:"force_basic_auth,omitempty"`

	// (Number) Create an HTTP connector at specified port
	// Create an HTTP connector at specified port
	HTTPPort *float64 `json:"httpPort,omitempty" tf:"http_port,omitempty"`

	// (Number) Create an HTTPS connector at specified port
	// Create an HTTPS connector at specified port
	HTTPSPort *float64 `json:"httpsPort,omitempty" tf:"https_port,omitempty"`

	// only: Whether to allow clients to use subdomain routing connector
	// Pro-only: Whether to allow clients to use subdomain routing connector
	Subdomain *string `json:"subdomain,omitempty" tf:"subdomain,omitempty"`

	// (Boolean) Whether to allow clients to use the V1 API to interact with this repository
	// Whether to allow clients to use the V1 API to interact with this repository
	V1Enabled *bool `json:"v1Enabled,omitempty" tf:"v1_enabled,omitempty"`
}

type DockerObservation struct {

	// (Boolean) Whether to force authentication (Docker Bearer Token Realm required if false)
	// Whether to force authentication (Docker Bearer Token Realm required if false)
	ForceBasicAuth *bool `json:"forceBasicAuth,omitempty" tf:"force_basic_auth,omitempty"`

	// (Number) Create an HTTP connector at specified port
	// Create an HTTP connector at specified port
	HTTPPort *float64 `json:"httpPort,omitempty" tf:"http_port,omitempty"`

	// (Number) Create an HTTPS connector at specified port
	// Create an HTTPS connector at specified port
	HTTPSPort *float64 `json:"httpsPort,omitempty" tf:"https_port,omitempty"`

	// only: Whether to allow clients to use subdomain routing connector
	// Pro-only: Whether to allow clients to use subdomain routing connector
	Subdomain *string `json:"subdomain,omitempty" tf:"subdomain,omitempty"`

	// (Boolean) Whether to allow clients to use the V1 API to interact with this repository
	// Whether to allow clients to use the V1 API to interact with this repository
	V1Enabled *bool `json:"v1Enabled,omitempty" tf:"v1_enabled,omitempty"`
}

type DockerParameters struct {

	// (Boolean) Whether to force authentication (Docker Bearer Token Realm required if false)
	// Whether to force authentication (Docker Bearer Token Realm required if false)
	// +kubebuilder:validation:Optional
	ForceBasicAuth *bool `json:"forceBasicAuth" tf:"force_basic_auth,omitempty"`

	// (Number) Create an HTTP connector at specified port
	// Create an HTTP connector at specified port
	// +kubebuilder:validation:Optional
	HTTPPort *float64 `json:"httpPort,omitempty" tf:"http_port,omitempty"`

	// (Number) Create an HTTPS connector at specified port
	// Create an HTTPS connector at specified port
	// +kubebuilder:validation:Optional
	HTTPSPort *float64 `json:"httpsPort,omitempty" tf:"https_port,omitempty"`

	// only: Whether to allow clients to use subdomain routing connector
	// Pro-only: Whether to allow clients to use subdomain routing connector
	// +kubebuilder:validation:Optional
	Subdomain *string `json:"subdomain,omitempty" tf:"subdomain,omitempty"`

	// (Boolean) Whether to allow clients to use the V1 API to interact with this repository
	// Whether to allow clients to use the V1 API to interact with this repository
	// +kubebuilder:validation:Optional
	V1Enabled *bool `json:"v1Enabled" tf:"v1_enabled,omitempty"`
}

// DockerGroupSpec defines the desired state of DockerGroup
type DockerGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DockerGroupParameters `json:"forProvider"`
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
	InitProvider DockerGroupInitParameters `json:"initProvider,omitempty"`
}

// DockerGroupStatus defines the observed state of DockerGroup.
type DockerGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DockerGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DockerGroup is the Schema for the DockerGroups API. Use this resource to create a group docker repository.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nexus}
type DockerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.docker) || (has(self.initProvider) && has(self.initProvider.docker))",message="spec.forProvider.docker is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.group) || (has(self.initProvider) && has(self.initProvider.group))",message="spec.forProvider.group is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storage) || (has(self.initProvider) && has(self.initProvider.storage))",message="spec.forProvider.storage is a required parameter"
	Spec   DockerGroupSpec   `json:"spec"`
	Status DockerGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DockerGroupList contains a list of DockerGroups
type DockerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DockerGroup `json:"items"`
}

// Repository type metadata.
var (
	DockerGroup_Kind             = "DockerGroup"
	DockerGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DockerGroup_Kind}.String()
	DockerGroup_KindAPIVersion   = DockerGroup_Kind + "." + CRDGroupVersion.String()
	DockerGroup_GroupVersionKind = CRDGroupVersion.WithKind(DockerGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&DockerGroup{}, &DockerGroupList{})
}
