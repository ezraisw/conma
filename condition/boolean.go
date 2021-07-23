package condition

type (
	orCond  []Condition
	andCond []Condition
	notCond struct {
		cond Condition
	}
)

// Matches to true if any of the subconditions matches to true.
func Or(conds ...Condition) Condition {
	if len(conds) == 0 {
		panic(ErrEmptyCond)
	}

	return orCond(conds)
}

func (c orCond) Test(mctx MatchContext) bool {
	for _, cond := range c {
		if cond.Test(mctx) {
			return true
		}
	}

	return false
}

// Matches to true if all of the subconditions matches to true.
func And(conds ...Condition) Condition {
	if len(conds) == 0 {
		panic(ErrEmptyCond)
	}

	return andCond(conds)
}

func (c andCond) Test(mctx MatchContext) bool {
	for _, cond := range c {
		if !cond.Test(mctx) {
			return false
		}
	}

	return true
}

// Negates the subcondition's matching result.
func Not(cond Condition) Condition {
	return notCond{cond: cond}
}

func (c notCond) Test(mctx MatchContext) bool {
	return !c.cond.Test(mctx)
}
