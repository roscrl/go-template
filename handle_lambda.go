package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

// HandleLambda is the entry point for the lambda function which will proxy to other handlers
func (s *Server) handleLambda(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return s.lambda.ProxyWithContext(ctx, req)
}
