package main

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestParseColors(t *testing.T) {
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
	fs := fstest.MapFS{
		"colors_test.go": &fstest.MapFile{
			Data: []byte(content),
		},
	}

	file, err := fs.Open("colors_test.go")
	if err != nil {
		t.Fatalf("Failed to open file from MapFS: %v", err)
	}
	defer file.Close()

	colors, err := parseColors("colors_test.go", file)
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
