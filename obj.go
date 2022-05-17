package cols

// Collection is struct that represent array or list
// that can use aggregat function like Map or Filter
// In inside it contains slice that can be filtered
//
// For now it's still in experiment mode, because
// It cannot contain Map
// This caused by design of the language
type Collection[T any] struct {
	content []T
}

// From is function to create Collection from slice.
func From[T any](slice []T) *Collection[T] {
	return &Collection[T]{
		content: slice,
	}
}

// Filter will return Collection that filtered by function
func (c *Collection[T]) Filter(f func(T) bool) *Collection[T] {
	newSlice := Filter(c.content, f)
	return From[T](newSlice)
}

// Slice will return native slice type in golang
func (c *Collection[T]) Slice() []T {
	return c.content
}

// Map
// This implementation ins't work yet because it still rais
// compile error 'method must have no type parameters'
// The logic behind this are in
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#no-parameterized-methods
/*
func (c *Collection[T]) Map[R any](f func(T) R) *Collection[R] {
	newSlice := Map(c.content, f)
	return From[R](newSlice)
}
*/
