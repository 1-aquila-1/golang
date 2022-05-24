package models

import "github.com/aws/aws-lambda-go/events"

type Usuario struct {
	Id      string
	Usuario string
	Email   string
}
type Response events.APIGatewayProxyResponse
