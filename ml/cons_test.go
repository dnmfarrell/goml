package ml

import (
	"testing"
)

func TestCons(t *testing.T) {
	xs := Cons[int]{0, Cons[int]{1, Cons[int]{3, Nil[int]{}}}}
	if xs.Empty() {
		t.Error("Cons.Empty is true")
	}
	if xs.Lazy() {
		t.Error("Cons.Lazy is true")
	}
	if xs.Head() != 0 {
		t.Errorf("Head() returns %d", xs.Head())
	}
	if xs.Tail().Head() != 1 {
		t.Errorf("Tail().Head() returns %d", xs.Tail().Head())
	}
	if xs.Last() != 3 {
		t.Errorf("[0,1,3].Last returns %d", xs.Last())
	}
}
