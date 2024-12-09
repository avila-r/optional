package optional

func New[T any](v T) With[T] {
	return With[T]{
		value: &v,
		empty: false,
	}
}

func Of[T any](v *T) With[T] {
	if v != nil {
		return With[T]{
			value: v,
			empty: false,
		}
	}
	
	return Empty[T]()
}

func None[T any]() With[T] {
	return With[T]{
		value: nil,
		empty: true,
	}
}

func Empty[T any]() With[T] {
	return None[T]()
}
