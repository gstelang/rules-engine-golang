package rules

type ConditionType string

const (
	IS          ConditionType = "IS"
	IS_NOT      ConditionType = "IS_NOT"
	MATCHES     ConditionType = "MATCHES"
	NOT_MATCHES ConditionType = "NOT_MATCHES"
	SQL         ConditionType = "SQL"
)

type Condition struct {
	CdtType    ConditionType
	Value      string
	EmailField string
}

type Rule struct {
	Name       string
	Min_Score  float64
	Conditions []Condition
}
