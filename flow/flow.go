package flow

// Define a Flow-like struct in Go
type Flow[T any] struct {
	ch chan T
}

// Function to emit values to the channel
func (f *Flow[T]) Emit(val T) {
	f.ch <- val
}

// Function to close the channel
func (f *Flow[T]) Close() {
	close(f.ch)
}

// Function to collect values from the channel
func (f *Flow[T]) Collect(action func(T)) {
	for val := range f.ch {
		action(val)
	}
}

// Function to create a new Flow
func NewFlow[T any]() *Flow[T] {
	return &Flow[T]{ch: make(chan T)}
}
