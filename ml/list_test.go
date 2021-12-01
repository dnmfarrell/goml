package ml

import (
	"testing"
)

func TestFoldl(t *testing.T) {
	n := Nil[string]{}
	AccStrlen := func(acc int, x string) int { return acc + len(x) }
	x := Foldl[int, string](AccStrlen, 0, n)
	if x != 0 {
		t.Errorf("Foldl f 0 Nil returns %d", x)
	}
	words := Slice2List([]string{"foo", "bar", "baz"})
	y := Foldl[int, string](AccStrlen, 0, words)
	if y != 9 {
		t.Errorf("Foldl f 0 [foo,bar,baz] returns %d", y)
	}
	z := Foldl[int, string](AccStrlen, 0, words.Tail())
	if z != 6 {
		t.Errorf("Foldl f 0 [bar,baz] returns %d", z)
	}
	a := Foldl[int, string](AccStrlen, 0, Cons[string]{words.Head(), n})
	if a != 3 {
		t.Errorf("Foldl f 0 [foo] returns %d", a)
	}
}

func TestScan(t *testing.T) {
	n := Nil[string]{}
	AccStrlen := func(acc int, x string) int { return acc + len(x) }
	x := Scan[int, string](AccStrlen, 0, n)
	if x.Head() != 0 {
		t.Errorf("Scan f 0 Nil returns %v", x)
	}
	words := Slice2List([]string{"foo", "bar", "baz"})
	y := Scan[int, string](AccStrlen, 0, words)
	if y.Head() != 0 {
		t.Errorf("Scan f 0 [foo,bar,baz] Head returns %d", y.Head())
	}
	if y.Tail().Head() != 3 {
		t.Errorf("Scan f 0 [foo,bar,baz] Tail Head returns %d", y.Tail().Head())
	}
	if y.Last() != 9 {
		t.Errorf("Scan f 0 [foo,bar,baz] Last returns %d", y.Last())
	}
	z := Scan[int, string](AccStrlen, 0, words.Tail())
	if z.Head() != 0 {
		t.Errorf("Scan f 0 [bar,baz] Head returns %d", z.Head())
	}
	if z.Tail().Head() != 3 {
		t.Errorf("Scan f 0 [bar,baz] Tail Head returns %d", z.Tail().Head())
	}
	if z.Last() != 6 {
		t.Errorf("Scan f 0 [bar,baz] Last returns %d", z.Last())
	}
	a := Scan[int, string](AccStrlen, 0, Cons[string]{words.Head(), n})
	if a.Head() != 0 {
		t.Errorf("Scan f 0 [bar] Head returns %d", a.Head())
	}
	if a.Tail().Head() != 3 {
		t.Errorf("Scan f 0 [bar] Tail Head returns %d", a.Tail().Head())
	}
	if a.Last() != 3 {
		t.Errorf("Scan f 0 [bar] Last returns %d", a.Last())
	}
}

func TestSlice2List(t *testing.T) {
	l := Slice2List([]string{})
	if !l.Empty() {
		t.Error("empty slice does not return an empty list")
	}
	l = Slice2List([]string{"foo"})
	if l.Length() != 1 {
		t.Errorf("[foo] returns a %d length list", l.Length())
	}
	if l.Head() != "foo" {
		t.Errorf("[foo] slice not converted into [%s]", l.Head())
	}
	l = Slice2List([]string{"foo", "bar"})
	if l.Length() != 2 {
		t.Errorf("[foo,bar] converted into a %d length list", l.Length())
	}
	if l.Head() != "foo" {
		t.Errorf("[foo,bar] slice Head is %s", l.Head())
	}
	if l.Last() != "bar" {
		t.Errorf("[foo,bar] slice Last is %s", l.Last())
	}
}

func TestMap(t *testing.T) {
	ints := Cons[int]{3, Cons[int]{4, Nil[int]{}}}
	doubleInt := func(x int) int { return x * 2 }
	got := Map[int, int](doubleInt, ints)
	if got.Head() != 6 {
		t.Errorf("Map doubleInt Head returns: %d", got.Head())
	}
	if got.Last() != 8 {
		t.Errorf("Map doubleInt Last returns: %d", got.Last())
	}
}

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
	negInts := RangeInf[int]{-1, -1}.Eval()
	abs := func(x int) int { return x * -1 }
	posInts := MapL[int, int](abs, negInts)
	if posInts.Head() != 1 {
		t.Errorf("[-1,-2..].MapL abs.Head returns: %d", posInts.Last())
	}
	if posInts.Tail().Head() != 2 {
		t.Errorf("[-1,-2..].MapL abs.Tail.Head returns: %d", posInts.Tail().Head())
	}
}

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
