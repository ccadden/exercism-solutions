package strain

func Keep[T any](collection []T, p func(T) bool) []T {
	res := []T{}
	for _, item := range(collection) {
		if p(item) {
			res = append(res, item)
		}
	}

	return res
}

func Discard[T any](collection []T, p func(T) bool) []T {
	res := []T{}
	for _, item := range(collection) {
		if !p(item) {
			res = append(res, item)
		}
	}

	return res
}
