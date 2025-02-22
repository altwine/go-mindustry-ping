# go-mindustry-ping v0.0.2
Go Mindustry Ping - command-line tool to query and display information about a Mindustry server without any external dependency.

## CLI Usage
```bash
# Windows usage example
go-mindustry-ping --host 121.127.37.17 --refresh 5000

# Linux & MacOS usage examples
./go-mindustry-ping --host 121.127.37.17 --refresh 5000
```
* `--host <host>` - server host; (Default: 127.0.0.1)
* `--port <port>` - server port; (Default: 6567)
* `--raw` - show fields as-is (without formatting); (Default: false)
* `--refresh <ms>` - refresh interval, in ms (0 - no refresh); (Default: 0)
* `--indent <size>` - indent size; (Default: 3)
* `--no-ansi` - prevent printing ANSI-codes; (Default: false)

![Example of CLI usage](assets/cli-usage-1.webp)

## Direct API Usage
```go
package main

import (
	"log"
	"time"

	"github.com/altwine/go-mindustry-ping/pkg/serverinfo"
)

func main() {
	// Fetch server info
	si, err := serverinfo.GetServerInfo("omnidustry.ru", 6567)
	if err != nil {
		log.Fatalf("Error fetching server info: %v", err)
	}

	si.FormatFieldsAnsi() // Replace all color tags with their corresponding ANSI codes

	log.Printf("%d players are playing on the omnidustry server right now!", si.Players)

	// Wait for 5 seconds before updating server info
	time.Sleep(5 * time.Second)

	// Update the server info
	if err := si.Update(); err != nil {
		log.Fatalf("Error updating server info: %v", err)
	}

	log.Printf("%d players are playing on the omnidustry server 5 seconds later!", si.Players)
}
```

## Building
1. Install required tools: [Go compiler](https://go.dev/dl/)
2. Run `build.sh` in the `scripts` directory.
3. Find binaries in the `build` directory.

## Plans
This project is essentially done. No further plans, only bug fixes.

# License
MIT. See the [LICENSE](LICENSE.txt) file.
