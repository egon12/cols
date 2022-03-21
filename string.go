package cols

import "strings"

// MapString will map any slice to string slice.
// The length of the input and the output should be same.
func MapString[T any](in []T, fun func(T) string) []string {
	return Map(in, fun)
}

// JoinString will return slice string joined into one string
// and seperated by comman (, ) comma space
func JoinString[T any](in []T, fun func(T) string) string {
	return JoinStringWithSep(in, fun, ",")
}

// JoinStringWithSep will return slice string joined into one string
// and seperated sep argument
func JoinStringWithSep[T any](in []T, fun func(T) string, sep string) string {
	return strings.Join(MapString(in, fun), sep)
}
