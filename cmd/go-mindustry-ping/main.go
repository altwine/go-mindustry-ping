package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/altwine/go-mindustry-ping/pkg/serverinfo"
)

const (
	DEFAULT_HOST = "127.0.0.1"
	DEFAULT_PORT = 6567
)

var BINARY_VERSION string

func handleExitRequests(callback func()) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		callback()
		os.Exit(0)
	}()
}

func main() {
	address := flag.String("host", DEFAULT_HOST, "Host")
	port := flag.Int("port", DEFAULT_PORT, "Port")
	colorize := flag.Bool("colorize", false, "Visualize color tags")
	refreshInterval := flag.Int("refresh", 0, "Refresh interval, in ms (0 - no refresh)")

	flag.Parse()

	iteration := 0
	exitCallback := func() {
		fmt.Printf("\033[?25h")
		if iteration > 0 {
			fmt.Printf("%s\033[%dA", strings.Repeat("\n", 12), 12)
		}
	}
	handleExitRequests(exitCallback)

	for {
		if iteration != 0 {
			fmt.Printf("\033[%dA", 12)
		}
		si, err := serverinfo.GetServerInfo(*address, *port)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if *colorize {
			si.ColorizeFields()
		}

		indent := strings.Repeat(" ", 3)

		// Made with https://dom111.github.io/image-to-ansi
		logo1 := "\033[38;5;243;48;5;247m▄\033[38;5;247;48;5;250m▄\033[48;5;250m     \033[48;5;222m \033[38;5;250;48;5;222m▄▄▄▄\033[48;5;222m \033[48;5;250m       \033[m"
		logo2 := "\033[48;5;243m  \033[48;5;247m \033[38;5;250;48;5;247m▄\033[38;5;250;48;5;248m▄\033[38;5;250;48;5;247m▄▄\033[38;5;180;48;5;180m▄▄\033[38;5;180;48;5;247m▄▄\033[38;5;180;48;5;180m▄▄\033[38;5;250;48;5;247m▄▄\033[38;5;250;48;5;248m▄\033[38;5;250;48;5;247m▄\033[48;5;247m \033[48;5;250m  \033[m"
		logo3 := "\033[48;5;243m  \033[38;5;247;48;5;248m▄\033[38;5;250;48;5;250m▄\033[38;5;242;48;5;242m▄▄\033[38;5;248;48;5;249m▄\033[38;5;248;48;5;246m▄\033[38;5;248;48;5;179m▄\033[38;5;179;48;5;179m▄▄\033[38;5;248;48;5;179m▄\033[38;5;248;48;5;246m▄\033[38;5;248;48;5;249m▄\033[38;5;242;48;5;242m▄▄\033[38;5;250;48;5;250m▄\033[38;5;247;48;5;248m▄\033[48;5;250m  \033[m"
		logo4 := "\033[38;5;173;48;5;243m▄▄\033[38;5;180;48;5;247m▄\033[38;5;180;48;5;249m▄\033[38;5;246;48;5;249m▄\033[38;5;247;48;5;248m▄\033[38;5;95;48;5;247m▄\033[48;5;95m  \033[38;5;137;48;5;8m▄▄\033[48;5;180m  \033[38;5;180;48;5;247m▄\033[38;5;247;48;5;248m▄\033[38;5;246;48;5;249m▄\033[38;5;180;48;5;249m▄\033[38;5;180;48;5;247m▄\033[38;5;222;48;5;250m▄▄\033[m"
		logo5 := "\033[48;5;173m \033[48;5;243m \033[38;5;247;48;5;180m▄\033[38;5;179;48;5;180m▄\033[38;5;179;48;5;179m▄\033[38;5;137;48;5;246m▄\033[48;5;95m \033[38;5;137;48;5;95m▄\033[38;5;180;48;5;137m▄\033[38;5;95;48;5;95m▄\033[38;5;180;48;5;95m▄\033[38;5;95;48;5;138m▄\033[38;5;138;48;5;180m▄\033[48;5;180m \033[38;5;137;48;5;246m▄\033[38;5;179;48;5;179m▄\033[38;5;179;48;5;180m▄\033[38;5;247;48;5;180m▄\033[38;5;250;48;5;250m▄\033[48;5;222m \033[m"
		logo6 := "\033[48;5;173m \033[48;5;243m \033[38;5;180;48;5;247m▄\033[38;5;179;48;5;179m▄▄\033[38;5;245;48;5;137m▄\033[38;5;180;48;5;137m▄\033[38;5;250;48;5;180m▄\033[38;5;250;48;5;95m▄\033[48;5;95m \033[48;5;180m \033[38;5;255;48;5;180m▄\033[38;5;255;48;5;95m▄\033[38;5;95;48;5;138m▄\033[38;5;245;48;5;137m▄\033[38;5;179;48;5;179m▄▄\033[38;5;180;48;5;247m▄\033[48;5;250m \033[48;5;222m \033[m"
		logo7 := "\033[38;5;243;48;5;173m▄▄\033[38;5;247;48;5;180m▄\033[38;5;249;48;5;180m▄\033[38;5;145;48;5;245m▄\033[38;5;247;48;5;246m▄\033[38;5;245;48;5;95m▄\033[38;5;95;48;5;249m▄\033[38;5;249;48;5;250m▄\033[48;5;250m \033[48;5;255m \033[38;5;255;48;5;255m▄\033[38;5;180;48;5;255m▄\033[38;5;245;48;5;180m▄\033[38;5;247;48;5;246m▄\033[38;5;145;48;5;245m▄\033[38;5;249;48;5;180m▄\033[38;5;247;48;5;180m▄\033[38;5;250;48;5;222m▄▄\033[m"
		logo8 := "\033[48;5;243m  \033[38;5;248;48;5;247m▄\033[38;5;250;48;5;250m▄\033[38;5;242;48;5;242m▄▄\033[38;5;145;48;5;247m▄\033[38;5;245;48;5;246m▄\033[38;5;179;48;5;95m▄\033[38;5;173;48;5;250m▄\033[38;5;173;48;5;255m▄\033[38;5;179;48;5;180m▄\033[38;5;245;48;5;246m▄\033[38;5;145;48;5;247m▄\033[38;5;242;48;5;242m▄▄\033[38;5;250;48;5;250m▄\033[38;5;248;48;5;247m▄\033[48;5;250m  \033[m"
		logo9 := "\033[48;5;243m  \033[48;5;247m \033[38;5;247;48;5;250m▄\033[38;5;248;48;5;250m▄\033[38;5;247;48;5;250m▄\033[38;5;247;48;5;249m▄\033[38;5;180;48;5;179m▄▄\033[38;5;246;48;5;179m▄▄\033[38;5;180;48;5;179m▄▄\033[38;5;247;48;5;249m▄\033[38;5;247;48;5;250m▄\033[38;5;248;48;5;250m▄\033[38;5;247;48;5;250m▄\033[48;5;247m \033[48;5;250m  \033[m"
		logo10 := "\033[48;5;243m       \033[38;5;173;48;5;173m▄\033[38;5;173;48;5;242m▄▄▄▄\033[38;5;173;48;5;173m▄\033[48;5;243m     \033[38;5;243;48;5;247m▄\033[38;5;247;48;5;250m▄\033[m"

		fmt.Printf("\n%s%s%s\033[1mHost\033[0m: %s\n", indent, logo1, indent, si.Host)
		fmt.Printf("%s%s%s\033[1mMap\033[0m: %s\n", indent, logo2, indent, si.Map)
		if si.Limit != 0 {
			fmt.Printf("%s%s%s\033[1mPlayers\033[0m: %d/%d\n", indent, logo3, indent, si.Players, si.Limit)
		} else {
			fmt.Printf("%s%s%s\033[1mPlayers\033[0m: %d\n", indent, logo3, indent, si.Players)
		}
		fmt.Printf("%s%s%s\033[1mWaves\033[0m: %d\n", indent, logo4, indent, si.Waves)
		fmt.Printf("%s%s%s\033[1mVersion\033[0m: %d\n", indent, logo5, indent, si.GameVersion)
		fmt.Printf("%s%s%s\033[1mType\033[0m: %s\n", indent, logo6, indent, si.VerType)
		fmt.Printf("%s%s%s\033[1mGamemode\033[0m: %v\n", indent, logo7, indent, si.Gamemode)
		fmt.Printf("%s%s%s\033[1mDesc\033[0m: %s\n", indent, logo8, indent, si.Desc)
		fmt.Printf("%s%s%s\033[1mLatency\033[0m: %dms\n", indent, logo9, indent, si.Latency)
		fmt.Printf("%s%s\n\n\033[?25l", indent, logo10)

		if *refreshInterval == 0 {
			break
		}
		time.Sleep(time.Millisecond * time.Duration(*refreshInterval))
		iteration += 1
	}
	exitCallback()
}
