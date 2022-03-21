// Package cols contains function to aggregate list and use it as standard
// function. For most known function like Map and Filter should work as is
// should
package cols

// Number are interface that can be fulfilled by all int and float in golang
type Number interface {
	int | int32 | int64 | float32 | float64
}

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
		group := res[k]
		res[k] = append(group, v)
	}
	return res
}

// Sum will return sum of every number returned by function
func Sum[T any, N Number](in []T, fun func(T) N) N {
	return SumNumber(Map(in, fun))
}

// Max will return higest of every number returned by function
func Max[T any, N Number](in []T, fun func(T) N) N {
	return MaxNumber(Map(in, fun))
}

// Min will return lowest of every number returned by function
func Min[T any, N Number](in []T, fun func(T) N) N {
	return MinNumber(Map(in, fun))
}

// SumNumber will return sum of every number in input
func SumNumber[T Number](in []T) T {
	var res T
	for _, v := range in {
		res += v
	}
	return res
}

// MaxNumber will return the higest number in input
func MaxNumber[T Number](in []T) T {
	var res T
	if len(in) == 0 {
		return res
	}

	res = in[0]

	for _, v := range in {
		if v > res {
			res = v
		}
	}

	return res
}

// MinNumber will return the higest number in input
func MinNumber[T Number](in []T) T {
	var res T
	if len(in) == 0 {
		return res
	}

	res = in[0]

	for _, v := range in {
		if v < res {
			res = v
		}
	}

	return res
}
