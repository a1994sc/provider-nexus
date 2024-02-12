{ pkgs ? import <nixpkgs-unstable> {}}:
pkgs.mkShell {
  packages = [
    pkgs.go
    pkgs.gotools
  ];
}
