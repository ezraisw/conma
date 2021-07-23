package condition

import (
	"reflect"
	"strings"
)

type (
	checkCond struct {
		fn func(x interface{}) bool
	}
	fieldCheckCond struct {
		target []string
		fn     func(x interface{}) bool
	}

	CheckFunc func(x interface{}) bool
)

// Matches to true if an element satisfies the given check function for the specified value.
func Check(fn CheckFunc) Condition {
	return checkCond{
		fn: fn,
	}
}

func (c checkCond) Test(mctx MatchContext) bool {
	return c.fn(mctx.CurrentValue())
}

// Matches to true if a struct element's value satisfies the given check function for the specified value.
func FieldCheck(target string, fn CheckFunc) Condition {
	return fieldCheckCond{
		target: strings.Split(target, "."),
		fn:     fn,
	}
}

func (c fieldCheckCond) Test(mctx MatchContext) bool {
	rv := reflect.ValueOf(mctx.CurrentValue())

	for i := 0; i < len(c.target); i++ {
		switch rv.Kind() {
		case reflect.Ptr, reflect.Interface:
			rv = rv.Elem()
		}

		switch rv.Kind() {
		case reflect.Struct:
			rv = rv.FieldByName(c.target[i])
		case reflect.Map:
			rv = rv.MapIndex(reflect.ValueOf(c.target[i]))
		default:
			return false
		}

		if !rv.IsValid() {
			return false
		}
	}

	val := rv.Interface()
	return c.fn(val)
}

// Shallow equality check with the given value.
func Eq(val interface{}) CheckFunc {
	return func(x interface{}) bool {
		return val == x
	}
}

// Deep (recursive) equality check with the given value.
func DeepEq(val interface{}) CheckFunc {
	return func(x interface{}) bool {
		return reflect.DeepEqual(val, x)
	}
}

// Length equality check of array, channel, map, slice, or string.
func Len(len int) CheckFunc {
	return func(x interface{}) bool {
		rv := reflect.ValueOf(x)

		switch rv.Kind() {
		case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
			return rv.Len() == len
		default:
			return false
		}
	}
}
