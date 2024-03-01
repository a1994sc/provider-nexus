package routing

import (
	"github.com/crossplane/upjet/pkg/config"
)

const (
	// Group is the short group for this provider.
	Group = "routing"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nexus_routing_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
}
