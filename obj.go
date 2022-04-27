package cols

// Collection is struct that represent array or list
// that can use aggregat function like Map or Filter
// In inside it contains slice that can be filtered
// or mapped
type Collection[T any] struct {
	content []T
}

// From is function to create Collection from slice.
func From[T any](slice []T) *Collection[T] {
	return &Collection[T]{
		content: slice,
	}
}

func (c *Collection[T]) Filter(f func(T) bool) *Collection[T] {
	newSlice := Filter(c.content, f)
	return From[T](newSlice)
}

func (c *Collection[T]) Slice() []T {
	return c.content
}
