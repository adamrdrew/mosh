#!/bin/bash
mkdir -p mosh_tmp
mkdir -p mosh_tmp/cache
export MOSH_CONFIG_DIR=./mosh_tmp
export MOSH_LOG_DIR=./mosh_tmp
export MOSH_PID_DIR=./mosh_tmp
export MOSH_PORT=9777
export MOSH_CACHE_DIR=./mosh_tmp/cache
VERSION=`git describe --tags`
LDFLAGS="-X main.Version=$VERSION"
go run -ldflags "$LDFLAGS" mosh.go $@