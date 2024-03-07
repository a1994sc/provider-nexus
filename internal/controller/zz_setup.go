// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	azure "github.com/a1994sc/provider-nexus/internal/controller/blobstore/azure"
	file "github.com/a1994sc/provider-nexus/internal/controller/blobstore/file"
	group "github.com/a1994sc/provider-nexus/internal/controller/blobstore/group"
	s3 "github.com/a1994sc/provider-nexus/internal/controller/blobstore/s3"
	providerconfig "github.com/a1994sc/provider-nexus/internal/controller/providerconfig"
	apthosted "github.com/a1994sc/provider-nexus/internal/controller/repository/apthosted"
	aptproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/aptproxy"
	dockergroup "github.com/a1994sc/provider-nexus/internal/controller/repository/dockergroup"
	dockerhosted "github.com/a1994sc/provider-nexus/internal/controller/repository/dockerhosted"
	dockerproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/dockerproxy"
	helmhosted "github.com/a1994sc/provider-nexus/internal/controller/repository/helmhosted"
	helmproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/helmproxy"
	rawgroup "github.com/a1994sc/provider-nexus/internal/controller/repository/rawgroup"
	rawhosted "github.com/a1994sc/provider-nexus/internal/controller/repository/rawhosted"
	rawproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/rawproxy"
	rubygemsgroup "github.com/a1994sc/provider-nexus/internal/controller/repository/rubygemsgroup"
	rubygemshosted "github.com/a1994sc/provider-nexus/internal/controller/repository/rubygemshosted"
	rubygemsproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/rubygemsproxy"
	yumgroup "github.com/a1994sc/provider-nexus/internal/controller/repository/yumgroup"
	yumhosted "github.com/a1994sc/provider-nexus/internal/controller/repository/yumhosted"
	yumproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/yumproxy"
	rule "github.com/a1994sc/provider-nexus/internal/controller/routing/rule"
	anonymous "github.com/a1994sc/provider-nexus/internal/controller/security/anonymous"
	contentselector "github.com/a1994sc/provider-nexus/internal/controller/security/contentselector"
	ldap "github.com/a1994sc/provider-nexus/internal/controller/security/ldap"
	ldaporder "github.com/a1994sc/provider-nexus/internal/controller/security/ldaporder"
	realms "github.com/a1994sc/provider-nexus/internal/controller/security/realms"
	role "github.com/a1994sc/provider-nexus/internal/controller/security/role"
	saml "github.com/a1994sc/provider-nexus/internal/controller/security/saml"
	user "github.com/a1994sc/provider-nexus/internal/controller/security/user"
	usertoken "github.com/a1994sc/provider-nexus/internal/controller/security/usertoken"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		azure.Setup,
		file.Setup,
		group.Setup,
		s3.Setup,
		providerconfig.Setup,
		apthosted.Setup,
		aptproxy.Setup,
		dockergroup.Setup,
		dockerhosted.Setup,
		dockerproxy.Setup,
		helmhosted.Setup,
		helmproxy.Setup,
		rawgroup.Setup,
		rawhosted.Setup,
		rawproxy.Setup,
		rubygemsgroup.Setup,
		rubygemshosted.Setup,
		rubygemsproxy.Setup,
		yumgroup.Setup,
		yumhosted.Setup,
		yumproxy.Setup,
		rule.Setup,
		anonymous.Setup,
		contentselector.Setup,
		ldap.Setup,
		ldaporder.Setup,
		realms.Setup,
		role.Setup,
		saml.Setup,
		user.Setup,
		usertoken.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
