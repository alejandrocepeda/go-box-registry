package box

import "context"

type Result struct {
	Output   any    `json:"output"`
	NewEvent string `json:"new_event,omitempty"`
}

type Box interface {
	ID() string
	Execute(ctx context.Context, input any) (Result, error)
}
