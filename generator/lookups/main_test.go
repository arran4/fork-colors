package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseColors(t *testing.T) {
	// Create a temporary Go file with some color definitions.
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "colors_test.go")

	content := `package testpkg

import "image/color"

var (
	Transparent = color.RGBA{0, 0, 0, 0}
	Red         = color.RGBA{255, 0, 0, 255}
	Green       = color.RGBA{0, 255, 0, 255}
	Blue        = color.RGBA{0, 0, 255, 255}
)

var OtherVar = "hello"
`
	if err := os.WriteFile(tempFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}

	colors, err := parseColors(tempFile)
	if err != nil {
		t.Fatalf("parseColors returned error: %v", err)
	}

	expected := []Color{
		{Name: "Transparent"},
		{Name: "Red"},
		{Name: "Green"},
		{Name: "Blue"},
		{Name: "OtherVar"},
	}

	if !reflect.DeepEqual(colors, expected) {
		t.Errorf("parseColors() returned %v, expected %v", colors, expected)
	}
}
