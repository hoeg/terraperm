package policy

import (
	"encoding/json"
)

const policyVersion = "2012-10-17"

type outStatement struct {
	Effect     Effect
	Service    string
	Actions    []string
	Arn        string
	Conditions []Condition
}

type iamPolicy struct {
	Version   string
	Statement []outStatement
}

func Print(stmts []Statement) (string, error) {
	p := iamPolicy{
		Version:   policyVersion,
		Statement: convertStatements(stmts),
	}

	out, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func convertStatements(stmts []Statement) []outStatement {
	var out []outStatement
	for _, s := range stmts {
		out = append(out, outStatement{
			Effect:     s.Effect,
			Service:    s.Service,
			Actions:    mapToSlice(s.Actions),
			Arn:        s.Arn,
			Conditions: s.Conditions,
		})
	}
	return out
}

func mapToSlice(in map[string]bool) []string {
	var out []string
	for k := range in {
		out = append(out, k)
	}
	return out
}
