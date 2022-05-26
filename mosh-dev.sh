#!/bin/bash
mkdir -p mosh_tmp
export MOSH_CONFIG_DIR=./mosh_tmp
export MOSH_LOG_DIR=./mosh_tmp
export MOSH_PID_DIR=./mosh_tmp
export MOSH_PORT=9777
go run mosh.go $@