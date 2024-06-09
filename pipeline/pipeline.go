package pipeline

func Generate[T any](source []T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for _, v := range source {
			out <- v
		}
	}()
	return out
}

func Map[T any, U any](in <-chan T, fn func(T) U) <-chan U {
	out := make(chan U)
	go func() {
		defer close(out)
		for v := range in {
			out <- fn(v)
		}
	}()
	return out
}

func Filter[T any](in <-chan T, fn func(T) bool) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for v := range in {
			if fn(v) {
				out <- v
			}
		}
	}()
	return out
}

func Collect[T any](done <-chan struct{}, in <-chan T) []T {
	result := make([]T, 0)
	for v := range in {
		select {
		case <-done:
			return result
		default:
			result = append(result, v)
		}
	}
	return result
}
