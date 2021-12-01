package ml

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	n := Nil[string]{}
	if !n.Empty() {
		t.Errorf("[] is not empty")
	}
	if n.Length() != 0 {
		t.Errorf("[] length: %d", n.Length())
	}
	nSharp := n.Concat(Nil[string]{})
	if !nSharp.Empty() {
		t.Errorf("concat []' is not empty")
	}
	nRev := n.Reverse()
	if !nRev.Empty() {
		t.Errorf("reverse [] is not empty")
	}
	bigInt := func(x int) bool { return x > 7 }
	nFtrd := Nil[int]{}.Filter(bigInt)
	if !nFtrd.Empty() {
		t.Errorf("filter [] is not empty")
	}
}