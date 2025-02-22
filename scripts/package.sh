#!/bin/bash
set -e

function package () {
    local base_directory=$(echo "$1" | cut -d "/" -f1)
    mkdir -p "../build/releases"
    tar czf "../build/releases/$base_directory.tar.gz" -C "../build/$base_directory" .
}

package "windows-amd64/go-mindustry-ping.exe"
package "linux-amd64/go-mindustry-ping"
package "darwin-amd64/go-mindustry-ping"