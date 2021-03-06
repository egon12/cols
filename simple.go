// Package cols contains function to aggregate list and use it as standard
// function. For most known function like Map and Filter should work as is
// should
package cols

// Map will map any slice to other slice. The length of the input and the
// output should be same.
func Map[T any, R any](in []T, fun func(T) R) []R {
	res := make([]R, len(in))
	for i, v := range in {
		res[i] = fun(v)
	}
	return res
}

// Filter will filter slice based on true/false returned by function
func Filter[T any](in []T, fun func(T) bool) []T {
	res := make([]T, 0, len(in))
	for _, v := range in {
		if fun(v) {
			res = append(res, v)
		}
	}
	return res
}

// Reduce is function reduce just like javascript reduce
func Reduce[T, R any](in []T, fun func(R, T) R, val R) R {
	for _, v := range in {
		val = fun(val, v)
	}
	return val
}

// Any will return true if in slice there are item that match with the function
func Any[T any](in []T, fun func(T) bool) bool {
	_, ok := Find(in, fun)
	return ok
}

// Find will return item based on true/false returned by function
func Find[T any](in []T, fun func(T) bool) (it T, ok bool) {
	for _, v := range in {
		if fun(v) {
			return v, true
		}
	}
	return it, false
}

// GroupBy will group slice []V into map[K][]V. The key (K) will be reciveved
// from the function that pass into second arguments
func GroupBy[K comparable, V any](in []V, fun func(V) K) map[K][]V {
	res := make(map[K][]V, len(in))
	for _, v := range in {
		k := fun(v)
		res[k] = append(res[k], v)
	}
	return res
}

// CountBy will count slice []V based on key K map[K][]V. The key (K) will be received
// from the function that pass into second arguments
func CountBy[K comparable, V any](in []V, fun func(V) K) map[K]int {
	res := make(map[K]int, len(in))
	for _, v := range in {
		res[fun(v)]++
	}
	return res
}

// Unique will return slice of unique values
func Unique[T comparable](in []T) []T {
	unique := make(map[T]struct{}, len(in))

	for _, v := range in {
		_, ok := unique[v]
		if !ok {
			// mark unique
			unique[v] = struct{}{}
		}
	}

	res := make([]T, len(unique))
	i := 0
	for v := range unique {
		res[i] = v
		i++
	}

	return res
}
