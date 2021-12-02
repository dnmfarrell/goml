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
func (ConsL[A]) Empty() bool { return false }

// Lazy always returns true as by definition ConsL's tail has not been
// evaluated yet
func (ConsL[A]) Lazy() bool { return true }

// Last returns the last element of the List
func (c ConsL[A]) Last() A {
	if c.Tail().Empty() {
		return c.Head()
	}
	return c.Tail().Last()
}
