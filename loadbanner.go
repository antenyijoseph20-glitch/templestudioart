// here is my final code below
package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(fileName string) (map[rune][]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		//fmt.Println("banner loading files")
		return nil, err
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")
	if len(lines) < 856 {
		return nil, fmt.Errorf("invalid banner file")
	}
	//input := strings.Split(lines, "\n")

	charmap := make(map[rune][]string)
	for i := 0; i < 95; i++ {
		char := rune(i + 32)
		start := i*9 + 1
		charmap[char] = lines[start : start+8]
	}
	return charmap, nil
}
