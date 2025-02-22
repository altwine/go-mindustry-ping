#!/bin/bash
set -e

function clean () {
    rm -rf "../build"
}

function build () {
    local binary_name="$1"
    local binary_version="$2"
    export GOOS="$3"
    export GOARCH="$4"
    echo "Building $GOOS-$GOARCH/$binary_name"
    go build \
        -C "../cmd/go-mindustry-ping" \
        -ldflags="-s -w -X 'main.BINARY_VERSION=$binary_version' -X 'main.BINARY_ARCH=$GOARCH' -X 'main.BINARY_OS=$GOOS'" \
        -o "../../build/$GOOS-$GOARCH/$binary_name"
    if [ $? -ne 0 ];
    then
        echo "Failed!"
        return
    fi
    echo "Success!"
}

clean
build "go-mindustry-ping.exe" "v0.0.3" "windows" "amd64"
build "go-mindustry-ping" "v0.0.3" "linux" "amd64"
build "go-mindustry-ping" "v0.0.3" "darwin" "amd64"