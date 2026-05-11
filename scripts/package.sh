#!/bin/bash
set -e

function package () {
	local base_directory=$(echo "$1" | cut -d "/" -f1)
	local filename=$(basename "$1")

	mkdir -p "../build/releases"

	if [[ "$filename" == *.exe ]]; then
		zip -j "../build/releases/$base_directory.zip" "../build/$1"
	else
		tar czf "../build/releases/$base_directory.tar.gz" -C "../build/$base_directory" .
	fi
}

package "windows-amd64/go-mindustry-ping.exe"
package "linux-amd64/go-mindustry-ping"
package "darwin-amd64/go-mindustry-ping"
