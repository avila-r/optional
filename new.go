package optional

func Of[T any](v *T) With[T] {
	if v != nil {
		return With[T]{
			value: v,
			empty: false,
		}
	}
	
	return Empty[T]()
}

func Empty[T any]() With[T] {
	return With[T]{
		value: nil,
		empty: true,
	}
}
