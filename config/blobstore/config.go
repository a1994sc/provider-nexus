package blobstore

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/crossplane/upjet/pkg/config"
)

const (
	// Group is the short group for this provider.
	Group = "blobstore"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nexus_blobstore_s3", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group

		// Allow for the access id to be set via a secret
		r.TerraformResource.Schema["bucket_configuration"].Elem.(*schema.Resource).Schema["bucket_security"].Elem.(*schema.Resource).Schema["access_key_id"].Sensitive = true
	})
	p.AddResourceConfigurator("nexus_blobstore_file", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("nexus_blobstore_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("nexus_blobstore_azure", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
}
