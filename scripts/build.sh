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

clean
build "go-mindustry-ping.exe" "v0.0.2" "windows" "amd64"
build "go-mindustry-ping" "v0.0.2" "linux" "amd64"
build "go-mindustry-ping" "v0.0.2" "darwin" "amd64"

"../build/windows-amd64/go-mindustry-ping.exe" --host "162.248.100.133"
"../build/windows-amd64/go-mindustry-ping.exe" --help
"../build/windows-amd64/go-mindustry-ping.exe" --host "mindustry.ddns.net"
"../build/windows-amd64/go-mindustry-ping.exe" --host "mdt.mdtleague.top"
"../build/windows-amd64/go-mindustry-ping.exe" --host "cn.mindustry.top"
"../build/windows-amd64/go-mindustry-ping.exe" --host "mindurka.fun"
