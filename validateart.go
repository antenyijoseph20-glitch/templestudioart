package main

import "fmt"

func validateArt(input string) (rune, error) {
	for _, char := range input{
		if char == '\n'{
			continue
		}
		if char < 32 || char > 126{
			return char, fmt.Errorf("invalid character: %q", char)
		}
	}
	return 0, nil
}
