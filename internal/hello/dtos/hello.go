package dtos

type HelloSimpleOutput struct {
	Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
}

type HelloSimpleResponse struct {
	Body HelloSimpleOutput
}
