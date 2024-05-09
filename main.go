package main

import (
	"fmt"

	"github.com/gstelang/rules-engine-golang.git/app"
	"github.com/gstelang/rules-engine-golang.git/rules"
)

func main() {

	// Example 1
	// Define a condition

	// Regex rule
	condition := rules.Condition{
		CdtType:    rules.MATCHES,
		Value:      ".*",
		EmailField: "Subject",
	}

	// Match rule
	// condition := rules.Condition{
	// 	CdtType:    rules.IS,
	// 	Value:      "Hello my dear friend",
	// 	EmailField: "Subject",
	// }

	// SQL rule
	// condition := rules.Condition{
	// 	CdtType:    rules.SQL,
	// 	Value:      "select * from bad_values",
	// 	EmailField: "Subject",
	// }

	// Define a rule using that condition
	ruleAlwaysMatch := rules.Rule{
		Name:      "Always Match",
		Min_Score: 0,
		Conditions: []rules.Condition{
			condition,
		},
	}

	// Email you want to score
	email := app.ScoredEmail{
		Email: app.Email{
			Subject:        "Hello my dear friend",
			Body:           "I wish you a wonderful day and need a lot of money.",
			From_email:     "hacker@email.com",
			Reply_to:       "hacker@email.com",
			Domain:         "email.com",
			Company_domain: "abc.com",
		},
		Score: 1,
	}

	// Rule engine
	alwaysMatchRulesEngine := &rules.RulesEngine{
		Rules: []rules.Rule{ruleAlwaysMatch},
	}

	// Evaluate
	rulenames := alwaysMatchRulesEngine.EvaluateScoredEmail(email)
	fmt.Println(rulenames)

	// Example 2...
}
