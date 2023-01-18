package collection

func Collection[E, F any](slice []E, collection func(e E) F) []F {
	result := make([]F, len(slice))
	for i := 0; i < len(slice); i++ {
		result[i] = collection(slice[i])
	}
	return result
}
