# go-mindustry-ping v2.0.0
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
	"fmt"
	"log"
	"time"

	"github.com/altwine/go-mindustry-ping/pkg/serverinfo"
)

func main() {
	si, err := serverinfo.GetServerInfo("omnidustry.ru", 6567)
	if err != nil {
		log.Fatal(err)
	}
	si.FormatFieldsAnsi() // Replace all color tags with their corresponding ANSI codes
	fmt.Printf("%d players are playing on the omnidustry server right now!!", si.Players)
	time.Sleep(time.Second * 5)
	err := si.Update()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d players is playing on the omnidustry 5 seconds later!!", si.Players)
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
