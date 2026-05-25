package colors

import (
	"image/color"
	"testing"

	"github.com/pborman/colors/crayola"
	"github.com/pborman/colors/web"
)

func TestFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected color.RGBA
		found    bool
	}{
		// Root exact match
		{"root exact", "AbsoluteZero", AbsoluteZero, true},
		// Root case insensitive match
		{"root lower", "absolutezero", AbsoluteZero, true},
		// Subpackage match (web)
		{"web specific", "AliceBlue", web.AliceBlue, true}, // assuming it's in web
		// Missing
		{"missing", "NonExistentColor123", color.RGBA{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, ok := FromString(tt.input)
			if ok != tt.found {
				t.Errorf("FromString(%q) found = %v, want %v", tt.input, ok, tt.found)
			}
			if c != tt.expected {
				t.Errorf("FromString(%q) color = %v, want %v", tt.input, c, tt.expected)
			}
		})
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		name     string
		input    color.RGBA
		expected string
	}{
		{"root color", AbsoluteZero, "AbsoluteZero"},
		{"crayola specific", crayola.OrangeRed, "OrangeRed"}, // it might match a root or web color if collisions exist, but we just check if it finds *some* name
		{"missing", color.RGBA{0, 1, 2, 3}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := ToString(tt.input)
			if tt.expected != "" && name == "" {
				t.Errorf("ToString(%v) returned empty, expected something", tt.input)
			}
			if tt.expected == "" && name != "" {
				t.Errorf("ToString(%v) returned %q, expected empty", tt.input, name)
			}
			// if it's the root color, it should perfectly match
			if tt.name == "root color" && name != tt.expected {
				t.Errorf("ToString(%v) = %q, want %q", tt.input, name, tt.expected)
			}
		})
	}
}
