package main

import (
	"strings"
	// //"fmt"
)

func renderArt(input string, bannerFile string) (string, error) {
	if input == "" {
		return "", nil
	}

	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.ReplaceAll(input, "\r", "\n")
	// Is the input valid?
	_, err := validateArt(input)
	if err != nil {
		return "", err
	}
	// Can we load the banner?
	banner, err := loadBanner(bannerFile)
	if err != nil {
		return "", err
	}
	// Build the ASCII art.
	result := generateArt(input, banner)

	return result, nil

}
