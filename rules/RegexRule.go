package rules

import (
	"regexp"

	"github.com/gstelang/rules-engine-golang.git/app"
)

type RegexRule struct {
	ScrEmail  app.ScoredEmail
	Condition Condition
}

func (regexRule RegexRule) Evaluate() bool {
	matched := false
	conditionType := regexRule.Condition.CdtType
	value := regexRule.Condition.Value
	emailField := regexRule.Condition.EmailField

	emailFieldValue := GetFieldValue(regexRule.ScrEmail.Email, emailField)
	re := regexp.MustCompile(value)
	match := re.MatchString(emailFieldValue)

	if conditionType == MATCHES {
		return match == true
	} else if conditionType == NOT_MATCHES {
		return match == false
	}

	return matched
}
