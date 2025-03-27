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
        # pkgs = nixpkgs.legacyPackages.${system};
        pkgs = import nixpkgs {
          inherit system;
          config.allowUnfreePredicate =
            pkg:
            builtins.elem (nixpkgs.lib.getName pkg) [
              "upbound"
            ];
        };
        treefmtEval = treefmt-nix.lib.evalModule pkgs (
          { pkgs, ... }:
          {
            # keep-sorted start block=yes prefix_order=projectRootFile,
            projectRootFile = "flake.nix";
            programs.keep-sorted.enable = true;
            programs.nixfmt = {
              enable = true;
              package = pkgs.nixfmt-rfc-style;
            };
            programs.statix.enable = true;
            # keep-sorted end
          }
        );
        shellHook = ''
          GOROOT="$(dirname $(dirname $(which go)))/share/go"
          BUILD_REGISTRY="ghcr.io"
          export GOROOT
          export DOCKER_HOST
          export BUILD_REGISTRY
          unset GOPATH;
        '';
      in
      with pkgs;
      {
        devShells = {
          # $ nix develop
          default = pkgs.mkShell {
            inherit shellHook;
            buildInputs = with pkgs; [
              unzip
              crossplane-cli
              crane # Manipulate oci repos
              upbound # used to build crossplane artifacts
            ];
            nativeBuildInputs = with pkgs; [
              go
              gotools
              gopls
              golint
              golangci-lint
            ];
          };
        };
        formatter = treefmtEval.config.build.wrapper;
      }
    );
}
