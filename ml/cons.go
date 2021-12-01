package ml

// Cons represents a cell in a linked list, it has a head and a tail which is
// another List
type Cons[A any] struct {
	head A
	tail List[A]
}

// Empty always returns false as by definition a Cons cell is not empty
func (_ Cons[A]) Empty() bool { return false }

// Lazy always returns false as by definition a Cons cell has already been
// evaluated.
func (_ Cons[A]) Lazy() bool { return false }

// Head returns the Cons head
func (c Cons[A]) Head() A { return c.head }

// Tail returns the Cons tail
func (c Cons[A]) Tail() List[A] { return c.tail }

// Length returns the number of elements in the List
func (c Cons[A]) Length() int { return 1 + c.Tail().Length() }

// Drop discards the first n elements of a List
func (c Cons[A]) Drop(n int) List[A] {
	if n > 1 {
		return c.Tail().Drop(n - 1)
	}
	return c.Tail()
}

// Take returns the first n elements of a List in a new List
func (c Cons[A]) Take(n int) List[A] {
	if n > 1 {
		return Cons[A]{c.Head(), c.Tail().Take(n - 1)}
	}
	return Cons[A]{c.Head(), Nil[A]{}}
}

// Last returns the last element of a List
func (c Cons[A]) Last() A {
	if c.Tail().Empty() {
		return c.Head()
	}
	return c.Tail().Last()
}

// Init returns all but the last element of a List
func (c Cons[A]) Init() List[A] {
	if c.Tail().Empty() {
		return Nil[A]{}
	}
	return Cons[A]{c.Head(), c.Tail().Init()}
}

// Concat joins two Lists
func (c Cons[A]) Concat(l List[A]) List[A] {
	if l.Empty() {
		return c
	}
	return Cons[A]{c.Head(), c.Tail().Concat(l)}
}

// Reverse reverses the order of the elements of a List, returning a new List
func (c Cons[A]) Reverse() List[A] {
	return c.Tail().Reverse().Concat(Cons[A]{c.Head(), Nil[A]{}})
}

// Filter applies a bool function to a List, returning a new List comprised
// of elements that the function returned true for
func (c Cons[A]) Filter(f func(A) bool) List[A] {
	if f(c.Head()) {
		return Cons[A]{c.Head(), c.Tail().Filter(f)}
	}
	return c.Tail().Filter(f)
}
