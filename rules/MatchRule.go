package rules

import "github.com/gstelang/rules-engine-golang.git/app"

type MatchRule struct {
	ScrEmail  app.ScoredEmail
	Condition Condition
}

func (matchRule MatchRule) Evaluate() bool {
	matched := false

	conditionType := matchRule.Condition.CdtType
	value := matchRule.Condition.Value
	emailField := matchRule.Condition.EmailField

	emailFieldValue := GetFieldValue(matchRule.ScrEmail.Email, emailField)

	if conditionType == IS {
		return value == emailFieldValue
	} else if conditionType == IS_NOT {
		return value != emailFieldValue
	}

	return matched
}
