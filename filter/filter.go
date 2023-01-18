package filter

func Filter[E any](slice []E, boolF func(e E) bool) []E {
	result := make([]E, 0, len(slice))
	for i := range slice {
		if boolF(slice[i]) {
			result = append(result, slice[i])
		}
	}
	return result
}
