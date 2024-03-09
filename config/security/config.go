package security

import (
	"github.com/crossplane/upjet/pkg/config"
)

const (
	// Group is the short group for this provider.
	Group = "security"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nexus_security_anonymous", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("nexus_security_content_selector", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("nexus_security_ldap", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("nexus_security_ldap_order", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("nexus_security_realms", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("nexus_security_role", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("nexus_security_saml", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("nexus_security_ssl_truststore", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group

		r.TerraformResource.Schema["pem"].Sensitive = true
	})
	p.AddResourceConfigurator("nexus_security_user", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("nexus_security_user_token", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
}
