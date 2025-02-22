package serverinfo

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type MindustryColor struct {
	R, G, B, A int
}

var MINDUSTRY_COLORS = map[string]MindustryColor{
	"white":      {255, 255, 255, 255},
	"lightGray":  {191, 191, 191, 255},
	"gray":       {127, 127, 127, 255},
	"darkGray":   {63, 63, 63, 255},
	"black":      {0, 0, 0, 255},
	"clear":      {0, 0, 0, 255},
	"blue":       {0, 0, 255, 255},
	"navy":       {0, 0, 128, 255},
	"royal":      {65, 105, 225, 255},
	"slate":      {112, 128, 144, 255},
	"sky":        {135, 206, 235, 255},
	"cyan":       {0, 255, 255, 255},
	"teal":       {0, 128, 128, 255},
	"green":      {0, 255, 0, 255},
	"acid":       {127, 255, 0, 255},
	"lime":       {50, 205, 50, 255},
	"forest":     {34, 139, 34, 255},
	"olive":      {107, 142, 35, 255},
	"yellow":     {255, 255, 0, 255},
	"gold":       {221, 185, 0, 255},
	"goldenrod":  {218, 165, 32, 255},
	"orange":     {165, 42, 42, 255},
	"brown":      {139, 69, 51, 255},
	"tan":        {210, 176, 157, 255},
	"brick":      {178, 34, 34, 255},
	"red":        {255, 0, 0, 255},
	"scarlet":    {207, 51, 51, 255},
	"crimson":    {220, 69, 63, 255},
	"coral":      {255, 99, 71, 255},
	"salmon":     {250, 128, 114, 255},
	"pink":       {255, 173, 216, 255},
	"magenta":    {255, 0, 255, 255},
	"purple":     {160, 32, 192, 255},
	"violet":     {238, 130, 238, 255},
	"maroon":     {176, 32, 64, 255},
	"accent":     {255, 203, 57, 255},
	"unlaunched": {137, 130, 237, 255},
	"stat":       {255, 211, 127, 255},
}

func getColorFromTag(tag string) (MindustryColor, bool) {
	tagTrimmed := tag[1 : len(tag)-1]
	color, found := MINDUSTRY_COLORS[tagTrimmed]
	if found {
		return color, true
	}
	if strings.HasPrefix(tagTrimmed, "#") && len(tagTrimmed) <= 9 && len(tagTrimmed) > 1 {
		newTag := tagTrimmed + strings.Repeat("f", 9-len(tagTrimmed))
		rr, err1 := strconv.ParseInt(newTag[1:3], 16, 16)
		gg, err2 := strconv.ParseInt(newTag[3:5], 16, 16)
		bb, err3 := strconv.ParseInt(newTag[5:7], 16, 16)
		_, err4 := strconv.ParseInt(newTag[7:9], 16, 16)
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			return MindustryColor{}, false
		}
		mindustryColor := MindustryColor{R: int(rr), G: int(gg), B: int(bb)}
		return mindustryColor, true
	}
	return MindustryColor{}, false
}

func (si *ServerInfo) forEachField(callback func(string) string) {
	si.Address = callback(si.Address)
	si.Host = callback(si.Host)
	si.Map = callback(si.Map)
	si.VerType = callback(si.VerType)
	si.Gamemode = callback(si.Gamemode)
	si.Desc = callback(si.Desc)
}

func (si *ServerInfo) forEachTag(callback func(string) string) {
	re := regexp.MustCompile(`\[.*?\]`)
	si.Address = re.ReplaceAllStringFunc(si.Address, callback)
	si.Host = re.ReplaceAllStringFunc(si.Host, callback)
	si.Map = re.ReplaceAllStringFunc(si.Map, callback)
	si.VerType = re.ReplaceAllStringFunc(si.VerType, callback)
	si.Gamemode = re.ReplaceAllStringFunc(si.Gamemode, callback)
	si.Desc = re.ReplaceAllStringFunc(si.Desc, callback)
}

func (si *ServerInfo) FormatFieldsHtml() {
	log.Fatal("FormatField for HTML is not implemented yet.")
}

func (si *ServerInfo) FormatFieldsAnsi() {
	lastColor := MindustryColor{255, 255, 255, 255}
	si.forEachTag(func(tag string) string {
		tagContent := tag[1 : len(tag)-1]
		if tagContent == "" {
			return fmt.Sprintf("\033[38;2;%d;%d;%dm", lastColor.R, lastColor.G, lastColor.B)
		}
		color, found := getColorFromTag(tag)
		if found {
			lastColor = color
			return fmt.Sprintf("\033[38;2;%d;%d;%dm", color.R, color.G, color.B)
		}
		return tag
	})
}
