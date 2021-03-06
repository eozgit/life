package game

import (
	"fmt"
	"strings"
)

type MenuItem struct {
	hotkey      string
	description string
}

func getHelpText() string {
	hotkeyFnMap := []MenuItem{
		{hotkey: "1-9", description: "set speed"},
		{hotkey: "t + 1-4", description: "select theme"},
		{hotkey: "LMB", description: "resurrect cell"},
		{hotkey: "x + LMB", description: "create glider"},
		{hotkey: "r + 1-9", description: "reset"},
		{hotkey: "h", description: "resume"},
	}
	max := 0
	for _, item := range hotkeyFnMap {
		length := len(item.hotkey)
		if length > max {
			max = length
		}
	}
	format := fmt.Sprintf("%%-%ds%%s", max+2)
	lines := []string{}
	for _, item := range hotkeyFnMap {
		line := fmt.Sprintf(format, item.hotkey, item.description)
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

var helptext = getHelpText()
