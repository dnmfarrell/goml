package ml

// Compose joins two functions together returning a new function. The return
// type of the first function must match the parameter type of the second one.
func Compose[A, B, C any](f1 func(A) B, f2 func(B) C) func(A) C {
	return func(x A) C { return f2(f1(x)) }
}

// Slice2List converts any slice into a List of the same type
func Slice2List[A any](s []A) List[A] {
	if len(s) == 0 {
		return Nil[A]{}
	} else if len(s) == 1 {
		return Cons[A]{s[0], Nil[A]{}}
	}
	return Concat[A](Cons[A]{s[0], Nil[A]{}}, Slice2List(s[1:]))
}
