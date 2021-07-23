package condition

type Condition interface {
	// Test a condition for an element at a given the match context.
	Test(mctx MatchContext) bool
}
