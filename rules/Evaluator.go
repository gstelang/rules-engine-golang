package rules

type RuleEvaluator interface {
	Evaluate() bool
}
