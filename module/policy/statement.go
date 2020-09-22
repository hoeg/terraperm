package policy

import "strings"

type Effect string

var allow Effect = "Allow"
var deny Effect = "Deny"

type Statement struct {
	effect     Effect
	service    string
	actions    []string
	arn        string
	conditions []Condition
}

type Condition struct {
	cond  string
	key   string
	value string
}

// RequestToStatement translates a request made by terraform to a Statement in the policy
func RequestToStatement(req Request) (Statement, error) {
	sp := strings.Split(req.apiKey, "/")
	service := sp[0]
	action := sp[1]
	return Statement{
		effect:  allow,
		service: service,
		actions: []string{action},
		arn:     "",
	}, nil
}

// PruneStatements groups actions together where multiple statements have the same effect and resource
func PruneStatements(stmts []Statement) ([]Statement, error) {
	return nil, nil
}
