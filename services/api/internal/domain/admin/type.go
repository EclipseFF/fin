package admin

import "time"

type Request struct {
	requestedAt   time.Time
	requestMethod string
	requestBody   string
}

func New(method, body string) (*Request, error) {
	return &Request{
		requestedAt:   time.Now().UTC(),
		requestMethod: method,
		requestBody:   body,
	}, nil
}
