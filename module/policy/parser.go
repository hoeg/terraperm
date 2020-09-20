package policy

import (
	"bufio"
	"io"
	"strings"
)

const apiTokenStart = "[aws-sdk-go] DEBUG: Request"
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
	r := bufio.NewReader(p.trace)
	requests := make([]Request, 1)
	for {
		if r.Buffered() == 0 {
			break
		}
		l, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		if start := strings.Index(l, apiTokenStart); start != -1 {
			end := strings.Index(l, apiTokenEnd)
			key := l[start:end] //is start before the token?
			body := ""
			for {
				bl, err := r.ReadString('\n')
				if err != nil {
					return nil, err
				}
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
