/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/a1994sc/provider-nexus/config/blobstore"
	"github.com/a1994sc/provider-nexus/config/registry"
)

const (
	resourcePrefix = "nexus"
	modulePath     = "github.com/a1994sc/provider-nexus"
	rootGroup      = resourcePrefix + ".adrp.xyz"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			// KnownReferencers(),
		),
		ujconfig.WithRootGroup(rootGroup))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		registry.Configure,
		blobstore.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
