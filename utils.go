package optional

func (o With[T]) IfPresent(do func()) {
	if o.IsPresent() {
		do()
	}
}

func (o With[T]) IfEmpty(do func()) {
	if o.IsEmpty() {
		do()
	}
}

func (o With[T]) Filter(predicate func(value T) bool) With[T] {
	if o.IsPresent() && predicate(*o.value) {
		return o
	}

	return Empty[T]()
}

func (o With[T]) Peek(action func(value T)) With[T] {
	if o.IsPresent() {
		action(*o.value)
	}

	return o
}

func (o With[T]) ToPointer() *T {
	return o.value
}

func (o With[T]) Match(onPresent func(value T), onEmpty func()) {
	if o.IsPresent() {
		onPresent(*o.value)
	} else {
		onEmpty()
	}
}

func (o With[T]) Recover(supplier func() T) With[T] {
	if o.IsEmpty() {
		v := supplier()

		return Of(&v)
	}

	return o
}

