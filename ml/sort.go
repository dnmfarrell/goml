package ml

// Ordered are all the types which can be ordered via comparison operators
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 |
		~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

// Qsort takes a List of Ordered elements and returns a List of the elements in
// ascending order.
func Qsort[A Ordered](l List[A]) List[A] {
	if l.Empty() || l.Tail().Empty() {
		return l
	}
	h := l.Head()
	lt := Qsort(Filter(func(x A) bool { return x < h }, l.Tail()))
	ge := Qsort(Filter(func(x A) bool { return x >= h }, l.Tail()))
	return Concat[A](lt, Cons[A]{h, ge})
}

// QsortR takes a List of Ordered elements and returns a List of the elements in
// descending order. Warning! Will not halt if given an infinite list
func QsortR[A Ordered](l List[A]) List[A] {
	if l.Empty() || l.Tail().Empty() {
		return l
	}
	h := l.Head()
	lt := QsortR(Filter(func(x A) bool { return x < h }, l.Tail()))
	ge := QsortR(Filter(func(x A) bool { return x >= h }, l.Tail()))
	return Concat[A](ge, Cons[A]{h, lt})
}
