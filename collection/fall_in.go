package collection

func FallIn(mains []string, left, right string) []string {
	return Collection(mains, func(s string) string { return left + s + right })
}
