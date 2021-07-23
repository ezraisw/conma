package condition_test

import (
	"testing"

	"github.com/pwnedgod/conma/condition"
	"github.com/stretchr/testify/assert"
)

func TestLookBeforeAny(t *testing.T) {
	intValues := dummyIntValues

	condVal := intValues[4]
	c := condition.LookBeforeAny(
		condition.Check(condition.Eq(condVal)),
	)

	test := CondTest{
		Values:       make([]interface{}, 0),
		Expectations: makeExpectations(len(intValues), []int{5, 6, 7, 8, 9}),
	}
	for _, x := range intValues {
		test.Values = append(test.Values, x)
	}

	testCond(t, c, test)
}

func TestLookBeforeAll(t *testing.T) {
	intValues := dummyIntValues

	c := condition.LookBeforeAll(
		condition.Check(func(x interface{}) bool {
			num, ok := x.(int)
			if !ok {
				return false
			}

			return num > 300
		}),
	)

	test := CondTest{
		Values:       make([]interface{}, 0),
		Expectations: makeExpectations(len(intValues), []int{1, 2}),
	}
	for _, x := range intValues {
		test.Values = append(test.Values, x)
	}

	testCond(t, c, test)
}

func TestLookAfterAny(t *testing.T) {
	intValues := dummyIntValues

	condVal := intValues[5]
	c := condition.LookAfterAny(
		condition.Check(condition.Eq(condVal)),
	)

	test := CondTest{
		Values:       make([]interface{}, 0),
		Expectations: makeExpectations(len(intValues), []int{0, 1, 2, 3, 4}),
	}
	for _, x := range intValues {
		test.Values = append(test.Values, x)
	}

	testCond(t, c, test)
}

func TestLookAfterAll(t *testing.T) {
	intValues := dummyIntValues

	c := condition.LookAfterAll(
		condition.Check(func(x interface{}) bool {
			num, ok := x.(int)
			if !ok {
				return false
			}

			return num > 200
		}),
	)

	test := CondTest{
		Values:       make([]interface{}, 0),
		Expectations: makeExpectations(len(intValues), []int{7, 8}),
	}
	for _, x := range intValues {
		test.Values = append(test.Values, x)
	}

	testCond(t, c, test)
}

func TestLookaroundPanicInvalidInterval(t *testing.T) {
	assert.PanicsWithError(
		t,
		condition.ErrInvalidInterval.Error(),
		func() {
			condition.Lookaround(condition.P(condition.Check(condition.Eq(0))), 0)
		},
	)
}

func TestLookaroundPanicInvalidMaxDist(t *testing.T) {
	assert.PanicsWithError(
		t,
		condition.ErrInvalidMaxDist.Error(),
		func() {
			condition.Lookaround(
				condition.P(condition.Check(condition.Eq(0))),
				1,
				condition.WithMaxDist(-10),
			)
		},
	)
}

func TestLookaroundPanicInvalidStartDist(t *testing.T) {
	assert.PanicsWithError(
		t,
		condition.ErrInvalidStartDist.Error(),
		func() {
			condition.Lookaround(
				condition.P(condition.Check(condition.Eq(0))),
				1,
				condition.WithStartDist(-10),
			)
		},
	)
}

func TestLookaroundPanicInvalidMaxOrStartDist(t *testing.T) {
	assert.PanicsWithError(
		t,
		condition.ErrInvalidMaxOrStartDist.Error(),
		func() {
			condition.Lookaround(
				condition.P(condition.Check(condition.Eq(0))),
				1,
				condition.WithMaxDist(5),
				condition.WithStartDist(10),
			)
		},
	)

	assert.PanicsWithError(
		t,
		condition.ErrInvalidMaxOrStartDist.Error(),
		func() {
			condition.Lookaround(
				condition.P(condition.Check(condition.Eq(0))),
				1,
				condition.WithStartDist(10),
				condition.WithMaxDist(5),
			)
		},
	)
}

func TestLookaroundNegativeIntervalWithMaxDist(t *testing.T) {
	intValues := []int{
		1000,
		300,
		70,
		70,
		300,
		70,
		300,
		300,
		70,
	}

	c := condition.And(
		condition.Check(condition.Eq(70)),
		condition.Lookaround(
			condition.P(
				condition.Check(func(x interface{}) bool {
					num, ok := x.(int)
					if !ok {
						return false
					}

					return num < 200
				}),
			),
			-1,
			condition.WithMaxDist(2),
		),
	)

	test := CondTest{
		Values:       make([]interface{}, 0),
		Expectations: makeExpectations(len(intValues), []int{3, 5}),
	}
	for _, x := range intValues {
		test.Values = append(test.Values, x)
	}

	testCond(t, c, test)
}

func TestLookaroundNegativeIntervalWithStartDist(t *testing.T) {
	intValues := []int{
		1000,
		300,
		70,
		70,
		300,
		70,
		300,
		300,
		70,
	}

	c := condition.And(
		condition.Check(condition.Eq(70)),
		condition.Lookaround(
			condition.P(
				condition.Check(func(x interface{}) bool {
					num, ok := x.(int)
					if !ok {
						return false
					}

					return num < 200
				}),
			),
			-1,
			condition.WithStartDist(3),
		),
	)

	test := CondTest{
		Values:       make([]interface{}, 0),
		Expectations: makeExpectations(len(intValues), []int{5, 8}),
	}
	for _, x := range intValues {
		test.Values = append(test.Values, x)
	}

	testCond(t, c, test)
}

func TestLookaroundNegativeIntervalWithMaxDistStartDistAndAll(t *testing.T) {
	intValues := []int{
		70,
		1000,
		300,
		70,
		70,
		300,
		300,
		70,
		70,
		300,
		1000,
	}

	c := condition.And(
		condition.Check(condition.Eq(300)),
		condition.Lookaround(
			condition.P(condition.Check(condition.Eq(70))),
			-1,
			condition.WithMaxDist(3),
			condition.WithStartDist(2),
			condition.WithAll(true),
		),
	)

	test := CondTest{
		Values:       make([]interface{}, 0),
		Expectations: makeExpectations(len(intValues), []int{2, 6}),
	}
	for _, x := range intValues {
		test.Values = append(test.Values, x)
	}

	testCond(t, c, test)
}

func TestLookaroundPositiveIntervalWithMaxDist(t *testing.T) {
	intValues := []int{
		70,
		300,
		300,
		70,
		300,
		70,
		70,
		300,
		1000,
		50,
	}

	c := condition.And(
		condition.Check(condition.Eq(70)),
		condition.Lookaround(
			condition.P(
				condition.Check(func(x interface{}) bool {
					num, ok := x.(int)
					if !ok {
						return false
					}

					return num < 200
				}),
			),
			1,
			condition.WithMaxDist(2),
		),
	)

	test := CondTest{
		Values:       make([]interface{}, 0),
		Expectations: makeExpectations(len(intValues), []int{3, 5}),
	}
	for _, x := range intValues {
		test.Values = append(test.Values, x)
	}

	testCond(t, c, test)
}

func TestLookaroundPositiveIntervalWithStartDist(t *testing.T) {
	intValues := []int{
		70,
		300,
		300,
		70,
		300,
		70,
		70,
		300,
		1000,
	}

	c := condition.And(
		condition.Check(condition.Eq(70)),
		condition.Lookaround(
			condition.P(
				condition.Check(func(x interface{}) bool {
					num, ok := x.(int)
					if !ok {
						return false
					}

					return num < 200
				}),
			),
			1,
			condition.WithStartDist(3),
		),
	)

	test := CondTest{
		Values:       make([]interface{}, 0),
		Expectations: makeExpectations(len(intValues), []int{0, 3}),
	}
	for _, x := range intValues {
		test.Values = append(test.Values, x)
	}

	testCond(t, c, test)
}

func TestLookaroundPositiveIntervalWithMaxDistStartDistAndAll(t *testing.T) {
	intValues := []int{
		70,
		70,
		300,
		300,
		70,
		70,
		300,
		1000,
	}

	c := condition.And(
		condition.Check(condition.Eq(70)),
		condition.Lookaround(
			condition.P(condition.Check(condition.Eq(300))),
			1,
			condition.WithMaxDist(3),
			condition.WithStartDist(2),
			condition.WithAll(true),
		),
	)

	test := CondTest{
		Values:       make([]interface{}, 0),
		Expectations: makeExpectations(len(intValues), []int{0}),
	}
	for _, x := range intValues {
		test.Values = append(test.Values, x)
	}

	testCond(t, c, test)
}
