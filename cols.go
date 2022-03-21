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

// GroupBy will group slice []V into map[K][]V. The key (K) will be reciveved
// from the function that pass into second arguments
func GroupBy[K comparable, V any](in []V, fun func(V) K) map[K][]V {
	res := make(map[K][]V, len(in))
	for _, v := range in {
		k := fun(v)
		group := res[k]
		res[k] = append(group, v)
	}
	return res
}
