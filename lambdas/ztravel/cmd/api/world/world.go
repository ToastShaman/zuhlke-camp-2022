package main

import (
	"context"
	"ztravel/pkg/api"
	"ztravel/pkg/aws"
	"ztravel/pkg/clock"
	"ztravel/pkg/env"
	"ztravel/pkg/world"

	"github.com/aws/aws-lambda-go/lambda"

	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
)

var chiLambda *chiadapter.ChiLambda

func init() {
	clock := clock.SystemClockUTC

	cfg := aws.LoadDefaultConfig()
	sm := aws.NewSecretsManager(aws.NewSecretsManagerClient(cfg))
	fn := sm.GetSecretValueFn(context.TODO(), env.Required("API_KEY_SM_NAME"))
	tokenAuth := api.NewTokenAuth(fn)
	signer := aws.NewSigner(env.Required("SIGNING_KEY_KMS_ID"), aws.NewKMSClient(cfg))
	api := world.NewWorldApi(tokenAuth, signer, clock)

	chiLambda = chiadapter.New(api)
}

func main() {
	lambda.Start(aws.NewAPIHandler(chiLambda))
}
