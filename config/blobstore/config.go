package blobstore

import "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "blobstore"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nexus_blobstore_s3", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
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
