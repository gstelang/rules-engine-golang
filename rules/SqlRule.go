package rules

import (
	"database/sql"
	"fmt"

	"github.com/gstelang/rules-engine-golang.git/app"
)

type BadValue struct {
	fieldType  string
	fieldvalue string
}

type SQLLite struct {
	db *sql.DB
}

type ReadService interface {
	ExecuteSelect(stmt string) ([]*BadValue, error)
}

func (s *SQLLite) ExecuteSelect(stmt string) ([]*BadValue, error) {
	rows, err := s.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var badValues []*BadValue

	for rows.Next() {
		var badValue BadValue
		err := rows.Scan(&badValue.fieldType, &badValue.fieldvalue)
		if err != nil {
			return nil, err
		}
		badValues = append(badValues, &badValue)
	}

	return badValues, nil
}

type EmailRepo struct {
	db ReadService
}

// You can now mock this by mocking ExecuteSelect
func NewEmailRepo(db ReadService) *EmailRepo {
	return &EmailRepo{db: db}
}

type SQLRule struct {
	ScrEmail  app.ScoredEmail
	Condition Condition
	Repo      EmailRepo
}

func (sqlRule SQLRule) Evaluate() bool {
	matched := false

	conditionType := sqlRule.Condition.CdtType
	value := sqlRule.Condition.Value

	// TODO: value startsWith("Select")
	// TODO: use strings Replacer method to replace in SQL query
	// Also the assumption is the rules engine is not user facing and hence SQL injection is not likely. 

	if conditionType == SQL {
		rows, err := sqlRule.Repo.db.ExecuteSelect(value)
		if err != nil {
			fmt.Println(err)
			panic("Error executing select")
		}
		matched = len(rows) > 0
	}

	return matched
}
