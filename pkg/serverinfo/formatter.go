package serverinfo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type MindustryColor struct {
	R, G, B, A int
}

var MINDUSTRY_COLORS = map[string]MindustryColor{
	"ACID":       {127, 255, 0, 255},
	"BLACK":      {0, 0, 0, 255},
	"BLUE":       {0, 0, 255, 255},
	"BRICK":      {178, 34, 34, 255},
	"BROWN":      {139, 69, 19, 255},
	"CLEAR":      {0, 0, 0, 0},
	"CORAL":      {255, 127, 80, 255},
	"CRIMSON":    {220, 20, 60, 255},
	"CYAN":       {0, 255, 255, 255},
	"DARK_GRAY":  {63, 63, 63, 255},
	"DARK_GREY":  {63, 63, 63, 255},
	"FOREST":     {34, 139, 34, 255},
	"GOLD":       {255, 215, 0, 255},
	"GOLDENROD":  {218, 165, 32, 255},
	"GRAY":       {127, 127, 127, 255},
	"GREEN":      {0, 255, 0, 255},
	"GREY":       {127, 127, 127, 255},
	"LIGHT_GRAY": {191, 191, 191, 255},
	"LIGHT_GREY": {191, 191, 191, 255},
	"LIME":       {50, 205, 50, 255},
	"MAGENTA":    {255, 0, 255, 255},
	"MAROON":     {176, 48, 96, 255},
	"NAVY":       {0, 0, 127, 255},
	"OLIVE":      {107, 142, 35, 255},
	"ORANGE":     {255, 165, 0, 255},
	"PINK":       {255, 105, 180, 255},
	"PURPLE":     {160, 32, 240, 255},
	"RED":        {255, 0, 0, 255},
	"ROYAL":      {65, 105, 225, 255},
	"SALMON":     {250, 128, 114, 255},
	"SCARLET":    {255, 52, 28, 255},
	"SKY":        {135, 206, 235, 255},
	"SLATE":      {112, 128, 144, 255},
	"TAN":        {210, 180, 140, 255},
	"TEAL":       {0, 127, 127, 255},
	"VIOLET":     {238, 130, 238, 255},
	"WHITE":      {255, 255, 255, 255},
	"YELLOW":     {255, 255, 0, 255},
	"acid":       {127, 255, 0, 255},
	"black":      {0, 0, 0, 255},
	"blue":       {0, 0, 255, 255},
	"brick":      {178, 34, 34, 255},
	"brown":      {139, 69, 19, 255},
	"clear":      {0, 0, 0, 0},
	"coral":      {255, 127, 80, 255},
	"crimson":    {220, 20, 60, 255},
	"cyan":       {0, 255, 255, 255},
	"darkgray":   {63, 63, 63, 255},
	"darkgrey":   {63, 63, 63, 255},
	"forest":     {34, 139, 34, 255},
	"gold":       {255, 215, 0, 255},
	"goldenrod":  {218, 165, 32, 255},
	"gray":       {127, 127, 127, 255},
	"green":      {0, 255, 0, 255},
	"grey":       {127, 127, 127, 255},
	"lightgray":  {191, 191, 191, 255},
	"lightgrey":  {191, 191, 191, 255},
	"lime":       {50, 205, 50, 255},
	"magenta":    {255, 0, 255, 255},
	"maroon":     {176, 48, 96, 255},
	"navy":       {0, 0, 127, 255},
	"olive":      {107, 142, 35, 255},
	"orange":     {255, 165, 0, 255},
	"pink":       {255, 105, 180, 255},
	"purple":     {160, 32, 240, 255},
	"red":        {255, 0, 0, 255},
	"royal":      {65, 105, 225, 255},
	"salmon":     {250, 128, 114, 255},
	"scarlet":    {255, 52, 28, 255},
	"sky":        {135, 206, 235, 255},
	"slate":      {112, 128, 144, 255},
	"tan":        {210, 180, 140, 255},
	"teal":       {0, 127, 127, 255},
	"violet":     {238, 130, 238, 255},
	"white":      {255, 255, 255, 255},
	"yellow":     {255, 255, 0, 255},
	"accent":     {255, 211, 127, 255},
	"unlaunched": {137, 130, 237, 255},
	"highlight":  {255, 224, 165, 255},
	"stat":       {255, 211, 127, 255},
	"negstat":    {229, 84, 84, 255},
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
	si.ModeName = callback(si.ModeName)
	si.Desc = callback(si.Desc)
}

func (si *ServerInfo) forEachTag(callback func(string) string) {
	re := regexp.MustCompile(`\[.*?\]`)
	si.Address = re.ReplaceAllStringFunc(si.Address, callback)
	si.Host = re.ReplaceAllStringFunc(si.Host, callback)
	si.Map = re.ReplaceAllStringFunc(si.Map, callback)
	si.VerType = re.ReplaceAllStringFunc(si.VerType, callback)
	si.Gamemode = re.ReplaceAllStringFunc(si.Gamemode, callback)
	si.ModeName = re.ReplaceAllStringFunc(si.ModeName, callback)
	si.Desc = re.ReplaceAllStringFunc(si.Desc, callback)
}

func (si *ServerInfo) FormatFieldsHtml() {
	processField := func(field string) string {
		if field == "" {
			return field
		}

		re := regexp.MustCompile(`\[(.*?)\]`)

		var result strings.Builder
		lastIndex := 0
		lastTagWasColor := false

		matches := re.FindAllStringSubmatchIndex(field, -1)

		for _, match := range matches {
			if match[0] > lastIndex {
				result.WriteString(field[lastIndex:match[0]])
			}

			tagContent := field[match[2]:match[3]]

			if tagContent == "" {
				if lastTagWasColor {
					result.WriteString(`</span>`)
					lastTagWasColor = false
				}
			} else {
				tag := fmt.Sprintf("[%s]", tagContent)
				color, found := getColorFromTag(tag)
				if found {
					if lastTagWasColor {
						result.WriteString(`</span>`)
					}
					hexColor := fmt.Sprintf("#%02X%02X%02X", color.R, color.G, color.B)
					result.WriteString(fmt.Sprintf(`<span style="color: %s;">`, hexColor))
					lastTagWasColor = true
				} else {
					if lastTagWasColor {
						result.WriteString(`</span>`)
						lastTagWasColor = false
					}
					result.WriteString(tag)
				}
			}

			lastIndex = match[1]
		}

		if lastIndex < len(field) {
			result.WriteString(field[lastIndex:])
		}

		if lastTagWasColor {
			result.WriteString(`</span>`)
		}

		return result.String()
	}

	si.Address = processField(si.Address)
	si.Host = processField(si.Host)
	si.Map = processField(si.Map)
	si.VerType = processField(si.VerType)
	si.Gamemode = processField(si.Gamemode)
	si.ModeName = processField(si.ModeName)
	si.Desc = processField(si.Desc)
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
