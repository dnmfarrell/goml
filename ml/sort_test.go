package ml

import (
	"testing"
)

func TestQsort(t *testing.T) {
	xs := Cons[int]{8, Cons[int]{3, Cons[int]{12, Nil[int]{}}}}
	ys := Qsort[int](xs)
	if ys.Head() != 3 {
		t.Errorf("Qsort [8,3,12] Head is: %d", ys.Head())
	}
	if ys.Tail().Head() != 8 {
		t.Errorf("Qsort [8,3,12] Tail.Head is: %d", ys.Head())
	}
	if ys.Last() != 12 {
		t.Errorf("Qsort [8,3,12] Last is: %d", ys.Last())
	}
}

func TestQsortR(t *testing.T) {
	xs := Cons[int]{8, Cons[int]{3, Cons[int]{12, Nil[int]{}}}}
	ys := QsortR[int](xs)
	if ys.Head() != 12 {
		t.Errorf("QsortR [8,3,12] Head is: %d", ys.Head())
	}
	if ys.Tail().Head() != 8 {
		t.Errorf("QsortR [8,3,12] Tail.Head is: %d", ys.Head())
	}
	if ys.Last() != 3 {
		t.Errorf("QsortR [8,3,12] Last is: %d", ys.Last())
	}
}
