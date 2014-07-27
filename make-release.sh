#!/bin/bash

# This script creates 64-bit release builds for OSX, Linux and Windows. It
# requires Go to be set up for cross-compiling to these platforms. Some package
# managers can set this up for you (homebrew can do this, for example). For
# instructions on how to set up cross-compiling manually, there's an excellent
# guide here:
# http://dave.cheney.net/2013/07/09/an-introduction-to-cross-compilation-with-go-1-1.

set -e

release_dir=release
name=commitfmt
compress_cmd="zip -9 -m -T -D"

mkdir -p $release_dir

GOOS=darwin GOARCH=amd64 go build -o $release_dir/$name
pushd $release_dir
$compress_cmd $name-osx.zip $name
popd

GOOS=linux GOARCH=amd64 go build -o $release_dir/$name
pushd $release_dir
$compress_cmd $name-linux.zip $name
popd

GOOS=windows GOARCH=amd64 go build -o $release_dir/$name.exe
pushd $release_dir
$compress_cmd $name-windows.zip $name.exe
popd
