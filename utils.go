package optional

func (o With[T]) IfPresent(do func(value T)) {
	if o.IsPresent() {
		do(*o.value)
	}
}

func (o With[T]) OrElse(def T) T {
	if o.IsPresent() {
		return *o.value
	}

	return def
}

func (o With[T]) OrElseGet(supplier func() T) T {
	if o.IsPresent() {
		return *o.value
	}

	return supplier()
}

func (o With[T]) OrElsePanic() T {
	if o.IsEmpty() {
		panic("Optional value is not present")
	}

	return *o.value
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

