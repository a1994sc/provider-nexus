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

type NugetProxyCleanupInitParameters struct {

	// (Set of String) List of policy names
	// List of policy names
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type NugetProxyCleanupObservation struct {

	// (Set of String) List of policy names
	// List of policy names
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type NugetProxyCleanupParameters struct {

	// (Set of String) List of policy names
	// List of policy names
	// +kubebuilder:validation:Optional
	// +listType=set
	PolicyNames []*string `json:"policyNames,omitempty" tf:"policy_names,omitempty"`
}

type NugetProxyHTTPClientAuthenticationInitParameters struct {

	// (String) The ntlm domain to connect
	// The ntlm domain to connect
	NtlmDomain *string `json:"ntlmDomain,omitempty" tf:"ntlm_domain,omitempty"`

	// (String) The ntlm host to connect
	// The ntlm host to connect
	NtlmHost *string `json:"ntlmHost,omitempty" tf:"ntlm_host,omitempty"`

	// (String, Sensitive) The password used by the proxy repository
	// The password used by the proxy repository
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// (String) Authentication type. Possible values: ntlm or username
	// Authentication type. Possible values: `ntlm` or `username`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) The username used by the proxy repository
	// The username used by the proxy repository
	UsernameSecretRef *v1.SecretKeySelector `json:"usernameSecretRef,omitempty" tf:"-"`
}

type NugetProxyHTTPClientAuthenticationObservation struct {

	// (String) The ntlm domain to connect
	// The ntlm domain to connect
	NtlmDomain *string `json:"ntlmDomain,omitempty" tf:"ntlm_domain,omitempty"`

	// (String) The ntlm host to connect
	// The ntlm host to connect
	NtlmHost *string `json:"ntlmHost,omitempty" tf:"ntlm_host,omitempty"`

	// (String) Authentication type. Possible values: ntlm or username
	// Authentication type. Possible values: `ntlm` or `username`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NugetProxyHTTPClientAuthenticationParameters struct {

	// (String) The ntlm domain to connect
	// The ntlm domain to connect
	// +kubebuilder:validation:Optional
	NtlmDomain *string `json:"ntlmDomain,omitempty" tf:"ntlm_domain,omitempty"`

	// (String) The ntlm host to connect
	// The ntlm host to connect
	// +kubebuilder:validation:Optional
	NtlmHost *string `json:"ntlmHost,omitempty" tf:"ntlm_host,omitempty"`

	// (String, Sensitive) The password used by the proxy repository
	// The password used by the proxy repository
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// (String) Authentication type. Possible values: ntlm or username
	// Authentication type. Possible values: `ntlm` or `username`
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// (String) The username used by the proxy repository
	// The username used by the proxy repository
	// +kubebuilder:validation:Optional
	UsernameSecretRef *v1.SecretKeySelector `json:"usernameSecretRef,omitempty" tf:"-"`
}

type NugetProxyHTTPClientConnectionInitParameters struct {

	// (Boolean) Whether to enable redirects to the same location (may be required by some servers)
	// Whether to enable redirects to the same location (may be required by some servers)
	EnableCircularRedirects *bool `json:"enableCircularRedirects,omitempty" tf:"enable_circular_redirects,omitempty"`

	// (Boolean) Whether to allow cookies to be stored and used
	// Whether to allow cookies to be stored and used
	EnableCookies *bool `json:"enableCookies,omitempty" tf:"enable_cookies,omitempty"`

	// (Number) Total retries if the initial connection attempt suffers a timeout
	// Total retries if the initial connection attempt suffers a timeout
	Retries *float64 `json:"retries,omitempty" tf:"retries,omitempty"`

	// (Number) Seconds to wait for activity before stopping and retrying the connection
	// Seconds to wait for activity before stopping and retrying the connection
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// (Boolean) Use certificates stored in the Nexus Repository Manager truststore to connect to external systems
	// Use certificates stored in the Nexus Repository Manager truststore to connect to external systems
	UseTrustStore *bool `json:"useTrustStore,omitempty" tf:"use_trust_store,omitempty"`

	// Agent header in HTTP requests
	// Custom fragment to append to User-Agent header in HTTP requests
	UserAgentSuffix *string `json:"userAgentSuffix,omitempty" tf:"user_agent_suffix,omitempty"`
}

type NugetProxyHTTPClientConnectionObservation struct {

	// (Boolean) Whether to enable redirects to the same location (may be required by some servers)
	// Whether to enable redirects to the same location (may be required by some servers)
	EnableCircularRedirects *bool `json:"enableCircularRedirects,omitempty" tf:"enable_circular_redirects,omitempty"`

	// (Boolean) Whether to allow cookies to be stored and used
	// Whether to allow cookies to be stored and used
	EnableCookies *bool `json:"enableCookies,omitempty" tf:"enable_cookies,omitempty"`

	// (Number) Total retries if the initial connection attempt suffers a timeout
	// Total retries if the initial connection attempt suffers a timeout
	Retries *float64 `json:"retries,omitempty" tf:"retries,omitempty"`

	// (Number) Seconds to wait for activity before stopping and retrying the connection
	// Seconds to wait for activity before stopping and retrying the connection
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// (Boolean) Use certificates stored in the Nexus Repository Manager truststore to connect to external systems
	// Use certificates stored in the Nexus Repository Manager truststore to connect to external systems
	UseTrustStore *bool `json:"useTrustStore,omitempty" tf:"use_trust_store,omitempty"`

	// Agent header in HTTP requests
	// Custom fragment to append to User-Agent header in HTTP requests
	UserAgentSuffix *string `json:"userAgentSuffix,omitempty" tf:"user_agent_suffix,omitempty"`
}

type NugetProxyHTTPClientConnectionParameters struct {

	// (Boolean) Whether to enable redirects to the same location (may be required by some servers)
	// Whether to enable redirects to the same location (may be required by some servers)
	// +kubebuilder:validation:Optional
	EnableCircularRedirects *bool `json:"enableCircularRedirects,omitempty" tf:"enable_circular_redirects,omitempty"`

	// (Boolean) Whether to allow cookies to be stored and used
	// Whether to allow cookies to be stored and used
	// +kubebuilder:validation:Optional
	EnableCookies *bool `json:"enableCookies,omitempty" tf:"enable_cookies,omitempty"`

	// (Number) Total retries if the initial connection attempt suffers a timeout
	// Total retries if the initial connection attempt suffers a timeout
	// +kubebuilder:validation:Optional
	Retries *float64 `json:"retries,omitempty" tf:"retries,omitempty"`

	// (Number) Seconds to wait for activity before stopping and retrying the connection
	// Seconds to wait for activity before stopping and retrying the connection
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// (Boolean) Use certificates stored in the Nexus Repository Manager truststore to connect to external systems
	// Use certificates stored in the Nexus Repository Manager truststore to connect to external systems
	// +kubebuilder:validation:Optional
	UseTrustStore *bool `json:"useTrustStore,omitempty" tf:"use_trust_store,omitempty"`

	// Agent header in HTTP requests
	// Custom fragment to append to User-Agent header in HTTP requests
	// +kubebuilder:validation:Optional
	UserAgentSuffix *string `json:"userAgentSuffix,omitempty" tf:"user_agent_suffix,omitempty"`
}

type NugetProxyHTTPClientInitParameters struct {

	// (Block List, Max: 1) Authentication configuration of the HTTP client (see below for nested schema)
	// Authentication configuration of the HTTP client
	Authentication []NugetProxyHTTPClientAuthenticationInitParameters `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// block outbound connections if remote peer is detected as unreachable/unresponsive
	// Whether to auto-block outbound connections if remote peer is detected as unreachable/unresponsive
	AutoBlock *bool `json:"autoBlock,omitempty" tf:"auto_block,omitempty"`

	// (Boolean) Whether to block outbound connections on the repository
	// Whether to block outbound connections on the repository
	Blocked *bool `json:"blocked,omitempty" tf:"blocked,omitempty"`

	// (Block List, Max: 1) Connection configuration of the HTTP client (see below for nested schema)
	// Connection configuration of the HTTP client
	Connection []NugetProxyHTTPClientConnectionInitParameters `json:"connection,omitempty" tf:"connection,omitempty"`
}

type NugetProxyHTTPClientObservation struct {

	// (Block List, Max: 1) Authentication configuration of the HTTP client (see below for nested schema)
	// Authentication configuration of the HTTP client
	Authentication []NugetProxyHTTPClientAuthenticationObservation `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// block outbound connections if remote peer is detected as unreachable/unresponsive
	// Whether to auto-block outbound connections if remote peer is detected as unreachable/unresponsive
	AutoBlock *bool `json:"autoBlock,omitempty" tf:"auto_block,omitempty"`

	// (Boolean) Whether to block outbound connections on the repository
	// Whether to block outbound connections on the repository
	Blocked *bool `json:"blocked,omitempty" tf:"blocked,omitempty"`

	// (Block List, Max: 1) Connection configuration of the HTTP client (see below for nested schema)
	// Connection configuration of the HTTP client
	Connection []NugetProxyHTTPClientConnectionObservation `json:"connection,omitempty" tf:"connection,omitempty"`
}

type NugetProxyHTTPClientParameters struct {

	// (Block List, Max: 1) Authentication configuration of the HTTP client (see below for nested schema)
	// Authentication configuration of the HTTP client
	// +kubebuilder:validation:Optional
	Authentication []NugetProxyHTTPClientAuthenticationParameters `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// block outbound connections if remote peer is detected as unreachable/unresponsive
	// Whether to auto-block outbound connections if remote peer is detected as unreachable/unresponsive
	// +kubebuilder:validation:Optional
	AutoBlock *bool `json:"autoBlock" tf:"auto_block,omitempty"`

	// (Boolean) Whether to block outbound connections on the repository
	// Whether to block outbound connections on the repository
	// +kubebuilder:validation:Optional
	Blocked *bool `json:"blocked" tf:"blocked,omitempty"`

	// (Block List, Max: 1) Connection configuration of the HTTP client (see below for nested schema)
	// Connection configuration of the HTTP client
	// +kubebuilder:validation:Optional
	Connection []NugetProxyHTTPClientConnectionParameters `json:"connection,omitempty" tf:"connection,omitempty"`
}

type NugetProxyInitParameters struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	Cleanup []NugetProxyCleanupInitParameters `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Min: 1, Max: 1) HTTP Client configuration for proxy repositories (see below for nested schema)
	// HTTP Client configuration for proxy repositories
	HTTPClient []NugetProxyHTTPClientInitParameters `json:"httpClient,omitempty" tf:"http_client,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Min: 1, Max: 1) Configuration of the negative cache handling (see below for nested schema)
	// Configuration of the negative cache handling
	NegativeCache []NugetProxyNegativeCacheInitParameters `json:"negativeCache,omitempty" tf:"negative_cache,omitempty"`

	// (String) Nuget protocol version
	// Nuget protocol version
	NugetVersion *string `json:"nugetVersion,omitempty" tf:"nuget_version,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// (Block List, Min: 1, Max: 1) Configuration for the proxy repository (see below for nested schema)
	// Configuration for the proxy repository
	Proxy []NugetProxyProxyInitParameters `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// (Number) How long to cache query results from the proxied repository (in seconds)
	// How long to cache query results from the proxied repository (in seconds)
	QueryCacheItemMaxAge *float64 `json:"queryCacheItemMaxAge,omitempty" tf:"query_cache_item_max_age,omitempty"`

	// (String) The name of the routing rule assigned to this repository
	// The name of the routing rule assigned to this repository
	RoutingRule *string `json:"routingRule,omitempty" tf:"routing_rule,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	Storage []NugetProxyStorageInitParameters `json:"storage,omitempty" tf:"storage,omitempty"`
}

type NugetProxyNegativeCacheInitParameters struct {

	// (Boolean) Whether to cache responses for content not present in the proxied repository
	// Whether to cache responses for content not present in the proxied repository
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Number) How long to cache the fact that a file was not found in the repository (in minutes)
	// How long to cache the fact that a file was not found in the repository (in minutes)
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type NugetProxyNegativeCacheObservation struct {

	// (Boolean) Whether to cache responses for content not present in the proxied repository
	// Whether to cache responses for content not present in the proxied repository
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Number) How long to cache the fact that a file was not found in the repository (in minutes)
	// How long to cache the fact that a file was not found in the repository (in minutes)
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type NugetProxyNegativeCacheParameters struct {

	// (Boolean) Whether to cache responses for content not present in the proxied repository
	// Whether to cache responses for content not present in the proxied repository
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// (Number) How long to cache the fact that a file was not found in the repository (in minutes)
	// How long to cache the fact that a file was not found in the repository (in minutes)
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl" tf:"ttl,omitempty"`
}

type NugetProxyObservation struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	Cleanup []NugetProxyCleanupObservation `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Min: 1, Max: 1) HTTP Client configuration for proxy repositories (see below for nested schema)
	// HTTP Client configuration for proxy repositories
	HTTPClient []NugetProxyHTTPClientObservation `json:"httpClient,omitempty" tf:"http_client,omitempty"`

	// (String) Used to identify resource at nexus
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Min: 1, Max: 1) Configuration of the negative cache handling (see below for nested schema)
	// Configuration of the negative cache handling
	NegativeCache []NugetProxyNegativeCacheObservation `json:"negativeCache,omitempty" tf:"negative_cache,omitempty"`

	// (String) Nuget protocol version
	// Nuget protocol version
	NugetVersion *string `json:"nugetVersion,omitempty" tf:"nuget_version,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// (Block List, Min: 1, Max: 1) Configuration for the proxy repository (see below for nested schema)
	// Configuration for the proxy repository
	Proxy []NugetProxyProxyObservation `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// (Number) How long to cache query results from the proxied repository (in seconds)
	// How long to cache query results from the proxied repository (in seconds)
	QueryCacheItemMaxAge *float64 `json:"queryCacheItemMaxAge,omitempty" tf:"query_cache_item_max_age,omitempty"`

	// (String) The name of the routing rule assigned to this repository
	// The name of the routing rule assigned to this repository
	RoutingRule *string `json:"routingRule,omitempty" tf:"routing_rule,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	Storage []NugetProxyStorageObservation `json:"storage,omitempty" tf:"storage,omitempty"`
}

type NugetProxyParameters struct {

	// (Block List) Cleanup policies (see below for nested schema)
	// Cleanup policies
	// +kubebuilder:validation:Optional
	Cleanup []NugetProxyCleanupParameters `json:"cleanup,omitempty" tf:"cleanup,omitempty"`

	// (Block List, Min: 1, Max: 1) HTTP Client configuration for proxy repositories (see below for nested schema)
	// HTTP Client configuration for proxy repositories
	// +kubebuilder:validation:Optional
	HTTPClient []NugetProxyHTTPClientParameters `json:"httpClient,omitempty" tf:"http_client,omitempty"`

	// (String) A unique identifier for this repository
	// A unique identifier for this repository
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Min: 1, Max: 1) Configuration of the negative cache handling (see below for nested schema)
	// Configuration of the negative cache handling
	// +kubebuilder:validation:Optional
	NegativeCache []NugetProxyNegativeCacheParameters `json:"negativeCache,omitempty" tf:"negative_cache,omitempty"`

	// (String) Nuget protocol version
	// Nuget protocol version
	// +kubebuilder:validation:Optional
	NugetVersion *string `json:"nugetVersion,omitempty" tf:"nuget_version,omitempty"`

	// (Boolean) Whether this repository accepts incoming requests
	// Whether this repository accepts incoming requests
	// +kubebuilder:validation:Optional
	Online *bool `json:"online,omitempty" tf:"online,omitempty"`

	// (Block List, Min: 1, Max: 1) Configuration for the proxy repository (see below for nested schema)
	// Configuration for the proxy repository
	// +kubebuilder:validation:Optional
	Proxy []NugetProxyProxyParameters `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// (Number) How long to cache query results from the proxied repository (in seconds)
	// How long to cache query results from the proxied repository (in seconds)
	// +kubebuilder:validation:Optional
	QueryCacheItemMaxAge *float64 `json:"queryCacheItemMaxAge,omitempty" tf:"query_cache_item_max_age,omitempty"`

	// (String) The name of the routing rule assigned to this repository
	// The name of the routing rule assigned to this repository
	// +kubebuilder:validation:Optional
	RoutingRule *string `json:"routingRule,omitempty" tf:"routing_rule,omitempty"`

	// (Block List, Min: 1, Max: 1) The storage configuration of the repository (see below for nested schema)
	// The storage configuration of the repository
	// +kubebuilder:validation:Optional
	Storage []NugetProxyStorageParameters `json:"storage,omitempty" tf:"storage,omitempty"`
}

type NugetProxyProxyInitParameters struct {

	// (Number) How long (in minutes) to cache artifacts before rechecking the remote repository
	// How long (in minutes) to cache artifacts before rechecking the remote repository
	ContentMaxAge *float64 `json:"contentMaxAge,omitempty" tf:"content_max_age,omitempty"`

	// (Number) How long (in minutes) to cache metadata before rechecking the remote repository.
	// How long (in minutes) to cache metadata before rechecking the remote repository.
	MetadataMaxAge *float64 `json:"metadataMaxAge,omitempty" tf:"metadata_max_age,omitempty"`

	// (String) Location of the remote repository being proxied
	// Location of the remote repository being proxied
	RemoteURL *string `json:"remoteUrl,omitempty" tf:"remote_url,omitempty"`
}

type NugetProxyProxyObservation struct {

	// (Number) How long (in minutes) to cache artifacts before rechecking the remote repository
	// How long (in minutes) to cache artifacts before rechecking the remote repository
	ContentMaxAge *float64 `json:"contentMaxAge,omitempty" tf:"content_max_age,omitempty"`

	// (Number) How long (in minutes) to cache metadata before rechecking the remote repository.
	// How long (in minutes) to cache metadata before rechecking the remote repository.
	MetadataMaxAge *float64 `json:"metadataMaxAge,omitempty" tf:"metadata_max_age,omitempty"`

	// (String) Location of the remote repository being proxied
	// Location of the remote repository being proxied
	RemoteURL *string `json:"remoteUrl,omitempty" tf:"remote_url,omitempty"`
}

type NugetProxyProxyParameters struct {

	// (Number) How long (in minutes) to cache artifacts before rechecking the remote repository
	// How long (in minutes) to cache artifacts before rechecking the remote repository
	// +kubebuilder:validation:Optional
	ContentMaxAge *float64 `json:"contentMaxAge,omitempty" tf:"content_max_age,omitempty"`

	// (Number) How long (in minutes) to cache metadata before rechecking the remote repository.
	// How long (in minutes) to cache metadata before rechecking the remote repository.
	// +kubebuilder:validation:Optional
	MetadataMaxAge *float64 `json:"metadataMaxAge,omitempty" tf:"metadata_max_age,omitempty"`

	// (String) Location of the remote repository being proxied
	// Location of the remote repository being proxied
	// +kubebuilder:validation:Optional
	RemoteURL *string `json:"remoteUrl" tf:"remote_url,omitempty"`
}

type NugetProxyStorageInitParameters struct {

	// (String) Blob store used to store repository contents
	// Blob store used to store repository contents
	BlobStoreName *string `json:"blobStoreName,omitempty" tf:"blob_store_name,omitempty"`

	// (Boolean) Whether to validate uploaded content's MIME type appropriate for the repository format
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation *bool `json:"strictContentTypeValidation,omitempty" tf:"strict_content_type_validation,omitempty"`
}

type NugetProxyStorageObservation struct {

	// (String) Blob store used to store repository contents
	// Blob store used to store repository contents
	BlobStoreName *string `json:"blobStoreName,omitempty" tf:"blob_store_name,omitempty"`

	// (Boolean) Whether to validate uploaded content's MIME type appropriate for the repository format
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation *bool `json:"strictContentTypeValidation,omitempty" tf:"strict_content_type_validation,omitempty"`
}

type NugetProxyStorageParameters struct {

	// (String) Blob store used to store repository contents
	// Blob store used to store repository contents
	// +kubebuilder:validation:Optional
	BlobStoreName *string `json:"blobStoreName" tf:"blob_store_name,omitempty"`

	// (Boolean) Whether to validate uploaded content's MIME type appropriate for the repository format
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	// +kubebuilder:validation:Optional
	StrictContentTypeValidation *bool `json:"strictContentTypeValidation,omitempty" tf:"strict_content_type_validation,omitempty"`
}

// NugetProxySpec defines the desired state of NugetProxy
type NugetProxySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NugetProxyParameters `json:"forProvider"`
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
	InitProvider NugetProxyInitParameters `json:"initProvider,omitempty"`
}

// NugetProxyStatus defines the observed state of NugetProxy.
type NugetProxyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NugetProxyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NugetProxy is the Schema for the NugetProxys API. Use this resource to create an NPM proxy repository.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nexus}
type NugetProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.httpClient) || (has(self.initProvider) && has(self.initProvider.httpClient))",message="spec.forProvider.httpClient is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.negativeCache) || (has(self.initProvider) && has(self.initProvider.negativeCache))",message="spec.forProvider.negativeCache is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nugetVersion) || (has(self.initProvider) && has(self.initProvider.nugetVersion))",message="spec.forProvider.nugetVersion is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.proxy) || (has(self.initProvider) && has(self.initProvider.proxy))",message="spec.forProvider.proxy is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.queryCacheItemMaxAge) || (has(self.initProvider) && has(self.initProvider.queryCacheItemMaxAge))",message="spec.forProvider.queryCacheItemMaxAge is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storage) || (has(self.initProvider) && has(self.initProvider.storage))",message="spec.forProvider.storage is a required parameter"
	Spec   NugetProxySpec   `json:"spec"`
	Status NugetProxyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NugetProxyList contains a list of NugetProxys
type NugetProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NugetProxy `json:"items"`
}

// Repository type metadata.
var (
	NugetProxy_Kind             = "NugetProxy"
	NugetProxy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NugetProxy_Kind}.String()
	NugetProxy_KindAPIVersion   = NugetProxy_Kind + "." + CRDGroupVersion.String()
	NugetProxy_GroupVersionKind = CRDGroupVersion.WithKind(NugetProxy_Kind)
)

func init() {
	SchemeBuilder.Register(&NugetProxy{}, &NugetProxyList{})
}
