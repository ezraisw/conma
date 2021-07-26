package mapping

type MapperFunc func(x interface{}) interface{}

// Create a mapper that directly returns the specified value.
// Essentially, this mapper does not care about the matched element.
func Value(val interface{}) MapperFunc {
	return func(x interface{}) interface{} {
		return val
	}
}
