package main

import "testing"

func TestRenderArtValid(t *testing.T) {
	result, err := renderArt("A", "bannerfiles/standard.txt")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == "" {
		t.Error("Expected ASCII art, got an empty string")
	}
}
func TestRenderArtInvalidBanner(t *testing.T) {
	_, err := renderArt("A", "bannerfiles/notfound.txt")

	if err == nil {
		t.Error("Expected an error for an invalid banner file")
	}
}
func TestRenderArtEmptyInput(t *testing.T) {
	result, err := renderArt("", "bannerfiles/standard.txt")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != "" {
		t.Errorf("Expected empty output, got %q", result)
	}
}
func TestRenderArtMultiLine(t *testing.T) {
	result, err := renderArt("A\nB", "bannerfiles/standard.txt")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result == "" {
		t.Error("Expected ASCII art for multiple lines")
	}
}
