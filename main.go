package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name    string `json:"what is your name?"`
	Email   string `json:"What is your email?"`
	Comment string `json:"send a message to the owner of this page"`
}

type ReturnEvent struct {
	Message string `json:"return email"`
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(event MyEvent) (ReturnEvent, error) {
	return ReturnEvent{
		Message: fmt.Sprintf("Name: %s, Email: %s, Comment: %s",
			event.Name, event.Email, event.Comment),
	}, nil
}
