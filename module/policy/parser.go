package policy

import (
	"bufio"
	"io"
	"strings"
)

const apiTokenStart = "[aws-sdk-go] DEBUG: Request "
const apiTokenEnd = " Details:"
const bodyEnd = "-----------------------------------------------------"

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
	s := bufio.NewScanner(p.trace)
	requests := make([]Request, 0)
	for s.Scan() {
		l := s.Text()
		if start := strings.Index(l, apiTokenStart); start != -1 {
			end := strings.Index(l, apiTokenEnd)
			key := l[start+len(apiTokenStart) : end]
			body := ""
			for s.Scan() {
				bl := s.Text()
				if strings.Contains(bl, bodyEnd) {
					break
				}
				body = body + bl
			}
			req := Request{
				apiKey: key,
				body:   body,
			}
			requests = append(requests, req)
		}
	}

	return requests, nil
}
