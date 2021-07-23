package mapping

type Mapper interface {
	Map(x interface{}) interface{}
}

type MapperFunc func(x interface{}) interface{}

func (m MapperFunc) Map(x interface{}) interface{} {
	return m(x)
}

func Value(val interface{}) MapperFunc {
	return func(x interface{}) interface{} {
		return val
	}
}
