package http

import (
	"context"

	"eventsguard/internal/hello/dtos"
)

type helloSimpleHandler struct {
}

func NewHelloSimpleHandler() *helloSimpleHandler {
	return &helloSimpleHandler{}
}

func (h helloSimpleHandler) Handler(
	ctx context.Context,
	input *struct{},
) (*dtos.HelloSimpleResponse, error) {

	resp := &dtos.HelloSimpleResponse{
		Body: dtos.HelloSimpleOutput{
			Message: "oli Huma, ke ase?",
		},
	}
	return resp, nil
}
