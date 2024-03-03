/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"nexus_blobstore_s3":    config.IdentifierFromProvider,
	"nexus_blobstore_file":  config.IdentifierFromProvider,
	"nexus_blobstore_group": config.IdentifierFromProvider,
	"nexus_blobstore_azure": config.IdentifierFromProvider,

	"nexus_repository_docker_proxy":  config.IdentifierFromProvider,
	"nexus_repository_docker_group":  config.IdentifierFromProvider,
	"nexus_repository_docker_hosted": config.IdentifierFromProvider,

	"nexus_repository_helm_proxy":  config.IdentifierFromProvider,
	"nexus_repository_helm_hosted": config.IdentifierFromProvider,

	"nexus_repository_rubygems_proxy":  config.IdentifierFromProvider,
	"nexus_repository_rubygems_group":  config.IdentifierFromProvider,
	"nexus_repository_rubygems_hosted": config.IdentifierFromProvider,

	"nexus_repository_yum_proxy":  config.IdentifierFromProvider,
	"nexus_repository_yum_group":  config.IdentifierFromProvider,
	"nexus_repository_yum_hosted": config.IdentifierFromProvider,

	"nexus_security_anonymous":        config.IdentifierFromProvider,
	"nexus_security_content_selector": config.IdentifierFromProvider,
	"nexus_security_ldap":             config.IdentifierFromProvider,
	"nexus_security_ldap_order":       config.IdentifierFromProvider,
	"nexus_security_realms":           config.IdentifierFromProvider,
	"nexus_security_role":             config.IdentifierFromProvider,
	"nexus_security_saml":             config.IdentifierFromProvider,
	"nexus_security_user":             config.IdentifierFromProvider,
	"nexus_security_user_token":       config.IdentifierFromProvider,

	"nexus_routing_rule": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
