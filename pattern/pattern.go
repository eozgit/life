package pattern

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetPattern(patternName string) map[int]map[int]struct{} {
	absPath, _ := filepath.Abs(fmt.Sprintf("./pattern/%s.txt", patternName))

	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal("Unable to read pattern file", err)
	}
	defer file.Close()

	lineNo := 0
	result := make(map[int]map[int]struct{})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for pos, char := range line {
			if char == 'X' {
				if result[lineNo] == nil {
					result[lineNo] = make(map[int]struct{})
				}
				result[lineNo][pos] = struct{}{}
			}
		}
		lineNo++
	}

	return result
}
