package rules

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/gstelang/rules-engine-golang.git/app"
	_ "github.com/mattn/go-sqlite3"
)

type EmailScoreEvaluator interface {
	EvaluateScoredEmail(email app.ScoredEmail) []string
}

type RulesEngine struct {
	Rules []Rule
}

var (
	db     *sql.DB
	dbOnce sync.Once
)

func GetDB() (*sql.DB, error) {
	var err error
	dbOnce.Do(func() {
		db, err = sql.Open("sqlite3", "./rules_db.sqlite")
	})
	return db, err
}

func getEmailRepo() *EmailRepo {
	db, err := GetDB()
	if err != nil {
		panic(err)
	}

	return NewEmailRepo(&SQLLite{db: db})
}

func (engine RulesEngine) EvaluateScoredEmail(email app.ScoredEmail) []string {
	matchedRules := []string{}

	for _, rule := range engine.Rules {
		if email.Score > rule.Min_Score {
			for _, condition := range rule.Conditions {
				isMatched := false
				switch condition.CdtType {
				case IS, IS_NOT:
					matchRule := MatchRule{
						ScrEmail: app.ScoredEmail{
							Email: email.Email,
							Score: rule.Min_Score,
						},
						Condition: condition,
					}
					isMatched = matchRule.Evaluate()
				case MATCHES, NOT_MATCHES:
					regexRule := RegexRule{
						ScrEmail: app.ScoredEmail{
							Email: email.Email,
							Score: rule.Min_Score,
						},
						Condition: condition,
					}
					isMatched = regexRule.Evaluate()
				case SQL:
					sqlRule := SQLRule{
						ScrEmail: app.ScoredEmail{
							Email: email.Email,
							Score: rule.Min_Score,
						},
						Condition: condition,
						Repo:      *getEmailRepo(),
					}
					isMatched = sqlRule.Evaluate()
				default:
					fmt.Println("Condition type is not valid")
				}

				if isMatched == true {
					matchedRules = append(matchedRules, rule.Name)
				}
			}
		}
	}

	return matchedRules
}
