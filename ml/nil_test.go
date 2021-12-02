package ml

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	n := Nil[string]{}
	if !n.Empty() {
		t.Errorf("[] is not empty")
	}
	if n.Lazy() {
		t.Errorf("[] is lazy")
	}
}
