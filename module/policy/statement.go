package policy

import "strings"

type Effect string

var allow Effect = "Allow"
var deny Effect = "Deny"

type Statement struct {
	Effect     Effect
	Service    string
	Actions    []string
	Arn        string
	Conditions []Condition
}

type Condition struct {
	Cond  string
	Key   string
	Value string
}

// RequestToStatement translates a request made by terraform to a Statement in the policy
func RequestToStatement(req Request) Statement {
	sp := strings.Split(req.apiKey, "/")
	service := sp[0]
	action := sp[1]
	return Statement{
		Effect:  allow,
		Service: service,
		Actions: []string{action},
		Arn:     "",
	}
}

// PruneStatements groups actions together where multiple statements have the same effect and resource
func PruneStatements(stmts []Statement) ([]Statement, error) {
	p := make(map[string]Statement)
	for _, s := range stmts {
		if upd, ok := p[s.Service]; ok {
			upd.Actions = append(p[s.Service].Actions, s.Actions...)
			p[s.Service] = upd
		} else {
			p[s.Service] = s
		}
	}
	return mapToSlice(p), nil
}

func mapToSlice(in map[string]Statement) []Statement {
	var out []Statement
	for _, v := range in {
		out = append(out, v)
	}
	return out
}
