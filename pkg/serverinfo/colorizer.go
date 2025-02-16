package serverinfo

import (
	"fmt"
	"regexp"
	"strconv"
)

// TODO: add support for rrggbbaa (simulate alpha channel by applying some gray tint maybe?)
func hexToANSI(hex string) string {
	if len(hex) == 7 && hex[0] == '#' {
		hex = hex[1:]
	}

	r, _ := strconv.ParseInt(hex[0:2], 16, 64)
	g, _ := strconv.ParseInt(hex[2:4], 16, 64)
	b, _ := strconv.ParseInt(hex[4:6], 16, 64)

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

// TODO: add support for more colors: https://mindustrygame.github.io/wiki/modding/5-markup/
func replaceColorCodes(input string) string {
	re := regexp.MustCompile(`\[#([0-9a-fA-F]{6})\]`)
	newStr := re.ReplaceAllStringFunc(input, func(match string) string {
		hex := match[2 : len(match)-1]
		ansiCode := hexToANSI(hex)
		return ansiCode
	})
	newStr += "\033[0m"
	return newStr
}

func (si *ServerInfo) ColorizeFields() {
	si.Host = replaceColorCodes(si.Host)
	si.Map = replaceColorCodes(si.Map)
	si.VerType = replaceColorCodes(si.VerType)
	si.Desc = replaceColorCodes(si.Desc)
}
