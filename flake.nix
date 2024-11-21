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
              golangci-lint
              crane # Manipulate oci repos
              upbound # used to build crossplane artifacts
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
