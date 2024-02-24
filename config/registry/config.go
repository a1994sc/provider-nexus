package registry

import "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "repository"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nexus_repository_docker_proxy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("nexus_repository_docker_hosted", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("nexus_repository_docker_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("nexus_repository_yum_proxy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("nexus_repository_yum_hosted", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("nexus_repository_yum_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
}
