package ml

// Nil represents an empty List
type Nil[A any] struct{}

// Empty always return true, as Nil is empty by definition
func (_ Nil[A]) Empty() bool { return true }

// Lazy always returns false as by definition Nil has already been evaluated
func (_ Nil[A]) Lazy() bool { return false }

// Head panics - an empty List has no head
func (_ Nil[A]) Head() A { panic("cannot Head() an empty List") }

// Tail panics - an empty List has no head
func (_ Nil[A]) Tail() List[A] { panic("cannot Tail() an empty List") }

// Length returns a count of the elements in the empty List: 0
func (_ Nil[A]) Length() int { return 0 }

// Drop panics as you cannot drop elements from an empty List
func (_ Nil[A]) Drop(_ int) List[A] { panic("cannot Drop() an empty List") }

// Take panics as you cannot take elements from an empty List
func (_ Nil[A]) Take(_ int) List[A] { panic("cannot Take() an empty List") }

// Init panics - as an empty List has no initial elements
func (_ Nil[A]) Init() List[A] { panic("cannot Init() an empty List") }

// Last panics as an empty List has no final element
func (_ Nil[A]) Last() A { panic("cannot Last() an empty List") }

// Reverse returns itself
func (n Nil[A]) Reverse() List[A] { return n }

// Concat returns the List parameter it receives unchanged
func (_ Nil[A]) Concat(l List[A]) List[A] { return l }

// Filter returns itself
func (n Nil[A]) Filter(func(_ A) bool) List[A] { return n }