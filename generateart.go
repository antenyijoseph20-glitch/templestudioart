package main

import "strings"

func generateArt(input string, banner map[rune][]string) string{
	var result strings.Builder
	
	lines := splitArt(input)
	for _, line := range lines{
		// handles an input line
		if line == "" {
			result.WriteString("\n")
			continue
		}
		// draw the 8 rows of ASCII art
		for row := 0; row < 8; row++{
			for _, char := range line{
				result.WriteString(banner[char][row])
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
