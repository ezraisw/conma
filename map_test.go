package conma_test

import (
	"testing"

	"github.com/ezraisw/conma"
	"github.com/ezraisw/conma/condition"
	"github.com/ezraisw/conma/mapping"
	"github.com/stretchr/testify/assert"
)

type exampleStruct struct {
	Name    string
	Code    int
	Message string
}

func TestMap(t *testing.T) {
	slice := []interface{}{
		exampleStruct{
			Name:    "john",
			Code:    500,
			Message: "Example 1",
		},
		exampleStruct{
			Name:    "sebastian",
			Code:    700,
			Message: "Example 2",
		},
		exampleStruct{
			Name:    "<placeholder>",
			Code:    0,
			Message: "This is a placeholder for the next fields",
		},
		exampleStruct{
			Name:    "john",
			Code:    500,
			Message: "Example 3",
		},
	}

	expectedMapped := []interface{}{
		"Doe",
		"Winters",
		"Example 3",
	}

	m := conma.NewWithEntries([]conma.Entry{
		{
			Cond: condition.And(
				condition.Not(condition.FieldCheck("Message", condition.Eq("Example 3"))),
				condition.FieldCheck("Name", condition.Eq("john")),
			),
			Mapper: mapping.Value("Doe"),
		},
		{
			Cond:   condition.FieldCheck("Message", condition.Eq("Example 2")),
			Mapper: mapping.Value("Winters"),
		},
		{
			Cond: condition.And(
				condition.FieldCheck("Name", condition.Eq("john")),
				condition.LookBeforeAny(
					condition.FieldCheck("Name", condition.Eq("<placeholder>")),
				),
			),
			Mapper: func(x interface{}) interface{} {
				d := x.(exampleStruct)
				return d.Message
			},
		},
	})

	mapped := m.MapSlice(slice)
	assert.Equal(t, expectedMapped, mapped)
}
