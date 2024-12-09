package optional

func (o With[T]) Or(fallback With[T]) With[T] {
	if o.IsEmpty() {
		return fallback
	}

	return o
}

func (o With[T]) OrElseGet(def T) With[T] {
	if o.IsPresent() {
		return o
	}

	return Some(def)
}

func (o With[T]) OrElse(supplier func() With[T]) With[T] {
	if o.IsPresent() {
		return o
	}

	return supplier()
}

func (o With[T]) GetOr(v T) T {
	if o.IsPresent() {
		return *o.value
	}

	return v
}

func (o With[T]) TakeOr(v T) T {
	return o.GetOr(v)
}

func (o With[T]) OrElsePanic() T {
	if o.IsEmpty() {
		panic("Optional value is not present")
	}

	return *o.value
}
