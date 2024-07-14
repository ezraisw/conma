package conma

import (
	"github.com/ezraisw/conma/condition"
	"github.com/ezraisw/conma/mapping"
)

type Entry struct {
	// The condition to satisfy.
	Cond condition.Condition

	// The mapper which will produce the value.
	Mapper mapping.MapperFunc
}

type Map struct {
	entries []Entry
}

// Create a new empty conditional map.
func New() *Map {
	return &Map{
		entries: make([]Entry, 0),
	}
}

// Create a conditional map with the given entries.
func NewWithEntries(entries []Entry) *Map {
	return &Map{
		entries: entries,
	}
}

// Set a new entry for the map.
func (m *Map) Set(cond condition.Condition, mapper mapping.MapperFunc) {
	m.entries = append(m.entries, Entry{
		Cond:   cond,
		Mapper: mapper,
	})
}

// Map a slice from the list of entries.
//
// Mapping a slice is a O(mn) operation where
// m is the number of entries and n the number of elements in the slice.
//
// It is always faster to use Go map when only equality is used.
func (m Map) MapSlice(values []interface{}) []interface{} {
	mapped := make([]interface{}, 0)
	for i := range values {
		for _, entry := range m.entries {
			mctx := condition.MatchContext{
				Values:       values,
				CurrentIndex: i,
			}

			if entry.Cond.Test(mctx) {
				mapped = append(mapped, entry.Mapper(mctx.CurrentValue()))
			}
		}
	}

	return mapped
}
