package policy

import (
	"encoding/json"
)

const policyVersion = "2012-10-17"

type iamPolicy struct {
	Version   string
	Statement []Statement
}

func Print(stms []Statement) (string, error) {
	p := iamPolicy{
		Version:   policyVersion,
		Statement: stms,
	}

	out, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}
