package policy

import "strings"

type Effect string

var allow Effect = "Allow"
var deny Effect = "Deny"

type Statements map[string]Statement

func NewStatements() Statements {
	return make(map[string]Statement)
}

func (s *Statements) List() []Statement {
	var out []Statement
	for _, v := range *s {
		out = append(out, v)
	}
	return out
}

type Statement struct {
	Effect     Effect
	Service    string
	Actions    map[string]bool
	Arn        string
	Conditions []Condition
}

type Condition struct {
	Cond  string
	Key   string
	Value string
}

func (s *Statements) AddRequests(reqs []Request) error {
	for _, r := range reqs {
		ns := requestToStatement(r)
		if _, ok := (*s)[ns.Service]; ok {
			for a, _ := range ns.Actions {
				(*s)[ns.Service].Actions[a] = true
			}
		} else {
			(*s)[ns.Service] = ns
		}
	}
	return nil
}

// requestToStatement translates a request made by terraform to a Statement in the policy
func requestToStatement(req Request) Statement {
	sp := strings.Split(req.apiKey, "/")
	service := sp[0]
	action := sp[1]
	actions := make(map[string]bool)
	actions[action] = true
	return Statement{
		Effect:  allow,
		Service: service,
		Actions: actions,
		Arn:     "",
	}
}

// PruneStatements groups actions together where multiple statements have the same effect and resource
func PruneStatements(stmts []Statement) ([]Statement, error) {
	p := make(map[string]Statement)
	for _, s := range stmts {
		if upd, ok := p[s.Service]; ok {
			for a := range s.Actions {
				upd.Actions[a] = true
				p[s.Service] = upd
			}
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
