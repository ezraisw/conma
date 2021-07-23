package condition_test

import (
	"testing"

	"github.com/pwnedgod/conma/condition"
)

func TestCheckDeepEq(t *testing.T) {
	condVal := dummyStructValues[0]
	c := condition.Check(condition.DeepEq(condVal))

	values := []interface{}{
		dummyStructValues[0],
		dummyStructValues[1],
	}

	test := CondTest{
		Values:       values,
		Expectations: makeExpectations(len(values), []int{0}),
	}

	testCond(t, c, test)
}

func TestCheckEq(t *testing.T) {
	condVal := dummyStringValues[0]
	c := condition.Check(condition.Eq(condVal))

	values := []interface{}{
		dummyStringValues[0],
		dummyStringValues[1],
	}

	test := CondTest{
		Values:       values,
		Expectations: makeExpectations(len(values), []int{0}),
	}

	testCond(t, c, test)
}

func TestCheckLen(t *testing.T) {
	condLen := 3
	c := condition.Check(condition.Len(condLen))

	values := []interface{}{
		dummyStructValues,
		dummyStringValues,
		dummyIntValues,
		dummyRogueStructValue,
	}

	test := CondTest{
		Values:       values,
		Expectations: makeExpectations(len(values), []int{0}),
	}

	testCond(t, c, test)
}

func TestCheck(t *testing.T) {
	condFn := func(x interface{}) bool {
		str, ok := x.(string)
		if !ok {
			return false
		}

		return len(str) > 15
	}
	c := condition.Check(condFn)

	values := []interface{}{
		dummyStringValues[0],
		dummyStringValues[1],
		dummyIntValues[0],
	}

	test := CondTest{
		Values:       values,
		Expectations: makeExpectations(len(values), []int{1}),
	}

	testCond(t, c, test)
}

func TestFieldCheckEq(t *testing.T) {
	condTarget := "Field1.Field2"
	condVal := dummyStructValues[0].Field1.Field2

	c := condition.FieldCheck(condTarget, condition.Eq(condVal))

	values := []interface{}{
		dummyStructValues[0],
		dummyStructValues[1],
		dummyStructValues[2],
		dummyMapValue,
		dummyRogueStructValue,
	}

	test := CondTest{
		Values:       values,
		Expectations: makeExpectations(len(values), []int{0, 3}),
	}

	testCond(t, c, test)
}

func TestFieldCheck(t *testing.T) {
	condTarget := "Field1.Field3"
	condFn := func(x interface{}) bool {
		num, ok := x.(int)
		if !ok {
			return false
		}

		return num < 60
	}

	c := condition.FieldCheck(condTarget, condFn)

	values := []interface{}{
		dummyStructValues[0],
		dummyStructValues[1],
		dummyStructValues[2],
		dummyRogueStructValue,
	}

	test := CondTest{
		Values:       values,
		Expectations: makeExpectations(len(values), []int{1}),
	}

	testCond(t, c, test)
}
