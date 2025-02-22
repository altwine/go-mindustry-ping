package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/altwine/go-mindustry-ping/pkg/serverinfo"
)

var BINARY_VERSION string
var BINARY_ARCH string
var BINARY_OS string

func main() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		w.Write([]byte("Usage of go-mindustry-ping:\n"))
		flag.PrintDefaults()
	}

	host := flag.String("host", "127.0.0.1", "Host")
	port := flag.Int("port", 6567, "Port")
	raw := flag.Bool("raw", false, "Show fields as-is (Without formatting)")
	refresh := flag.Int("refresh", 0, "Refresh interval, in ms (0 - no refresh)")
	indentSize := flag.Int("indent", 3, "Indent size")
	noAnsi := flag.Bool("no-ansi", false, "Prevent printing ANSI-codes")
	version := flag.Bool("version", false, "Print the version number")

	flag.Parse()

	if *version {
		fmt.Printf("go-mindustry-ping %s %s/%s", BINARY_VERSION, BINARY_OS, BINARY_ARCH)
		return
	}

	printer := &Printer{IndentSize: *indentSize, NoAnsi: *noAnsi}
	si, err := serverinfo.GetServerInfo(*host, *port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		switch si.VerType {
		case "MindustryX":
			printer.Logo = MINDUSTRYX_LOGO
		case "official":
			printer.Logo = MINDUSTRY_LOGO
		default:
			printer.Logo = NO_LOGO
		}

		if !*raw {
			si.FormatFieldsAnsi()
		}

		printer.SetLine(0, fmt.Sprintf("%s\033[90m:\033[0m%d \033[90m(%dms)\033[0m", si.Address, si.Port, si.Latency))
		printer.SetLine(1, fmt.Sprintf("\033[90m%s\033[0m", strings.Repeat("⎯", 32)))
		printer.SetLine(2, fmt.Sprintf("\033[1mHost\033[0m\033[90m:\033[0m %s\033[0m", si.Host))
		printer.SetLine(3, fmt.Sprintf("\033[1mMap\033[0m\033[90m:\033[0m %s\033[0m", si.Map))
		if si.Limit != 0 {
			printer.SetLine(4, fmt.Sprintf("\033[1mPlayers\033[0m\033[90m:\033[0m %d/%d\033[0m", si.Players, si.Limit))
		} else {
			printer.SetLine(4, fmt.Sprintf("\033[1mPlayers\033[0m\033[90m:\033[0m %d\033[0m", si.Players))
		}
		printer.SetLine(5, fmt.Sprintf("\033[1mWaves\033[0m\033[90m:\033[0m %d\033[0m", si.Waves))
		printer.SetLine(6, fmt.Sprintf("\033[1mVersion\033[0m\033[90m:\033[0m %d\033[0m", si.GameVersion))
		printer.SetLine(7, fmt.Sprintf("\033[1mType\033[0m\033[90m:\033[0m %s\033[0m", si.VerType))
		printer.SetLine(8, fmt.Sprintf("\033[1mGamemode\033[0m\033[90m:\033[0m %s\033[0m", si.Gamemode))
		printer.SetLine(9, fmt.Sprintf("\033[1mDescription\033[0m\033[90m:\033[0m %s\033[0m", si.Desc))
		printer.Update()

		if *refresh == 0 {
			break
		}

		time.Sleep(time.Millisecond * time.Duration(*refresh))
		err := si.Update()
		if err != nil {
			log.Fatal(err)
		}
	}
}
