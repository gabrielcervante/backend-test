package validator

func IsEmpty[T comparable](value T) bool {
	var v T

	return value == v
}
