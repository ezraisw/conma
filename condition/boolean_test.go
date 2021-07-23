package condition_test

import (
	"testing"

	"github.com/pwnedgod/conma/condition"
	"github.com/stretchr/testify/assert"
)

func TestAndPanicEmptyCond(t *testing.T) {
	assert.PanicsWithError(
		t,
		condition.ErrEmptyCond.Error(),
		func() {
			condition.And()
		},
	)
}

func TestAnd(t *testing.T) {
	intValues := dummyIntValues

	c := condition.And(
		condition.Check(func(x interface{}) bool {
			num, ok := x.(int)
			if !ok {
				return false
			}

			return num > 350
		}),
		condition.Check(func(x interface{}) bool {
			num, ok := x.(int)
			if !ok {
				return false
			}

			return num < 400
		}),
	)

	test := CondTest{
		Values:       make([]interface{}, 0),
		Expectations: makeExpectations(len(intValues), []int{0, 1, 8}),
	}
	for _, x := range intValues {
		test.Values = append(test.Values, x)
	}

	testCond(t, c, test)
}

func TestOrPanicEmptyCond(t *testing.T) {
	assert.PanicsWithError(
		t,
		condition.ErrEmptyCond.Error(),
		func() {
			condition.Or()
		},
	)
}

func TestOr(t *testing.T) {
	intValues := dummyIntValues

	c := condition.Or(
		condition.Check(condition.Eq(intValues[0])),
		condition.Check(condition.Eq(intValues[1])),
	)

	test := CondTest{
		Values:       make([]interface{}, 0),
		Expectations: makeExpectations(len(intValues), []int{0, 1}),
	}
	for _, x := range intValues {
		test.Values = append(test.Values, x)
	}

	testCond(t, c, test)
}

func TestNot(t *testing.T) {
	intValues := dummyIntValues

	c := condition.Not(
		condition.Check(func(x interface{}) bool {
			num, ok := x.(int)
			if !ok {
				return false
			}

			return num > 400
		}),
	)

	test := CondTest{
		Values:       make([]interface{}, 0),
		Expectations: makeExpectations(len(intValues), []int{0, 1, 2, 4, 5, 6, 7, 8, 9}),
	}
	for _, x := range intValues {
		test.Values = append(test.Values, x)
	}

	testCond(t, c, test)
}
