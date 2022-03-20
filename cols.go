package cols

// Map will map any slice
func Map[T any, R any](in []T, fun func(T) R) []R {
	res := make([]R, len(in))
	for i, v := range in {
		res[i] = fun(v)
	}
	return res
}

// Filter will filter slice based on true/false returned by
// function
func Filter[T any](in []T, fun func(T) bool) []T {
	res := make([]T, 0, len(in))
	for _, v := range in {
		if fun(v) {
			res = append(res, v)
		}
	}
	return res
}

// Sum
//func Sum[T N](in []T) T {
//	total := 0
//	for _, v := range in {
//		total += v
//	}
//	return total
//}
//
//type N interface {
//	int | int32 | int64 | float32 | float64
//}
