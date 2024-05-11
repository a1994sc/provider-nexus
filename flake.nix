{
  inputs = {
    # keep-sorted start block=yes case=no
    flake-utils.url = "github:numtide/flake-utils";
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    systems.url = "github:nix-systems/default";
    treefmt-nix = {
      inputs.nixpkgs.follows = "nixpkgs";
      url = "github:numtide/treefmt-nix";
    };
    # keep-sorted end
  };

  outputs =
    {
      self,
      flake-utils,
      nixpkgs,
      treefmt-nix,
      systems,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        treefmtEval = treefmt-nix.lib.evalModule pkgs (
          { pkgs, ... }:
          {
            # keep-sorted start block=yes prefix_order=projectRootFile,
            projectRootFile = "flake.nix";
            programs.deadnix.enable = true;
            programs.keep-sorted.enable = true;
            programs.nixfmt = {
              enable = true;
              package = pkgs.nixfmt-rfc-style;
            };
            programs.statix.enable = true;
            # Disabled on "flake.nix" because of some false positivies.
            settings.formatter.deadnix.excludes = [ "**/flake.nix" ];
            # keep-sorted end
          }
        );
        shellHook = ''
          GOROOT="$(dirname $(dirname $(which go)))/share/go"
          DOCKER_HOST="unix://$(podman info | yq .host.remoteSocket.path)"
          BUILD_REGISTRY="ghcr.io"
          export GOROOT
          export DOCKER_HOST
          export BUILD_REGISTRY
          unset GOPATH;
        '';
        podmanSetupScript =
          let
            registriesConf = pkgs.writeText "registries.conf" ''
              [registries.search]
              registries = ['docker.io']
              [registries.block]
              registries = []
            '';
          in
          pkgs.writeScript "podman-setup" ''
            #!${pkgs.runtimeShell}
            # Dont overwrite customised configuration
            if ! test -f ~/.config/containers/policy.json; then
              install -Dm555 ${pkgs.skopeo.src}/default-policy.json ~/.config/containers/policy.json
            fi
            if ! test -f ~/.config/containers/registries.conf; then
              install -Dm555 ${registriesConf} ~/.config/containers/registries.conf
            fi
          '';

        # Provides a fake "docker" binary mapping to podman
        dockerCompat = pkgs.runCommandNoCC "docker-podman-compat" { } ''
          mkdir -p $out/bin
          ln -s ${pkgs.podman}/bin/podman $out/bin/docker
        '';
      in
      with pkgs;
      {
        devShells = {
          # $ nix develop
          default = pkgs.mkShell {
            inherit shellHook;
            buildInputs = with pkgs; [
              dockerCompat
              unzip
              podman # Docker compat
              runc # Container runtime
              conmon # Container runtime monitor
              skopeo # Interact with container registry
              slirp4netns # User-mode networking for unprivileged namespaces
              fuse-overlayfs # CoW for images, much faster than default vfs
            ];
            nativeBuildInputs = with pkgs; [ 
              go
              gotools
              gopls
              golint
            ];
          };
        };
        formatter = treefmtEval.config.build.wrapper;
      }
    );
}
