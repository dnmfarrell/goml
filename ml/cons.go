package ml

// Cons represents a cell in a linked list, it has a head and a tail which is
// another List
type Cons[A any] struct {
	head A
	tail List[A]
}

// Empty always returns false as by definition a Cons cell is not empty
func (Cons[A]) Empty() bool { return false }

// Lazy always returns false as by definition a Cons cell has already been
// evaluated.
func (Cons[A]) Lazy() bool { return false }

// Head returns the Cons head
func (c Cons[A]) Head() A { return c.head }

// Tail returns the Cons tail
func (c Cons[A]) Tail() List[A] { return c.tail }

// Last returns the last element of the List
func (c Cons[A]) Last() A {
	if c.Tail().Empty() {
		return c.Head()
	}
	return c.Tail().Last()
}
