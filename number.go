package cols

// Number are interface that can be fulfilled by all int and float in golang
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// SumBy will return map that have sum of each group.
// Will be usefull if you want to create subtotal
func SumBy[T any, N Number, K comparable](
	in []T,
	keyFun func(T) K,
	numFun func(T) N,
) map[K]N {
	map1 := GroupBy(in, keyFun)

	map2 := make(map[K]N, len(map1))
	for k, l := range map1 {
		map2[k] = Sum(l, numFun)
	}

	return map2
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
