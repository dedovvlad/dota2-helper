package common

func TransformSlice[T, K any](slice []T, transform func(item T) K) []K {
	result := make([]K, len(slice))
	for i, item := range slice {
		result[i] = transform(item)
	}

	return result
}
