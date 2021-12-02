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

func TestFoldr(t *testing.T) {
	n := Nil[string]{}
	strlen := func(x string, acc int) int { return acc + len(x) }
	x := Foldr[string, int](strlen, 0, n)
	if x != 0 {
		t.Errorf("Foldr f 0 Nil returns %d", x)
	}
	words := Slice2List([]string{"foo", "bar", "baz"})
	y := Foldr[string, int](strlen, 0, words)
	if y != 9 {
		t.Errorf("Foldr f 0 [foo,bar,baz] returns %d", y)
	}
	z := Foldr[string, int](strlen, 0, words.Tail())
	if z != 6 {
		t.Errorf("Foldr f 0 [bar,baz] returns %d", z)
	}
	a := Foldr[string, int](strlen, 0, Cons[string]{words.Head(), n})
	if a != 3 {
		t.Errorf("Foldr f 0 [foo] returns %d", a)
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

func TestDrop(t *testing.T) {
	xs := RangeInf[int]{1, 1}.Eval()
	ys := Drop(5, xs)
	if ys.Head() != 6 {
		t.Errorf("Drop(5, [1..]).Head returns: %d", ys.Head())
	}
	zs := Drop[int](1, Nil[int]{})
	if !zs.Empty() {
		t.Error("Drop(1, Nil) does not return Nil")
	}
}

func TestTake(t *testing.T) {
	xs := RangeInf[int]{1, 1}.Eval()
	ys := Take(5, xs)
	if ys.Head() != 1 {
		t.Errorf("Take(5, [1..]).Head returns: %d", ys.Head())
	}
	if ys.Last() != 5 {
		t.Errorf("Take(5, [1..]).Last returns: %d", ys.Last())
	}
	zs := Take[int](1, Nil[int]{})
	if !zs.Empty() {
		t.Error("Take(1, Nil) does not return Nil")
	}
}

func TestReverse(t *testing.T) {
	xs := Range[int]{1, 1, 10}.Eval()
	ys := Reverse(xs)
	if ys.Head() != 10 {
		t.Errorf("Reverse([1..]).Head returns: %d", ys.Head())
	}
	if ys.Last() != 1 {
		t.Errorf("Reverse(5, [1..]).Last returns: %d", ys.Last())
	}
	zs := Reverse[int](Nil[int]{})
	if !zs.Empty() {
		t.Error("Reverse(Nil) does not return Nil")
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

func TestFilter(t *testing.T) {
	xs := Cons[int]{6, Cons[int]{7, Cons[int]{8, Nil[int]{}}}}
	bigInt := func(x int) bool { return x > 7 }
	ys := Filter[int](bigInt, xs)
	if ys.Head() != 8 {
		t.Errorf("[6,7,8].Filter >7.Head returns: %d", ys.Head())
	}
	if Length(ys) != 1 {
		t.Errorf("[6,7,8].Filter >7.Length returns: %d", Length(ys))
	}
}
