package aws

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/rs/zerolog/log"
)

type APIHandlerFunc = func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func NewAPIHandler(lambda *chiadapter.ChiLambda) APIHandlerFunc {
	return APIHandlerFunc(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		lc := NewLogContext(ctx, req)
		log.Logger = lc.InitialiseLogger()
		return lambda.ProxyWithContext(lc.NewContext(ctx), req)
	})
}
