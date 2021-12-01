package ml

import (
	"testing"
)

func TestConsLFilterInf(t *testing.T) {
	xs := RangeInf[int]{0, 1}.Eval()
	if xs.Head() != 0 {
		t.Errorf("Head() returns %d", xs.Head())
	}
	if xs.Tail().Head() != 1 {
		t.Errorf("Tail().Head() returns %d", xs.Tail().Head())
	}
	if xs.Tail().Head() != 1 {
		t.Errorf("Tail().Head() twice returns %d", xs.Tail().Head())
	}
	evens := xs.Filter(func(x int) bool { return x%2 == 0 }).Take(3)
	if evens.Length() != 3 {
		t.Errorf("Take returns the wrong sized list: %d", evens.Length())
	}
	if evens.Head()%2 != 0 {
		t.Errorf("evens.Head is not even: %d", evens.Head())
	}
	if evens.Tail().Head()%2 != 0 {
		t.Errorf("evens.Tail.Head is not even: %d", evens.Tail().Head())
	}
	if evens.Last()%2 != 0 {
		t.Errorf("evens.Last is not even: %d", evens.Last())
	}
}

func TestConsLFilterHalt(t *testing.T) {
	xs := Range[int]{0, 5, 25}.Eval()
	if xs.Length() != 6 {
		t.Errorf("Range{0,5,25}.Tail.Length: %d [%v]", xs.Length(), xs.Take(xs.Length()))
	}
}
