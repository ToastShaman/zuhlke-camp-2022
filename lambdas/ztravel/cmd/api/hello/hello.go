package main

import (
	"ztravel/pkg/aws"
	"ztravel/pkg/env"
	"ztravel/pkg/hello"

	"github.com/aws/aws-lambda-go/lambda"

	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
)

var chiLambda *chiadapter.ChiLambda

func init() {
	cfg := aws.LoadDefaultConfig()
	s3 := aws.NewS3Storage(env.Required("BUCKET_NAME"), aws.NewS3Client(cfg))
	sqs := aws.NewSQS(env.Required("QUEUE_URL"), aws.NewSQSClient(cfg))
	api := hello.NewHelloApi(s3, sqs)

	chiLambda = chiadapter.New(api)
}

func main() {
	lambda.Start(aws.NewAPIHandler(chiLambda))
}
