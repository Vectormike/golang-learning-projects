package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	name    string `json:"Your name"`
	message string `json:"Birthday wish"`
}

type MyResponse struct {
	Message string `json:"My response"`
}

func sendBirthdayMessage(event MyEvent) (MyMessage, error) {
	return MyResponse{Message: fmt.Sprintf("Hi %s! %d.")}
}

func main() {
	lambda.Start(sendBirthdayMessage)
}
