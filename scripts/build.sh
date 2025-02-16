#!/bin/bash
set -e

function clean () {
    rm -rf "../build"
}

function build () {
    local binary_name="$1"
    local binary_version="$2"
    TARGET_OS="$3"
    TARGET_ARCH="$4"
    echo "Building $TARGET_OS-$TARGET_ARCH/$binary_name"
    go build \
        -C "../cmd/go-mindustry-ping" \
        -ldflags="-s -w -X 'main.BINARY_VERSION=$binary_version'" \
        -o "../../build/$TARGET_OS-$TARGET_ARCH/$binary_name"
    if [ $? -ne 0 ];
    then
        echo "Failed!"
        return
    fi
    echo "Success!"
}

function compress () {
    echo "Compressing $1"
    upx --lzma --best -q "../build/$1" > /dev/null 2>&1
    if [ $? -ne 0 ];
    then
        echo "Failed!"
        return
    fi
    echo "Success!"
}

clean
build "go-mindustry-ping.exe" "0.0.1" "windows" "amd64"
build "go-mindustry-ping" "0.0.1" "linux" "amd64"
build "go-mindustry-ping" "0.0.1" "darwin" "amd64"
compress "windows-amd64/go-mindustry-ping.exe"
compress "linux-amd64/go-mindustry-ping"
compress "darwin-amd64/go-mindustry-ping"
