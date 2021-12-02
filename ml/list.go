// Package ml provides functional programming tools using generics, inspired
// by ml-tyle languages like Haskell.
package ml

// List defines the behavior all List types must implement
type List[A any] interface {
	Empty() bool
	Lazy() bool
	Head() A
	Last() A
	Tail() List[A]
}

// Length returns the number of elements in the List.
// Warning! Will not return if given an infinite list.
func Length[A any](l List[A]) uint {
	if l.Empty() {
		return 0
	}
	return 1 + Length(l.Tail())
}

// Drop discards the first n elements of a List.
func Drop[A any](n uint, l List[A]) List[A] {
	if l.Empty() {
		return Nil[A]{}
	}
	if n > 1 {
		return Drop(n-1, l.Tail())
	}
	return l.Tail()
}

// Take returns the first n elements of a List in a new List.
func Take[A any](n uint, l List[A]) List[A] {
	if l.Empty() {
		return Nil[A]{}
	}
	if n > 1 {
		return Cons[A]{l.Head(), Take(n-1, l.Tail())}
	}
	return Cons[A]{l.Head(), Nil[A]{}}
}

// Concat joins two Lists, returning a new List.
// Warning! Will not return if given an infinite list.
func Concat[A any](l List[A], l2 List[A]) List[A] {
	if l.Empty() {
		return l2
	}
	return Cons[A]{l.Head(), Concat(l.Tail(), l2)}
}

// Reverse inverts the order of the elements of a List, returning a new List.
// Warning! Will not return if given an infinite list.
func Reverse[A any](l List[A]) List[A] {
	if l.Empty() {
		return l
	}
	return Concat[A](Reverse(l.Tail()), Cons[A]{l.Head(), Nil[A]{}})
}

// Filter selects all elements of the List which satisfy a predicate, returning
// a new List.
// Warning! Will not return if given an infinite list.
func Filter[A any](f func(A) bool, l List[A]) List[A] {
	if l.Empty() {
		return l
	}
	if f(l.Head()) {
		return Cons[A]{l.Head(), Filter(f, l.Tail())}
	}
	return Filter(f, l.Tail())
}

// Map transforms a List by applying a function to every element, returning a
// new List.
func Map[A, B any](f func(A) B, l List[A]) List[B] {
	if l.Empty() {
		return Nil[B]{}
	}
	return Cons[B]{f(l.Head()), Map(f, l.Tail())}
}

// Foldl applies a function to every element in a List, accumulating the
// return values into a single value.
// Warning! Will not return if given an infinite list.
func Foldl[A, B any](f func(x A, y B) A, acc A, l List[B]) A {
	if l.Empty() {
		return acc
	}
	return Foldl(f, f(acc, l.Head()), l.Tail())
}

// Foldr applies a function to every element in a List, accumulating the
// return values into a single value.
// Warning! Will not return if given an infinite list.
func Foldr[A, B any](f func(x A, y B) B, acc B, l List[A]) B {
	if l.Empty() {
		return acc
	}
	return f(l.Head(), Foldr(f, acc, l.Tail()))
}

// Scan applies a function to every element in a List, returning a new List
// containing the output.
// Warning! Will not return if given an infinite list.
func Scan[A, B any](f func(x A, y B) A, acc A, l List[B]) List[A] {
	if l.Empty() {
		return Cons[A]{acc, Nil[A]{}}
	}
	return Cons[A]{acc, Scan(f, f(acc, l.Head()), l.Tail())}
}
