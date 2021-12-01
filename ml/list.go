// Package ml provides functional programming tools using generics, inspired
// by ml-tyle languages like Haskell.
package ml

// List defines the behavior all List types must implement
type List[A any] interface {
	Empty() bool
	Lazy() bool
	Head() A
	Tail() List[A]
	Length() int
	Drop(int) List[A]
	Take(int) List[A]
	Last() A
	Init() List[A]
	Concat(List[A]) List[A]
	Reverse() List[A]
	Filter(func(A) bool) List[A]
}

// Compose joins two functions together returning a new function. The return
// type of the first function must match the parameter type of the second one.
func Compose[A, B, C any](f1 func(A) B, f2 func(B) C) func(A) C {
	return func(x A) C { return f2(f1(x)) }
}

// Map transforms a List by applying a function to every element, returning a
// new List. Warning! Will not halt if called on an infinite list
func Map[A, B any](f func(A) B, xs List[A]) List[B] {
	if xs.Empty() {
		return Nil[B]{}
	}
	return Cons[B]{f(xs.Head()), Map(f, xs.Tail())}
}

// MapThunk is a Thunk which uses a unary function to transform the output of
// another Thunk
type MapThunk[A, B any] struct {
	f func(A) B
	t Thunk[A]
}

// Eval evaluates the inner Thunk transforming its return value from A to B
func (mt MapThunk[A, B]) Eval() List[B] {
	l := mt.t.Eval()
	if l.Empty() {
		return Nil[B]{}
	}
	return ConsL[B]{mt.f(l.Head()), MapThunk[A, B]{mt.f, l.(ConsL[A]).tail}}
}

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

// MapL returns a ConsL which lazily applies a unary function to the input
// List. Unlike Map it is safe to use with infinite lists.
func MapL[A, B any](f func(A) B, xs List[A]) List[B] {
	if xs.Empty() {
		return Nil[B]{}
	}
	if !xs.Lazy() {
		xs = ListThunk[A]{xs}.Eval()
	}
	return ConsL[B]{f(xs.Head()), MapThunk[A, B]{f, xs.(ConsL[A]).tail}}
}

// Ordered are all the types which can be ordered via comparison operators
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 |
		~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

// Qsort takes a List of Ordered elements and returns a List of the elements in
// ascending order. Warning! Will not halt if given an infinite list
func Qsort[A Ordered](xs List[A]) List[A] {
	if xs.Empty() || xs.Tail().Empty() {
		return xs
	}
	h := xs.Head()
	lt := Qsort(xs.Tail().Filter(func(x A) bool { return x < h }))
	ge := Qsort(xs.Tail().Filter(func(x A) bool { return x >= h }))
	return lt.Concat(Cons[A]{h, ge})
}

// QsortR takes a List of Ordered elements and returns a List of the elements in
// descending order. Warning! Will not halt if given an infinite list
func QsortR[A Ordered](xs List[A]) List[A] {
	if xs.Empty() || xs.Tail().Empty() {
		return xs
	}
	h := xs.Head()
	lt := QsortR(xs.Tail().Filter(func(x A) bool { return x < h }))
	ge := QsortR(xs.Tail().Filter(func(x A) bool { return x >= h }))
	return ge.Concat(Cons[A]{h, lt})
}

// Foldl applies a function to every element in a List, accumulating the
// return values into a single value. Warning! Will not halt if given an
// infinite list
func Foldl[A, B any](f func(x A, y B) A, acc A, xs List[B]) A {
	if xs.Empty() {
		return acc
	}
	return Foldl(f, f(acc, xs.Head()), xs.Tail())
}

// Scan applies a function to every element in a List, returning a new List
// containing the output. Warning! Will not halt if given an infinite list
func Scan[A, B any](f func(x A, y B) A, acc A, xs List[B]) List[A] {
	if xs.Empty() {
		return Cons[A]{acc, Nil[A]{}}
	}
	return Cons[A]{acc, Scan(f, f(acc, xs.Head()), xs.Tail())}
}

// Slice2List converts any slice into a List of the same type
func Slice2List[A any](s []A) List[A] {
	if len(s) == 0 {
		return Nil[A]{}
	} else if len(s) == 1 {
		return Cons[A]{s[0], Nil[A]{}}
	}
	return Cons[A]{s[0], Nil[A]{}}.Concat(Slice2List(s[1:]))
}