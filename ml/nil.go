package ml

// Nil represents an empty List
type Nil[A any] struct{}

// Empty always return true, as Nil is empty by definition
func (Nil[A]) Empty() bool { return true }

// Lazy always returns false as by definition Nil has already been evaluated
func (Nil[A]) Lazy() bool { return false }

// Head panics - an empty List has no head
func (Nil[A]) Head() A { panic("cannot Head() an empty List") }

// Last panics - an empty List has no last element
func (Nil[A]) Last() A { panic("cannot Last() an empty List") }

// Tail panics - an empty List has no tail
func (Nil[A]) Tail() List[A] { panic("cannot Tail() an empty List") }
