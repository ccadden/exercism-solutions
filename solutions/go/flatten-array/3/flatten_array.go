package flatten

func Flatten(nested any) []any {
	flattened := []any{}

	values, ok := nested.([]any)

	if !ok {
		return flattened
	}

	for _, v := range(values) {
		switch v.(type) {
		case nil:
			continue
		case []any:
				flattened = append(flattened, Flatten(v)...)
		default:
				flattened = append(flattened, v)

		}
	}

	return flattened
}
