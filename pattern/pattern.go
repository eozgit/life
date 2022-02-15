package pattern

import (
	"bufio"
	_ "embed"
	"strings"
)

//go:embed cell.txt
var cell string

//go:embed block.txt
var block string

//go:embed glider.txt
var glider string

//go:embed lwss.txt
var lwss string

//go:embed mwss.txt
var mwss string

//go:embed hwss.txt
var hwss string

var patternMap = map[string]string{
	"cell": cell, "block": block, "glider": glider, "lwss": lwss, "mwss": mwss, "hwss": hwss,
}

func getPattern(patternName string) [][]bool {
	text := strings.NewReader(patternMap[patternName])

	scanner := bufio.NewScanner(text)
	lineNo := 0
	result := make([][]bool, 0)
	for scanner.Scan() {
		for len(result) <= lineNo {
			result = append(result, make([]bool, 0))
		}
		line := scanner.Text()
		for pos, char := range line {
			for len(result[lineNo]) <= pos {
				result[lineNo] = append(result[lineNo], false)
			}
			result[lineNo][pos] = char == 'X'
		}
		lineNo++
	}
	return result
}

func GetPattern(patternName string) map[int]map[int]struct{} {
	pattern := getPattern(patternName)

	result := make(map[int]map[int]struct{})

	for i, val := range pattern {
		for j, alive := range val {
			if alive {
				if result[i] == nil {
					result[i] = make(map[int]struct{})
				}
				result[i][j] = struct{}{}
			}
		}
	}

	return result
}
