package optional

type With[T any] struct {
	value *T
	empty bool
}

func (o With[T]) IsPresent() bool {
	return !o.empty
}

func (o With[T]) IsEmpty() bool {
	return o.empty
}

func (o With[T]) Unwrap() *T {
	return o.value
}

func (o With[T]) Get() T {
	return *o.value
}
