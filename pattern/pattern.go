package pattern

import (
	"bufio"
	_ "embed"
	"math/rand"
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

func flipCoin() bool {
	return rand.Float32() < .5
}

func transformPattern(pattern [][]bool) [][]bool {
	length := len(pattern)

	flipX, flipY, transpose := flipCoin(), flipCoin(), flipCoin()

	result := make([][]bool, length)
	for i := 0; i < length; i++ {
		result[i] = make([]bool, length)
	}

	for i := 0; i < length; i++ {
		m := i
		if flipX {
			m = length - i - 1
		}
		for j := 0; j < length; j++ {
			n := j
			if flipY {
				n = length - j - 1
			}

			if transpose {
				result[n][m] = pattern[i][j]
			} else {
				result[m][n] = pattern[i][j]
			}
		}
	}

	return result
}

func GetPattern(patternName string) map[int]map[int]struct{} {
	pattern := getPattern(patternName)

	pattern = transformPattern(pattern)

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
