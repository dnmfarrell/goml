package ml

// Thunk is a wrapper for unevaluated expressions
type Thunk[A any] interface {
	Eval() List[A]
}

// ConsL is a lazy Cons. It has a head like Cons, but its tail is a Thunk,
// evaluated on demand. This means ConsL can represent infinite lists and
// streams.
type ConsL[A any] struct {
	head A
	tail Thunk[A]
}

// Head returns the ConsL head
func (c ConsL[A]) Head() A { return c.head }

// Tail evaluates the ConsL thunk, returning a new List
func (c ConsL[A]) Tail() List[A] {
	return c.tail.Eval()
}

// Empty always returns false as a ConsL is, by definition never empty
func (_ ConsL[A]) Empty() bool { return false }

// Lazy always returns true as by definition ConsL's tail has not been
// evaluated yet
func (_ ConsL[A]) Lazy() bool { return true }

// Last returns the last element of a List. Warning! Will not halt if called on
// an infinite list
func (c ConsL[A]) Last() A {
	if c.Tail().Empty() {
		return c.Head()
	}
	return c.Tail().Last()
}

// Init returns all but the last element of a List as a Cons. Warning! Will not
// halt if called on an infinite list
func (c ConsL[A]) Init() List[A] {
	if c.Tail().Empty() {
		return Nil[A]{}
	}
	return Cons[A]{c.Head(), c.Tail().Init()}
}

// Concat appends a List to this list. If the List parameter is empty, it
// returns the ConsL unchanged. Otherwise it returns a Cons. Warning! Will not
// halt if called on an infinite list
func (c ConsL[A]) Concat(l List[A]) List[A] {
	if l.Empty() {
		return c
	}
	return Cons[A]{c.Head(), c.Tail().Concat(l)}
}

// Reverse swaps the order of the elements of a ConsL, returning a Cons.
// Warning! Will not halt if called on an infinite list
func (c ConsL[A]) Reverse() List[A] {
	return c.Tail().Reverse().Concat(Cons[A]{c.Head(), Nil[A]{}})
}

// Drop discards the first n elements of a List
func (c ConsL[A]) Drop(n int) List[A] {
	if n > 1 {
		return c.Tail().Drop(n - 1)
	}
	return c.Tail()
}

// Take returns the first n elements of a ConsL as a Cons
func (c ConsL[A]) Take(n int) List[A] {
	if n > 1 {
		return Cons[A]{c.Head(), c.Tail().Take(n - 1)}
	}
	return Cons[A]{c.Head(), Nil[A]{}}
}

// Length returns the number of elements in the List. Warning! This will never
// return if the ConsL is an infinite list
func (c ConsL[A]) Length() int { return 1 + c.Tail().Length() }

// FilterThunk is a Thunk which uses a bool function to filter the output of
// another Thunk
type FilterThunk[A any] struct {
	f func(A) bool
	t Thunk[A]
}

// Eval evaluates the inner Thunk; if its bool function returns true for the
// result, or if the result is empty, it returns it. Otherwise it evaluates the
// next element. Warning! This will never return if the bool function returns
// false for every element of an infinite list
func (ft FilterThunk[A]) Eval() List[A] {
	list := ft.t.Eval()
	if list.Empty() {
		return list
	} else if ft.f(list.Head()) {
		return ConsL[A]{list.Head(), FilterThunk[A]{ft.f, list.(ConsL[A]).tail}}
	}
	return FilterThunk[A]{ft.f, list.(ConsL[A]).tail}.Eval()
}

// Filter applies a bool function to a ConsL, returning a new ConsL which
// lazily filters its elements. Multiple filters can be applied to any ConsL
func (c ConsL[A]) Filter(f func(A) bool) List[A] {
	if f(c.Head()) {
		return ConsL[A]{c.Head(), FilterThunk[A]{f, c.tail}}
	}
	return c.Tail().Filter(f)
}

// Num is a type for all primitive number types
type Num interface {
	int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 |
		~uint64 | ~uintptr | ~float32 | ~float64
}

// Range[Num] is a Thunk which returns an list of nums between its start and
// end values inclusive, incremented by its step value
type Range[A Num] struct {
	start A
	step  A
	end   A
}

// Eval increments the start Num by the step value, terminating if it's
// greater than the end Num, otherwise returning another ConsL with an
// incremented Range thunk
func (r Range[A]) Eval() List[A] {
	next := r.start + r.step
	if next > r.end {
		return Cons[A]{r.start, Nil[A]{}}
	}
	return ConsL[A]{r.start, Range[A]{next, r.step, r.end}}
}

// RangeInf[Num] is a Thunk which returns an infinite list of nums incremented
// by its step value. This can wrap-around in the case of overflow
type RangeInf[A Num] struct {
	start A
	step  A
}

// Eval increments the start Num by the step value, returning another ConsL
// with an incremented Range thunk
func (r RangeInf[A]) Eval() List[A] {
	return ConsL[A]{r.start, RangeInf[A]{r.start + r.step, r.step}}
}