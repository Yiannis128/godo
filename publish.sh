#!/usr/bin/env sh

version=$1
GOPROXY=proxy.golang.org go list -m github.com/Yiannis128/manews_common@$version