package optional

type Pair[L, R any] struct {
	Left  L
	Right R
}

func Zip[L, R any](left With[L], right With[R]) With[Pair[L, R]] {
	if left.IsEmpty() || right.IsEmpty() {
		return Empty[Pair[L, R]]()
	}

	return New(Pair[L, R]{
		Left:  *left.value,
		Right: *right.value,
	})
}

func Unzip[L, R any](zipped With[Pair[L, R]]) (With[L], With[R]) {
	if zipped.IsEmpty() {
		return None[L](), None[R]()
	}

	pair := zipped.value
	return New(pair.Left), New(pair.Right)
}
