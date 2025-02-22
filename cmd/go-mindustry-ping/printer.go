package main

import (
	"fmt"
	"regexp"
	"strings"
)

var ansiFilter = regexp.MustCompile(`[\x1b\x9b][[()#;?]*(?:[0-9]{1,4}(?:;[0-9]{0,4})*)?[0-9A-ORZcf-nqry=><]`)

type Printer struct {
	Logo          [10]string
	IndentSize    int
	NoAnsi        bool
	lines         [64]string
	lastLineIndex int
	printedLines  int
}

func (p *Printer) Clear() {
	fmt.Print(strings.Repeat("\033[A", p.printedLines))
	p.printedLines = 0
}

func (p *Printer) Update() {
	p.Clear()
	lines := p.lines[:p.lastLineIndex+1]
	indent := strings.Repeat(" ", p.IndentSize)
	noLogoIndent := strings.Repeat(" ", 20+p.IndentSize*2)
	text := "\n"
	p.printedLines += len(lines) + 2
	for lineIndex, line := range lines {
		if lineIndex < 10 {
			text += fmt.Sprintf("%s%s\n\033[0m", indent+p.Logo[lineIndex]+indent, line)
		} else {
			text += fmt.Sprintf("%s%s\n\033[0m", noLogoIndent, line)
		}
	}
	text += "\n"
	if p.NoAnsi {
		text = ansiFilter.ReplaceAllString(text, "")
	}
	fmt.Print(text)
}

func (p *Printer) SetLine(line int, content string) {
	p.lines[line] = content
	p.lastLineIndex = max(line, p.lastLineIndex)
}
