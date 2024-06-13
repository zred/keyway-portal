{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.nodejs-18_x
    pkgs.go
    pkgs.sqlite
    pkgs.yarn
  ];
  
  shellHook = ''
    echo "Setting up environment..."
    
    # Install Node.js dependencies
    if [ ! -d "node_modules" ]; then
        yarn install
    fi
  '';
}

