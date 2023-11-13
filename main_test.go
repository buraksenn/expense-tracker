package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestLocal(t *testing.T) {
	handler(context.Background(), events.APIGatewayProxyRequest{
		Body: `{
			"update_id": 123456789,
			"message": {
				"message_id": 1,
				"from": {
					"id": 123456789,
					"is_bot": false,
					"first_name": "Burak",
					"last_name": "Sennaroglu",
					"username": "buraksenn",
					"language_code": "en"
				},
				"chat": {
					"id": 123456789,
					"first_name": "Burak",
					"last_name": "Sennaroglu",
					"username": "buraksenn",
					"type": "private"
				},
				"date": 1620000000,
				"text": "/register_expense"
			}`,
	})
}
