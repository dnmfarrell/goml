package ml

import (
	"testing"
)

func TestCompose(t *testing.T) {
	sqr := func(x int) int { return x * x }
	inc := func(x int) int { return x + 1 }
	sqrinc := Compose(sqr, inc)
	x := sqrinc(5)
	if x != 26 {
		t.Errorf("5 * %% + 1 returns: %d", x)
	}
}

func TestSlice2List(t *testing.T) {
	l := Slice2List([]string{})
	if !l.Empty() {
		t.Error("empty slice does not return an empty list")
	}
	l = Slice2List([]string{"foo"})
	if Length(l) != 1 {
		t.Errorf("[foo] returns a %d length list", Length(l))
	}
	if l.Head() != "foo" {
		t.Errorf("[foo] slice not converted into [%s]", l.Head())
	}
	l = Slice2List([]string{"foo", "bar"})
	if Length(l) != 2 {
		t.Errorf("[foo,bar] converted into a %d length list", Length(l))
	}
	if l.Head() != "foo" {
		t.Errorf("[foo,bar].Head is %s", l.Head())
	}
	if l.Tail().Head() != "bar" {
		t.Errorf("[foo,bar].Tail.Head is %s", l.Tail().Head())
	}
}
