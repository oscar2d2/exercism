package flatten

func Flatten(v interface{}) []interface{} {
	out := []interface{}{}
	switch v.(type) {
	case []interface{}:
		for _, elem := range v.([]interface{}) {
			out = append(out, Flatten(elem)...)
		}
	default:
		if v != nil {
			out = append(out, v)
		}
	}

	return out
}
