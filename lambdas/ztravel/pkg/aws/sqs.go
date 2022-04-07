package aws

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Queue interface {
	Send(ctx context.Context, payload interface{}) (string, error)
}

type SQS struct {
	Client *sqs.Client
	URL    string
}

func (s *SQS) Send(ctx context.Context, payload interface{}) (string, error) {
	message, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	output, err := s.Client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(s.URL),
		MessageBody: aws.String(string(message)),
	})

	if err != nil {
		return "", err
	}

	return *output.MessageId, nil
}

func NewSQSClient(cfg aws.Config) *sqs.Client {
	return sqs.NewFromConfig(cfg)
}

func NewSQS(url string, client *sqs.Client) *SQS {
	return &SQS{
		Client: client,
		URL:    url,
	}
}
