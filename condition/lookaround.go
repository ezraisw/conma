package condition

type (
	lookaroundCond struct {
		fn        LookaroundCondFunc
		interval  int
		maxDist   int
		startDist int
		all       bool
	}

	LookaroundOption   func(c *lookaroundCond)
	LookaroundCondFunc func(x interface{}) Condition
)

// Matches to true if any elements before it satisfies the given condition.
func LookBeforeAny(cond Condition) Condition {
	return Lookaround(P(cond), -1)
}

// Matches to true if all elements before it satisfies the given condition.
func LookBeforeAll(cond Condition) Condition {
	return Lookaround(P(cond), -1, WithAll(true))
}

// Matches to true if any element after it satisfies the given condition.
func LookAfterAny(cond Condition) Condition {
	return Lookaround(P(cond), 1)
}

// Matches to true if all element after it satisfies the given condition.
func LookAfterAll(cond Condition) Condition {
	return Lookaround(P(cond), 1, WithAll(true))
}

// Matches to true if elements around the current element is satisfies the given condition.
func Lookaround(fn LookaroundCondFunc, interval int, options ...LookaroundOption) Condition {
	if interval == 0 {
		panic(ErrInvalidInterval)
	}

	c := lookaroundCond{
		fn:       fn,
		interval: interval,
	}

	for _, option := range options {
		option(&c)
	}

	return c
}

// Condition function for lookaround at the current element.
func P(cond Condition) LookaroundCondFunc {
	return func(x interface{}) Condition {
		return cond
	}
}

// The maximum distance from the current element.
func WithMaxDist(maxDist int) LookaroundOption {
	return func(c *lookaroundCond) {
		if maxDist < 0 {
			panic(ErrInvalidMaxDist)
		}

		if maxDist != 0 && c.startDist != 0 && maxDist < c.startDist {
			panic(ErrInvalidMaxOrStartDist)
		}

		c.maxDist = maxDist
	}
}

// The distance to start searching around the current element.
func WithStartDist(startDist int) LookaroundOption {
	return func(c *lookaroundCond) {
		if startDist < 0 {
			panic(ErrInvalidStartDist)
		}

		if startDist != 0 && c.maxDist != 0 && startDist > c.maxDist {
			panic(ErrInvalidMaxOrStartDist)
		}

		c.startDist = startDist
	}
}

// All elements around the current element must satisfy the condition.
func WithAll(all bool) LookaroundOption {
	return func(c *lookaroundCond) {
		c.all = all
	}
}

func (c lookaroundCond) Test(mctx MatchContext) bool {
	low := 0
	cLow := mctx.CurrentIndex - c.maxDist
	if c.maxDist != 0 && cLow >= 0 {
		low = cLow
	}

	high := len(mctx.Values) - 1
	cHigh := mctx.CurrentIndex + c.maxDist
	if c.maxDist != 0 && cHigh < len(mctx.Values) {
		high = cHigh
	}

	// Default the start to be:
	// - <current index> + 1 if interval is positive
	// - <current index> - 1 if interval is negative
	var start int
	if c.interval < 0 {
		if c.startDist != 0 {
			start = mctx.CurrentIndex - c.startDist
		} else {
			start = mctx.CurrentIndex - 1
		}
	} else if c.interval > 0 {
		if c.startDist != 0 {
			start = mctx.CurrentIndex + c.startDist
		} else {
			start = mctx.CurrentIndex + 1
		}
	}

	cond := c.fn(mctx.CurrentValue())

	if c.all {
		matched := false

		for j := start; j >= low && j <= high; j += c.interval {
			submctx := MatchContext{
				Values:       mctx.Values,
				CurrentIndex: j,
			}

			if !cond.Test(submctx) {
				return false
			}

			matched = true
		}

		return matched
	}

	for j := start; j >= low && j <= high; j += c.interval {
		submctx := MatchContext{
			Values:       mctx.Values,
			CurrentIndex: j,
		}

		if cond.Test(submctx) {
			return true
		}
	}

	return false
}
