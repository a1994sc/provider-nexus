package registry

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nexus_repository_yum_proxy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "repository"

		// r.References["storage"] = config.Reference{
		// 	Type: "github.com/crossplane-contrib/provider-keycloak/apis/role/v1alpha1.Role",
		// }
		// r.References = map[string]config.Reference{
		// 	"storage": {
		// 		Type: "Registry",
		// 	},
		// }
	})
}
