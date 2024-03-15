/*
Copyright 2022 Upbound Inc.
*/

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
	bowergroup "github.com/a1994sc/provider-nexus/internal/controller/repository/bowergroup"
	bowerhosted "github.com/a1994sc/provider-nexus/internal/controller/repository/bowerhosted"
	bowerproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/bowerproxy"
	cocoapodsproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/cocoapodsproxy"
	conanproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/conanproxy"
	condaproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/condaproxy"
	dockergroup "github.com/a1994sc/provider-nexus/internal/controller/repository/dockergroup"
	dockerhosted "github.com/a1994sc/provider-nexus/internal/controller/repository/dockerhosted"
	dockerproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/dockerproxy"
	gitlfshosted "github.com/a1994sc/provider-nexus/internal/controller/repository/gitlfshosted"
	gogroup "github.com/a1994sc/provider-nexus/internal/controller/repository/gogroup"
	goproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/goproxy"
	helmhosted "github.com/a1994sc/provider-nexus/internal/controller/repository/helmhosted"
	helmproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/helmproxy"
	mavengroup "github.com/a1994sc/provider-nexus/internal/controller/repository/mavengroup"
	mavenhosted "github.com/a1994sc/provider-nexus/internal/controller/repository/mavenhosted"
	mavenproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/mavenproxy"
	npmgroup "github.com/a1994sc/provider-nexus/internal/controller/repository/npmgroup"
	npmhosted "github.com/a1994sc/provider-nexus/internal/controller/repository/npmhosted"
	npmproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/npmproxy"
	nugetgroup "github.com/a1994sc/provider-nexus/internal/controller/repository/nugetgroup"
	nugethosted "github.com/a1994sc/provider-nexus/internal/controller/repository/nugethosted"
	nugetproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/nugetproxy"
	p2proxy "github.com/a1994sc/provider-nexus/internal/controller/repository/p2proxy"
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
	ssltruststore "github.com/a1994sc/provider-nexus/internal/controller/security/ssltruststore"
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
		bowergroup.Setup,
		bowerhosted.Setup,
		bowerproxy.Setup,
		cocoapodsproxy.Setup,
		conanproxy.Setup,
		condaproxy.Setup,
		dockergroup.Setup,
		dockerhosted.Setup,
		dockerproxy.Setup,
		gitlfshosted.Setup,
		gogroup.Setup,
		goproxy.Setup,
		helmhosted.Setup,
		helmproxy.Setup,
		mavengroup.Setup,
		mavenhosted.Setup,
		mavenproxy.Setup,
		npmgroup.Setup,
		npmhosted.Setup,
		npmproxy.Setup,
		nugetgroup.Setup,
		nugethosted.Setup,
		nugetproxy.Setup,
		p2proxy.Setup,
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
		ssltruststore.Setup,
		user.Setup,
		usertoken.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
