package registry

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// Group is the short group for this provider.
	Group = "repository"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nexus_repository_apt_hosted", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("nexus_repository_apt_proxy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group

		// Allow for the username to be set via a secret
		r.TerraformResource.Schema["http_client"].Elem.(*schema.Resource).Schema["authentication"].Elem.(*schema.Resource).Schema["username"].Sensitive = true
	})

	p.AddResourceConfigurator("nexus_repository_docker_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("nexus_repository_docker_hosted", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("nexus_repository_docker_proxy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group

		// Allow for the username to be set via a secret
		r.TerraformResource.Schema["http_client"].Elem.(*schema.Resource).Schema["authentication"].Elem.(*schema.Resource).Schema["username"].Sensitive = true
	})

	p.AddResourceConfigurator("nexus_repository_bower_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("nexus_repository_bower_hosted", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("nexus_repository_bower_proxy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group

		// Allow for the username to be set via a secret
		r.TerraformResource.Schema["http_client"].Elem.(*schema.Resource).Schema["authentication"].Elem.(*schema.Resource).Schema["username"].Sensitive = true
	})

	p.AddResourceConfigurator("nexus_repository_cocoapods_proxy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group

		// Allow for the username to be set via a secret
		r.TerraformResource.Schema["http_client"].Elem.(*schema.Resource).Schema["authentication"].Elem.(*schema.Resource).Schema["username"].Sensitive = true
	})

	p.AddResourceConfigurator("nexus_repository_conan_proxy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group

		// Allow for the username to be set via a secret
		r.TerraformResource.Schema["http_client"].Elem.(*schema.Resource).Schema["authentication"].Elem.(*schema.Resource).Schema["username"].Sensitive = true
	})

	p.AddResourceConfigurator("nexus_repository_conda_proxy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group

		// Allow for the username to be set via a secret
		r.TerraformResource.Schema["http_client"].Elem.(*schema.Resource).Schema["authentication"].Elem.(*schema.Resource).Schema["username"].Sensitive = true
	})

	p.AddResourceConfigurator("nexus_repository_helm_hosted", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("nexus_repository_helm_proxy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group

		// Allow for the username to be set via a secret
		r.TerraformResource.Schema["http_client"].Elem.(*schema.Resource).Schema["authentication"].Elem.(*schema.Resource).Schema["username"].Sensitive = true
	})

	p.AddResourceConfigurator("nexus_repository_raw_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("nexus_repository_raw_hosted", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("nexus_repository_raw_proxy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group

		// Allow for the username to be set via a secret
		r.TerraformResource.Schema["http_client"].Elem.(*schema.Resource).Schema["authentication"].Elem.(*schema.Resource).Schema["username"].Sensitive = true
	})

	p.AddResourceConfigurator("nexus_repository_rubygems_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("nexus_repository_rubygems_hosted", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("nexus_repository_rubygems_proxy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group

		// Allow for the username to be set via a secret
		r.TerraformResource.Schema["http_client"].Elem.(*schema.Resource).Schema["authentication"].Elem.(*schema.Resource).Schema["username"].Sensitive = true
	})

	p.AddResourceConfigurator("nexus_repository_yum_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("nexus_repository_yum_hosted", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("nexus_repository_yum_proxy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group

		// Allow for the username to be set via a secret
		r.TerraformResource.Schema["http_client"].Elem.(*schema.Resource).Schema["authentication"].Elem.(*schema.Resource).Schema["username"].Sensitive = true
	})
}
