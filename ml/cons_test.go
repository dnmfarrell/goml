package ml

import (
	"testing"
)

func TestCons(t *testing.T) {
	x := Cons[int]{3, Nil[int]{}}
	if x.Length() != 1 {
		t.Errorf("[x] length: %d", x.Length())
	}
	xs := Cons[int]{3, x}
	if xs.Length() != 2 {
		t.Errorf("[xx] length: %d", xs.Length())
	}
	got := xs.Drop(2)
	if !got.Empty() {
		t.Error("Drop 2 didn't return []")
	}
	got = xs.Take(2)
	if got.Length() != 2 {
		t.Errorf("Take 2 didn't return [xs]: %d", got.Length())
	}
	if xs.Last() != 3 {
		t.Errorf("Last returns %d instead of 3", xs.Last())
	}
	got = xs.Init()
	if got.Length() != 1 {
		t.Errorf("[xs] init length: %d", got.Length())
	}
	got = xs.Concat(Cons[int]{4, Nil[int]{}})
	if got.Length() != 3 {
		t.Errorf("[xs]++[x] length: %d", got.Length())
	}
}

func TestConsReverse(t *testing.T) {
	xs := Cons[string]{"foo", Cons[string]{"bar", Cons[string]{"baz", Nil[string]{}}}}
	ys := xs.Reverse()
	if ys.Head() != "baz" {
		t.Errorf("[foo,bar,baz].Reverse.Head returns: %s", ys.Head())
	}
	if ys.Last() != "foo" {
		t.Errorf("[foo,bar,baz].Reverse.Last returns: %s", ys.Last())
	}
}

func TestConsFilter(t *testing.T) {
	xs := Cons[int]{6, Cons[int]{7, Cons[int]{8, Nil[int]{}}}}
	bigInt := func(x int) bool { return x > 7 }
	ys := xs.Filter(bigInt)
	if ys.Head() != 8 {
		t.Errorf("[6,7,8].Filter >7.Head returns: %d", ys.Head())
	}
	if ys.Length() != 1 {
		t.Errorf("[6,7,8].Filter >7.Length returns: %d", ys.Length())
	}
}
