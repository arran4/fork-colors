package web

import (
	"testing"
)

func TestWebLookup(t *testing.T) {
	c, ok := FromString("AliceBlue")
	if !ok || c != AliceBlue {
		t.Errorf("FromString(AliceBlue) failed")
	}

	name := ToString(AliceBlue)
	if name != "AliceBlue" {
		t.Errorf("ToString(AliceBlue) failed")
	}
}
