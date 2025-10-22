package presentation

import (
	"context"
)

type HttpRequest struct {
	Headers map[string][]string
	Body    interface{}
	Params  map[string]string
	Query   map[string][]string
	Ctx     context.Context
}

type HttpResponse struct {
	Status  int
	Headers map[string]string
	Body    interface{}
}

type IHandler interface {
	Handle(req *HttpRequest) (*HttpResponse, error)
}
