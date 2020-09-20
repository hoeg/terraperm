package policy

import "io"

type Parser struct {
	trace io.Reader
}

type Request struct {
	apiKey string
	body   string
}

func NewParser(trace io.Reader) Parser {
	return Parser{trace: trace}
}

func (p *Parser) Requests() ([]Request, error) {
	return nil, nil
}
