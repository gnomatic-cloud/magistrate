#!/usr/bin/env bash

readonly PACKAGE_DIR="$(dirname "${BASH_SOURCE[0]}")"
cd $PACKAGE_DIR

# Uses: https://github.com/globusdigital/deep-copy
deep-copy \
    -pointer-receiver \
    -o ./config.deepcopy.go \
    -type RuntimeConfig \
    ./
