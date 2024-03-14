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

type AdvancedBucketConnectionInitParameters struct {

	// (String) A custom endpoint URL for third party object stores using the S3 API.
	// A custom endpoint URL for third party object stores using the S3 API.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// style access being used for all requests.
	// Setting this flag will result in path-style access being used for all requests.
	ForcePathStyle *bool `json:"forcePathStyle,omitempty" tf:"force_path_style,omitempty"`

	// (Number) Setting this value will override the default connection pool size of Nexus of the s3 client for this blobstore.
	// Setting this value will override the default connection pool size of Nexus of the s3 client for this blobstore.
	MaxConnectionPoolSize *float64 `json:"maxConnectionPoolSize,omitempty" tf:"max_connection_pool_size,omitempty"`

	// (String) An API signature version which may be required for third party object stores using the S3 API.
	// An API signature version which may be required for third party object stores using the S3 API.
	SignerType *string `json:"signerType,omitempty" tf:"signer_type,omitempty"`
}

type AdvancedBucketConnectionObservation struct {

	// (String) A custom endpoint URL for third party object stores using the S3 API.
	// A custom endpoint URL for third party object stores using the S3 API.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// style access being used for all requests.
	// Setting this flag will result in path-style access being used for all requests.
	ForcePathStyle *bool `json:"forcePathStyle,omitempty" tf:"force_path_style,omitempty"`

	// (Number) Setting this value will override the default connection pool size of Nexus of the s3 client for this blobstore.
	// Setting this value will override the default connection pool size of Nexus of the s3 client for this blobstore.
	MaxConnectionPoolSize *float64 `json:"maxConnectionPoolSize,omitempty" tf:"max_connection_pool_size,omitempty"`

	// (String) An API signature version which may be required for third party object stores using the S3 API.
	// An API signature version which may be required for third party object stores using the S3 API.
	SignerType *string `json:"signerType,omitempty" tf:"signer_type,omitempty"`
}

type AdvancedBucketConnectionParameters struct {

	// (String) A custom endpoint URL for third party object stores using the S3 API.
	// A custom endpoint URL for third party object stores using the S3 API.
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// style access being used for all requests.
	// Setting this flag will result in path-style access being used for all requests.
	// +kubebuilder:validation:Optional
	ForcePathStyle *bool `json:"forcePathStyle,omitempty" tf:"force_path_style,omitempty"`

	// (Number) Setting this value will override the default connection pool size of Nexus of the s3 client for this blobstore.
	// Setting this value will override the default connection pool size of Nexus of the s3 client for this blobstore.
	// +kubebuilder:validation:Optional
	MaxConnectionPoolSize *float64 `json:"maxConnectionPoolSize,omitempty" tf:"max_connection_pool_size,omitempty"`

	// (String) An API signature version which may be required for third party object stores using the S3 API.
	// An API signature version which may be required for third party object stores using the S3 API.
	// +kubebuilder:validation:Optional
	SignerType *string `json:"signerType,omitempty" tf:"signer_type,omitempty"`
}

type BucketInitParameters struct {

	// 1 to disable)
	// How many days until deleted blobs are finally removed from the S3 bucket (-1 to disable)
	Expiration *float64 `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// (String) Blobstore name
	// The name of the S3 bucket
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The S3 blob store (i.e S3 object) key prefix
	// The S3 blob store (i.e S3 object) key prefix
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// (String) The AWS region to create a new S3 bucket in or an existing S3 bucket's region
	// The AWS region to create a new S3 bucket in or an existing S3 bucket's region
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type BucketObservation struct {

	// 1 to disable)
	// How many days until deleted blobs are finally removed from the S3 bucket (-1 to disable)
	Expiration *float64 `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// (String) Blobstore name
	// The name of the S3 bucket
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The S3 blob store (i.e S3 object) key prefix
	// The S3 blob store (i.e S3 object) key prefix
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// (String) The AWS region to create a new S3 bucket in or an existing S3 bucket's region
	// The AWS region to create a new S3 bucket in or an existing S3 bucket's region
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type BucketParameters struct {

	// 1 to disable)
	// How many days until deleted blobs are finally removed from the S3 bucket (-1 to disable)
	// +kubebuilder:validation:Optional
	Expiration *float64 `json:"expiration" tf:"expiration,omitempty"`

	// (String) Blobstore name
	// The name of the S3 bucket
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) The S3 blob store (i.e S3 object) key prefix
	// The S3 blob store (i.e S3 object) key prefix
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// (String) The AWS region to create a new S3 bucket in or an existing S3 bucket's region
	// The AWS region to create a new S3 bucket in or an existing S3 bucket's region
	// +kubebuilder:validation:Optional
	Region *string `json:"region" tf:"region,omitempty"`
}

type BucketSecurityInitParameters struct {

	// (String) An IAM role to assume in order to access the S3 bucket
	// An IAM role to assume in order to access the S3 bucket
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type BucketSecurityObservation struct {

	// (String) An IAM role to assume in order to access the S3 bucket
	// An IAM role to assume in order to access the S3 bucket
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type BucketSecurityParameters struct {

	// (String) An IAM access key ID for granting access to the S3 bucket
	// An IAM access key ID for granting access to the S3 bucket
	// +kubebuilder:validation:Optional
	AccessKeyIDSecretRef *v1.SecretKeySelector `json:"accessKeyIdSecretRef,omitempty" tf:"-"`

	// (String) An IAM role to assume in order to access the S3 bucket
	// An IAM role to assume in order to access the S3 bucket
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// (String, Sensitive) The secret access key associated with the specified IAM access key ID
	// The secret access key associated with the specified IAM access key ID
	// +kubebuilder:validation:Optional
	SecretAccessKeySecretRef *v1.SecretKeySelector `json:"secretAccessKeySecretRef,omitempty" tf:"-"`

	// (String, Sensitive) An AWS STS session token associated with temporary security credentials which grant access to the S3 bucket
	// An AWS STS session token associated with temporary security credentials which grant access to the S3 bucket
	// +kubebuilder:validation:Optional
	SessionTokenSecretRef *v1.SecretKeySelector `json:"sessionTokenSecretRef,omitempty" tf:"-"`
}

type EncryptionInitParameters struct {

	// (String) The encryption key.
	// The encryption key.
	EncryptionKey *string `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`

	// (String) The type of S3 server side encryption to use.
	// The type of S3 server side encryption to use.
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`
}

type EncryptionObservation struct {

	// (String) The encryption key.
	// The encryption key.
	EncryptionKey *string `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`

	// (String) The type of S3 server side encryption to use.
	// The type of S3 server side encryption to use.
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`
}

type EncryptionParameters struct {

	// (String) The encryption key.
	// The encryption key.
	// +kubebuilder:validation:Optional
	EncryptionKey *string `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`

	// (String) The type of S3 server side encryption to use.
	// The type of S3 server side encryption to use.
	// +kubebuilder:validation:Optional
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`
}

type S3BucketConfigurationInitParameters struct {

	// (Block List, Max: 1) Additional connection configurations (see below for nested schema)
	// Additional connection configurations
	AdvancedBucketConnection []AdvancedBucketConnectionInitParameters `json:"advancedBucketConnection,omitempty" tf:"advanced_bucket_connection,omitempty"`

	// (Block List, Min: 1, Max: 1) The S3 bucket configuration (see below for nested schema)
	// The S3 bucket configuration
	Bucket []BucketInitParameters `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// (Block List, Max: 1) Additional security configurations (see below for nested schema)
	// Additional security configurations
	BucketSecurity []BucketSecurityInitParameters `json:"bucketSecurity,omitempty" tf:"bucket_security,omitempty"`

	// (Block List, Max: 1) Additional bucket encryption configurations (see below for nested schema)
	// Additional bucket encryption configurations
	Encryption []EncryptionInitParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`
}

type S3BucketConfigurationObservation struct {

	// (Block List, Max: 1) Additional connection configurations (see below for nested schema)
	// Additional connection configurations
	AdvancedBucketConnection []AdvancedBucketConnectionObservation `json:"advancedBucketConnection,omitempty" tf:"advanced_bucket_connection,omitempty"`

	// (Block List, Min: 1, Max: 1) The S3 bucket configuration (see below for nested schema)
	// The S3 bucket configuration
	Bucket []BucketObservation `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// (Block List, Max: 1) Additional security configurations (see below for nested schema)
	// Additional security configurations
	BucketSecurity []BucketSecurityObservation `json:"bucketSecurity,omitempty" tf:"bucket_security,omitempty"`

	// (Block List, Max: 1) Additional bucket encryption configurations (see below for nested schema)
	// Additional bucket encryption configurations
	Encryption []EncryptionObservation `json:"encryption,omitempty" tf:"encryption,omitempty"`
}

type S3BucketConfigurationParameters struct {

	// (Block List, Max: 1) Additional connection configurations (see below for nested schema)
	// Additional connection configurations
	// +kubebuilder:validation:Optional
	AdvancedBucketConnection []AdvancedBucketConnectionParameters `json:"advancedBucketConnection,omitempty" tf:"advanced_bucket_connection,omitempty"`

	// (Block List, Min: 1, Max: 1) The S3 bucket configuration (see below for nested schema)
	// The S3 bucket configuration
	// +kubebuilder:validation:Optional
	Bucket []BucketParameters `json:"bucket" tf:"bucket,omitempty"`

	// (Block List, Max: 1) Additional security configurations (see below for nested schema)
	// Additional security configurations
	// +kubebuilder:validation:Optional
	BucketSecurity []BucketSecurityParameters `json:"bucketSecurity,omitempty" tf:"bucket_security,omitempty"`

	// (Block List, Max: 1) Additional bucket encryption configurations (see below for nested schema)
	// Additional bucket encryption configurations
	// +kubebuilder:validation:Optional
	Encryption []EncryptionParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`
}

type S3InitParameters struct {

	// (Block List, Min: 1, Max: 1) The S3 bucket configuration. (see below for nested schema)
	// The S3 bucket configuration.
	BucketConfiguration []S3BucketConfigurationInitParameters `json:"bucketConfiguration,omitempty" tf:"bucket_configuration,omitempty"`

	// (String) Blobstore name
	// Blobstore name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) Soft quota of the blobstore (see below for nested schema)
	// Soft quota of the blobstore
	SoftQuota []S3SoftQuotaInitParameters `json:"softQuota,omitempty" tf:"soft_quota,omitempty"`
}

type S3Observation struct {

	// (Number) Count of blobs
	// Count of blobs
	BlobCount *float64 `json:"blobCount,omitempty" tf:"blob_count,omitempty"`

	// (Block List, Min: 1, Max: 1) The S3 bucket configuration. (see below for nested schema)
	// The S3 bucket configuration.
	BucketConfiguration []S3BucketConfigurationObservation `json:"bucketConfiguration,omitempty" tf:"bucket_configuration,omitempty"`

	// (String) Used to identify resource at nexus
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Blobstore name
	// Blobstore name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) Soft quota of the blobstore (see below for nested schema)
	// Soft quota of the blobstore
	SoftQuota []S3SoftQuotaObservation `json:"softQuota,omitempty" tf:"soft_quota,omitempty"`

	// (Number) The total size of the blobstore in Bytes
	// The total size of the blobstore in Bytes
	TotalSizeInBytes *float64 `json:"totalSizeInBytes,omitempty" tf:"total_size_in_bytes,omitempty"`
}

type S3Parameters struct {

	// (Block List, Min: 1, Max: 1) The S3 bucket configuration. (see below for nested schema)
	// The S3 bucket configuration.
	// +kubebuilder:validation:Optional
	BucketConfiguration []S3BucketConfigurationParameters `json:"bucketConfiguration,omitempty" tf:"bucket_configuration,omitempty"`

	// (String) Blobstore name
	// Blobstore name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) Soft quota of the blobstore (see below for nested schema)
	// Soft quota of the blobstore
	// +kubebuilder:validation:Optional
	SoftQuota []S3SoftQuotaParameters `json:"softQuota,omitempty" tf:"soft_quota,omitempty"`
}

type S3SoftQuotaInitParameters struct {

	// (Number) The limit in Bytes. Minimum value is 1000000
	// The limit in Bytes. Minimum value is 1000000
	Limit *float64 `json:"limit,omitempty" tf:"limit,omitempty"`

	// (String) The type to use such as spaceRemainingQuota, or spaceUsedQuota
	// The type to use such as spaceRemainingQuota, or spaceUsedQuota
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type S3SoftQuotaObservation struct {

	// (Number) The limit in Bytes. Minimum value is 1000000
	// The limit in Bytes. Minimum value is 1000000
	Limit *float64 `json:"limit,omitempty" tf:"limit,omitempty"`

	// (String) The type to use such as spaceRemainingQuota, or spaceUsedQuota
	// The type to use such as spaceRemainingQuota, or spaceUsedQuota
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type S3SoftQuotaParameters struct {

	// (Number) The limit in Bytes. Minimum value is 1000000
	// The limit in Bytes. Minimum value is 1000000
	// +kubebuilder:validation:Optional
	Limit *float64 `json:"limit" tf:"limit,omitempty"`

	// (String) The type to use such as spaceRemainingQuota, or spaceUsedQuota
	// The type to use such as spaceRemainingQuota, or spaceUsedQuota
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// S3Spec defines the desired state of S3
type S3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     S3Parameters `json:"forProvider"`
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
	InitProvider S3InitParameters `json:"initProvider,omitempty"`
}

// S3Status defines the observed state of S3.
type S3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        S3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// S3 is the Schema for the S3s API. Use this resource to create a Nexus S3 blobstore.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nexus}
type S3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bucketConfiguration) || (has(self.initProvider) && has(self.initProvider.bucketConfiguration))",message="spec.forProvider.bucketConfiguration is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   S3Spec   `json:"spec"`
	Status S3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3List contains a list of S3s
type S3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3 `json:"items"`
}

// Repository type metadata.
var (
	S3_Kind             = "S3"
	S3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: S3_Kind}.String()
	S3_KindAPIVersion   = S3_Kind + "." + CRDGroupVersion.String()
	S3_GroupVersionKind = CRDGroupVersion.WithKind(S3_Kind)
)

func init() {
	SchemeBuilder.Register(&S3{}, &S3List{})
}
