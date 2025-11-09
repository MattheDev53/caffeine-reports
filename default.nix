{
  pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
    in
    import (fetchTree nixpkgs.locked) {
      overlays = [
        (import "${fetchTree gomod2nix.locked}/overlay.nix")
      ];
    }
  ),
  buildGoApplication ? pkgs.buildGoApplication,
}:

buildGoApplication {
  pname = "caffeine-reports";
  version = "0.1.2";
  pwd = ./.;
  src = ./.;
  modules = ./gomod2nix.toml;

  buildInputs = with pkgs; [
    xorg.libX11
    xorg.libXrandr
    xorg.libXcursor
    xorg.libXinerama
    xorg.xinput
    xorg.libXi
    xorg.libXxf86vm
    libxkbcommon
    wayland

    libGL
  ];

  nativeBuildInputs = with pkgs; [
    pkg-config
  ];

  /*
  desktopItems = [
    (makeDesktopItem {
      name = "CaffeineReports";
      exec = pname;
      icon = pname;
      desktopName = pname;
    })
  ];
  */
}
