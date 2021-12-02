package ml

import (
	"testing"
)

func TestMapL(t *testing.T) {
	words := Slice2List([]string{"foo", "bar", "bazz"})
	strlen := func(x string) int { return len(x) }
	lens := MapL[string, int](strlen, words)
	if lens.Head() != 3 {
		t.Errorf("[foo,bar,baz].MapL strlen.Head returns: %d", lens.Head())
	}
	if lens.Last() != 4 {
		t.Errorf("[foo,bar,baz].MapL strlen.Last returns: %d", lens.Last())
	}
	ns := MapL[string, int](strlen, Nil[string]{})
	if !ns.Empty() {
		t.Error("MapL(strlen, Nil) is not empty")
	}
}

func TestRange(t *testing.T) {
	tostr := func(x rune) string { return string(x) }
	xs := Range[rune]{'A', 1, 'Z'}.Eval()
	ys := MapL(tostr, xs)
	if ys.Head() != "A" {
		t.Errorf("Range[A,1,Z].Head returns: %s", ys.Head())
	}
	if ys.Tail().Head() != "B" {
		t.Errorf("Range[A,1,Z].Tail.Head returns: %s", ys.Tail().Head())
	}
	if ys.Last() != "Z" {
		t.Errorf("Range[A,1,Z].Last returns: %s", ys.Last())
	}
}

func TestRangeDec(t *testing.T) {
	xs := RangeDec[int]{10, 1, -10}.Eval()
	if xs.Head() != 10 {
		t.Errorf("RangeDec{10,1,-10}.Head returns: %d", xs.Head())
	}
	if xs.Last() != -10 {
		t.Errorf("RangeDec{10,1,-10}.Last returns: %d", xs.Last())
	}
	if Length(xs) != 21 {
		t.Errorf("RangeDec{10,1,-10}.Length returns: %d", Length(xs))
	}
}

func TestFilterL(t *testing.T) {
	xs := RangeInf[int]{0, 1}.Eval()
	even := func(x int) bool { return x%2 == 0 }
	ys := Take(3, FilterL(even, xs))
	if ys.Head() != 0 {
		t.Errorf("ys.Head is not 0: %d", ys.Head())
	}
	if ys.Tail().Head() != 2 {
		t.Errorf("ys.Tail.Head is not 2 even: %d", ys.Tail().Head())
	}
	if ys.Last() != 4 {
		t.Errorf("ys.Last is not 4: %d", ys.Last())
	}
	yes := func(x int) bool { return true }
	ns := FilterL[int](yes, Nil[int]{})
	if !ns.Empty() {
		t.Error("FilterL(yes, Nil) is not empty")
	}
	zs := FilterL[int](yes, Cons[int]{0, Cons[int]{1, Nil[int]{}}})
	if zs.Head() != 0 {
		t.Errorf("FilterL(yes, [0..1]).Head returns: %d", zs.Head())
	}
	as := FilterL[int](yes, ConsL[int]{1, ListThunk[int]{Nil[int]{}}})
	if !as.Tail().Empty() {
		t.Error("FilterL(yes, Nil) is not empty")
	}
}
