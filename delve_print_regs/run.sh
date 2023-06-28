#!/usr/bin/env bash

DLV_PATH="$HOME/Code/src/github.com/go-delve/delve/"
pushd $DLV_PATH
make build
popd
$DLV_PATH/dlv debug main.go --init ./delve.script
