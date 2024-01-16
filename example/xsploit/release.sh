#!/bin/bash
set -ex
# go get github.com/mitchellh/gox
cd "$(dirname "$(readlink -m "$0")")"
mkdir -p bin/release
cd bin/release
CGO_ENABLED=0 gox -cgo=0 -osarch="linux/amd64" -osarch="linux/arm64" -ldflags "${LDFLAGS}" xsploit/cmd/xsploit
CGO_ENABLED=0 gox -cgo=0 -osarch="linux/amd64" -osarch="linux/arm64" -ldflags "${LDFLAGS}" xsploit/cmd/checksec
cd -
upx bin/release/*
