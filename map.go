package conma

import (
	"github.com/pwnedgod/conma/condition"
	"github.com/pwnedgod/conma/mapping"
)

type Entry struct {
	Cond   condition.Condition
	Mapper mapping.Mapper
}

type Map struct {
	entries []Entry
}

func New() *Map {
	return &Map{
		entries: make([]Entry, 0),
	}
}

func NewWithEntries(entries []Entry) *Map {
	return &Map{
		entries: entries,
	}
}

func (m *Map) Set(c condition.Condition, mapper mapping.Mapper) {
	m.entries = append(m.entries, Entry{
		Cond:   c,
		Mapper: mapper,
	})
}

func (m Map) MapSlice(values []interface{}) []interface{} {
	mapped := make([]interface{}, 0)
	for i := range values {
		for _, entry := range m.entries {
			mctx := condition.MatchContext{
				Values:       values,
				CurrentIndex: i,
			}

			if entry.Cond.Test(mctx) {
				mapped = append(mapped, entry.Mapper.Map(mctx.CurrentValue()))
			}
		}
	}

	return mapped
}
