package condition_test

import (
	"testing"

	"github.com/pwnedgod/conma/condition"
	"github.com/stretchr/testify/assert"
)

type CondTestExpectation struct {
	Index   int
	Success bool
}

type CondTest struct {
	Values       []interface{}
	Expectations []CondTestExpectation
}

type exampleStruct struct {
	Field1 *exampleSubStruct1
	Field2 *exampleSubStruct2
	Field3 string
	Field4 string
}

type exampleSubStruct1 struct {
	Field1 string
	Field2 int
	Field3 interface{}
}

type exampleSubStruct2 struct {
	Field1 exampleSubSubStruct
	Field2 string
}

type exampleSubSubStruct struct {
	Field1 string
	Field2 string
}

type exampleRogueStruct struct {
	RogueField string
}

var (
	dummyStructValues = []exampleStruct{
		{
			Field1: &exampleSubStruct1{
				Field1: "example",
				Field2: 420,
				Field3: nil,
			},
			Field2: &exampleSubStruct2{
				Field1: exampleSubSubStruct{
					Field1: "subvalue1",
					Field2: "subvalue2",
				},
				Field2: "testing",
			},
			Field3: "verylongstring-verylongstring",
			Field4: "short",
		},
		{
			Field1: &exampleSubStruct1{
				Field1: "empty",
				Field2: 34,
				Field3: 50,
			},
			Field2: nil,
			Field3: "verylongstring-verylongstring",
			Field4: "short",
		},
		{
			Field1: nil,
			Field2: nil,
			Field3: "verylongstring-verylongstring",
			Field4: "short",
		},
	}

	dummyMapValue = map[string]interface{}{
		"Field1": map[string]int{
			"Field2": dummyStructValues[0].Field1.Field2,
		},
	}

	dummyRogueStructValue = exampleRogueStruct{
		RogueField: "invalid",
	}

	dummyStringValues = []string{
		"dummy value 1",
		"dummy value 2 (but very long)",
	}

	dummyIntValues = []int{
		399,
		391,
		71,
		439,
		136,
		37,
		71,
		24,
		371,
		245,
	}
)

func testCond(t *testing.T, c condition.Condition, test CondTest) {
	for _, ex := range test.Expectations {
		if ex.Success {
			assert.True(
				t,
				c.Test(condition.MatchContext{
					Values:       test.Values,
					CurrentIndex: ex.Index,
				}),
				"Index: %d",
				ex.Index,
			)
		} else {
			assert.Falsef(
				t,
				c.Test(condition.MatchContext{
					Values:       test.Values,
					CurrentIndex: ex.Index,
				}),
				"Index: %d",
				ex.Index,
			)
		}
	}
}

func makeExpectations(len int, successIdxs []int) []CondTestExpectation {
	successes := make([]bool, len)

	for _, i := range successIdxs {
		successes[i] = true
	}

	expectations := make([]CondTestExpectation, 0)
	for i := 0; i < len; i++ {
		expectations = append(expectations, CondTestExpectation{
			Index:   i,
			Success: successes[i],
		})
	}

	return expectations
}
