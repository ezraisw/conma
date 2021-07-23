package condition

type MatchContext struct {
	// The array of values to be matched with.
	Values []interface{}

	// The current index of the value to be matched.
	CurrentIndex int
}

// Obtain the current value derived from the current index.
func (c MatchContext) CurrentValue() interface{} {
	return c.Values[c.CurrentIndex]
}
