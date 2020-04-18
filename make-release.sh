#!/bin/sh

# This script creates 64-bit release builds for OSX, Linux and Windows. It
# requires Go to be set up for cross-compiling to these platforms. Some package
# managers can set this up for you (homebrew can do this, for example). For
# instructions on how to set up cross-compiling manually, there's an excellent
# guide here:
# http://dave.cheney.net/2013/07/09/an-introduction-to-cross-compilation-with-go-1-1.

set -e

mkdir -p release

GOOS=darwin GOARCH=amd64 go build -o release/commitfmt
cd release
tar -czf commitfmt-macos.tar.gz commitfmt
rm commitfmt
cd ..

GOOS=linux GOARCH=amd64 go build -o release/commitfmt
cd release
tar -czf commitfmt-linux.tar.gz commitfmt
rm commitfmt
cd ..

GOOS=windows GOARCH=amd64 go build -o release/commitfmt.exe
cd release
zip -9 -m -T -D -q commitfmt-windows.zip commitfmt.exe
cd ..
