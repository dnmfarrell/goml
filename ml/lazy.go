package ml

// ListThunk wraps any List as a Thunk. This is used by MapL to convert a Cons
// into a ConsL so it can be lazily mapped
type ListThunk[A any] struct {
	l List[A]
}

// Eval returns the next element in the List
func (lt ListThunk[A]) Eval() List[A] {
	if lt.l.Empty() {
		return Nil[A]{}
	}
	return ConsL[A]{lt.l.Head(), ListThunk[A]{lt.l.Tail()}}
}

// MapThunk uses a unary function to transform the output of another Thunk
type MapThunk[A, B any] struct {
	f func(A) B
	t Thunk[A]
}

// Eval evaluates the inner Thunk transforming its return value from A to B
func (mt MapThunk[A, B]) Eval() List[B] {
	l := mt.t.Eval()
	if l.Empty() {
		return Nil[B]{}
	} else if l.Lazy() {
		return ConsL[B]{mt.f(l.Head()), MapThunk[A, B]{mt.f, l.(ConsL[A]).tail}}
	}
	return ConsL[B]{mt.f(l.Head()), MapThunk[A, B]{mt.f, ListThunk[A]{l.Tail()}}}
}

// MapL returns a ConsL which lazily applies a unary function to the input
// List. Unlike Map it is safe to use with infinite lists.
func MapL[A, B any](f func(A) B, l List[A]) List[B] {
	if l.Empty() {
		return Nil[B]{}
	}
	if l.Lazy() {
		return ConsL[B]{f(l.Head()), MapThunk[A, B]{f, l.(ConsL[A]).tail}}
	}
	return ConsL[B]{f(l.Head()), MapThunk[A, B]{f, ListThunk[A]{l.Tail()}}}
}

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
	l := ft.t.Eval()
	if l.Empty() {
		return l
	}
	var th Thunk[A]
	if l.Lazy() {
		th = l.(ConsL[A]).tail
	} else {
		th = ListThunk[A]{l.Tail()}
	}
	if ft.f(l.Head()) {
		return ConsL[A]{l.Head(), FilterThunk[A]{ft.f, th}}
	}
	return FilterThunk[A]{ft.f, th}.Eval()
}

// FilterL lazily applies a predicate to a List. Multiple filters can be
// applied
func FilterL[A any](f func(A) bool, l List[A]) List[A] {
	if l.Empty() {
		return l
	}
	if f(l.Head()) {
		var th Thunk[A]
		if l.Lazy() {
			th = l.(ConsL[A]).tail
		} else {
			th = ListThunk[A]{l.Tail()}
		}
		return ConsL[A]{l.Head(), FilterThunk[A]{f, th}}
	}
	return FilterL(f, l.Tail())
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

// RangeDec[Num] is a Thunk which returns an list of nums between its start and
// end values inclusive, decremented by its step value
type RangeDec[A Num] struct {
	start A
	step  A
	end   A
}

// Eval decrements the start Num by the step value, terminating if it's
// greater than the end Num, otherwise returning another ConsL with a
// decremented RangeDec thunk
func (r RangeDec[A]) Eval() List[A] {
	next := r.start - r.step
	if next < r.end {
		return Cons[A]{r.start, Nil[A]{}}
	}
	return ConsL[A]{r.start, RangeDec[A]{next, r.step, r.end}}
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
