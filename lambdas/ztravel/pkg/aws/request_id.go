package aws

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// See https://github.com/awslabs/aws-lambda-go-api-proxy/pull/110
// See https://aws.amazon.com/blogs/compute/techniques-and-tools-for-better-serverless-api-logging-with-amazon-api-gateway-and-aws-lambda/

type ctxLogContextKey int

const logContextKey ctxLogContextKey = 0

type LogContext struct {
	ApiRequestId    string
	LambdaRequestId string
}

func (lc *LogContext) NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, logContextKey, lc)
}

func (lc *LogContext) InitialiseLogger() zerolog.Logger {
	return log.With().
		Str("apiRequestId", lc.ApiRequestId).
		Str("lambdaRequestId", lc.LambdaRequestId).
		Logger()
}

func NewLogContext(ctx context.Context, req events.APIGatewayProxyRequest) *LogContext {
	c, ok := lambdacontext.FromContext(ctx)
	if !ok {
		log.Fatal().Msg("Failed to create lambda context")
	}

	return &LogContext{
		ApiRequestId:    req.RequestContext.RequestID,
		LambdaRequestId: c.AwsRequestID,
	}
}

func LogContextFromContext(ctx context.Context) (*LogContext, bool) {
	l, ok := ctx.Value(logContextKey).(*LogContext)
	return l, ok
}

func RequestIdHeaders(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		lc, ok := LogContextFromContext(r.Context())
		if !ok {
			log.Fatal().Msg("Failed to retrieve log context")
		}

		r.Header.Set("X-Request-Id", lc.ApiRequestId)
		r.Header.Set("X-Api-Request-Id", lc.ApiRequestId)
		r.Header.Set("X-Lambda-Request-Id", lc.LambdaRequestId)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
