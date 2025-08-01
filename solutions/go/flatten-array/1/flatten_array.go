package flatten

import "reflect"

func Flatten(nested any) []interface{} {
	flattened := []any{}

	nestedSlice, ok := nested.([]interface{})

	if !ok {
		return flattened
	}

	for _, v := range(nestedSlice) {
		t := reflect.TypeOf(v) 

		if t != nil {
			k := t.Kind()
			switch k {
			case reflect.Slice:
				flattened = append(flattened, Flatten(v)...)
			default:
				flattened = append(flattened, v)
			}
		}
	}

	return flattened
}
