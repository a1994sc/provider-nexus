# Provider Nexus

`provider-nexus` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Nexus API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/a1994sc/provider-nexus):
```
up ctp provider install ghcr.io/a1994sc/crossplane/provider-nexus:v0.0.6
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-nexus
spec:
  package: ghcr.io/a1994sc/crossplane/provider-nexus:v0.0.6
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/a1994sc/provider-nexus).

## Terraform Provider
The upstream terraform provider is created and maintained by [@datadrivers](https://github.com/datadrivers)

* [datadrivers/terraform-provider-nexus](https://github.com/datadrivers/terraform-provider-nexus)

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/a1994sc/provider-nexus/issues).
